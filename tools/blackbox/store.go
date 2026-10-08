package main

// On-disk format: a directory of gzip-compressed JSON-lines segments named
// bb-<UTC start time>.jsonl.gz. The writer flushes the gzip stream
// periodically (never fsyncs), so a segment cut short by a crash or power loss
// is readable up to its last flush. Old segments are deleted to keep the
// directory under a fixed size.

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync/atomic"
	"time"
)

const segPrefix, segSuffix = "bb-", ".jsonl.gz"
const segTimeFormat = "20060102T150405.000Z"

type StoreConfig struct {
	Dir         string
	FlushEvery  time.Duration
	SegMaxAge   time.Duration
	SegMaxBytes int64
	MaxTotal    int64
}

type countWriter struct {
	w io.Writer
	n int64
}

func (c *countWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.n += int64(n)
	return n, err
}

type Store struct {
	cfg      StoreConfig
	ch       chan []byte
	flushReq chan struct{}
	done     chan struct{}
	dropped  atomic.Uint64

	f        *os.File
	cw       *countWriter
	gz       *gzip.Writer
	segStart time.Time
}

// OpenStore starts the writer goroutine. Writing happens off the sampling
// path so that a stalled disk (the very thing being diagnosed) never stops
// sampling; if the queue fills, records are dropped and counted.
func OpenStore(cfg StoreConfig) (*Store, error) {
	if err := os.MkdirAll(cfg.Dir, 0o755); err != nil {
		return nil, err
	}
	s := &Store{
		cfg:      cfg,
		ch:       make(chan []byte, 512),
		flushReq: make(chan struct{}, 1),
		done:     make(chan struct{}),
	}
	s.prune()
	go s.run()
	return s, nil
}

func (s *Store) Put(rec *Record) {
	b, err := json.Marshal(rec)
	if err != nil {
		log.Printf("encode: %v", err)
		return
	}
	select {
	case s.ch <- append(b, '\n'):
	default:
		s.dropped.Add(1)
	}
}

// TakeDropped returns and resets the number of dropped records.
func (s *Store) TakeDropped() uint64 { return s.dropped.Swap(0) }

// FlushSoon asks the writer to push buffered data to the file.
func (s *Store) FlushSoon() {
	select {
	case s.flushReq <- struct{}{}:
	default:
	}
}

func (s *Store) Close() {
	close(s.ch)
	<-s.done
}

func (s *Store) run() {
	defer close(s.done)
	t := time.NewTicker(s.cfg.FlushEvery)
	defer t.Stop()
	for {
		select {
		case b, ok := <-s.ch:
			if !ok {
				s.closeSeg()
				return
			}
			s.write(b)
		case <-s.flushReq:
			s.flush()
		case <-t.C:
			s.flush()
			if s.gz != nil && (time.Since(s.segStart) >= s.cfg.SegMaxAge || s.cw.n >= s.cfg.SegMaxBytes) {
				s.closeSeg()
				s.prune()
			}
		}
	}
}

func (s *Store) write(b []byte) {
	if s.gz == nil {
		if err := s.openSeg(); err != nil {
			log.Printf("open segment: %v", err)
			return
		}
	}
	if _, err := s.gz.Write(b); err != nil {
		log.Printf("write: %v", err)
		s.closeSeg()
	}
}

func (s *Store) openSeg() error {
	now := time.Now()
	name := filepath.Join(s.cfg.Dir, segPrefix+now.UTC().Format(segTimeFormat)+segSuffix)
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return err
	}
	s.f, s.cw, s.segStart = f, &countWriter{w: f}, now
	s.gz, _ = gzip.NewWriterLevel(s.cw, gzip.DefaultCompression)
	return nil
}

func (s *Store) flush() {
	if s.gz == nil {
		return
	}
	if err := s.gz.Flush(); err != nil {
		log.Printf("flush: %v", err)
		s.closeSeg()
	}
}

func (s *Store) closeSeg() {
	if s.gz == nil {
		return
	}
	if err := s.gz.Close(); err != nil {
		log.Printf("close segment: %v", err)
	}
	s.f.Close()
	s.f, s.cw, s.gz = nil, nil, nil
}

// prune deletes the oldest closed segments until the directory fits.
func (s *Store) prune() {
	segs := listSegments(s.cfg.Dir)
	var total int64
	for _, sg := range segs {
		total += sg.size
	}
	for i := 0; total > s.cfg.MaxTotal && i < len(segs); i++ {
		if s.f != nil && segs[i].path == s.f.Name() {
			continue
		}
		if err := os.Remove(segs[i].path); err == nil {
			total -= segs[i].size
		}
	}
}

type segment struct {
	path  string
	start time.Time
	size  int64
}

func listSegments(dir string) []segment {
	ents, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var out []segment
	for _, e := range ents {
		n := e.Name()
		if !strings.HasPrefix(n, segPrefix) || !strings.HasSuffix(n, segSuffix) {
			continue
		}
		t, err := time.Parse(segTimeFormat, strings.TrimSuffix(strings.TrimPrefix(n, segPrefix), segSuffix))
		if err != nil {
			continue
		}
		var size int64
		if fi, err := e.Info(); err == nil {
			size = fi.Size()
		}
		out = append(out, segment{filepath.Join(dir, n), t, size})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].start.Before(out[j].start) })
	return out
}

// ReadRecords calls fn for every record with since <= T < until, in order.
// A zero since or until leaves that end open. Segments truncated by a crash
// are read up to the point of truncation.
func ReadRecords(dir string, since, until time.Time, fn func(*Record)) error {
	segs := listSegments(dir)
	if len(segs) == 0 {
		return errors.New("no data files in " + dir)
	}
	sinceMs, untilMs := int64(0), int64(1<<62)
	if !since.IsZero() {
		sinceMs = since.UnixMilli()
	}
	if !until.IsZero() {
		untilMs = until.UnixMilli()
	}
	for i, sg := range segs {
		if i+1 < len(segs) && !since.IsZero() && segs[i+1].start.Before(since) {
			continue
		}
		if !until.IsZero() && sg.start.After(until) {
			break
		}
		if err := readSegment(sg.path, sinceMs, untilMs, fn); err != nil {
			log.Printf("%s: %v", sg.path, err)
		}
	}
	return nil
}

func readSegment(path string, sinceMs, untilMs int64, fn func(*Record)) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	gz, err := gzip.NewReader(bufio.NewReader(f))
	if err != nil {
		if err == io.EOF {
			return nil // created but nothing flushed yet
		}
		return err
	}
	sc := bufio.NewScanner(gz)
	sc.Buffer(make([]byte, 64<<10), 8<<20)
	for sc.Scan() {
		var rec Record
		if json.Unmarshal(sc.Bytes(), &rec) != nil {
			continue // partial last line of a truncated segment
		}
		if rec.T >= sinceMs && rec.T < untilMs {
			fn(&rec)
		}
	}
	if err := sc.Err(); err != nil && !errors.Is(err, io.ErrUnexpectedEOF) {
		return err
	}
	return nil
}
