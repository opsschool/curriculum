# blackbox

A flight recorder for small Linux systems, written for a UniFi Gateway Fiber
but usable on any Linux box with procfs. It runs continuously, keeps a bounded
compressed history on disk, and when load or I/O pressure goes over a
threshold it records which threads are stuck waiting on the disk, where they
are waiting in the kernel, and which file they are working on.

It reads `/proc` only. It needs no eBPF, perf, kernel modules, or libraries,
and it ships as a single static binary.

## Overhead

Measured on a 4-CPU x86 VM during an I/O incident, with stack capture on:

| | |
|---|---|
| RSS | 4–6 MB (Go heap soft-limited to 12 MB, systemd cap 48 MB) |
| CPU | ~0.1% of one core idle, ~0.5% during an incident |
| Disk | a few MB per day compressed; directory capped at 64 MB |
| Binary | 2.4 MB, static |

A slower ARM core costs more CPU per `/proc` read, so expect a few times
these CPU figures on the gateway. Check the real cost with `blackbox show`:
every sample records the recorder's own RSS and CPU (`self_rss` and
`self_cpu` in `export`).

The sampling path never waits on the disk. Records go through a queue to a
separate writer, and the writer never calls fsync. If the disk stalls long
enough to fill the queue, records are dropped and the number lost is noted
in the next sample.

## What it records

Every 5 s (every 2 s during an incident):

- load average, run queue, `procs_blocked`
- CPU split (user, system, iowait, irq, softirq, steal, idle), context
  switches, forks
- memory: available, cached, dirty, writeback, shmem, slab, swap used
- major faults, page-in/out, swap-in/out, OOM kills
- PSI stall percentages for cpu, memory and io, if the kernel has PSI
- per disk: IOPS, throughput, utilisation, average wait, queue depth
- per active process: CPU %, RSS, threads, major faults, bytes read and
  written to storage, cancelled writes, `rchar`/`wchar`, and block-I/O delay
  (time spent waiting on the disk)

Every minute it takes a full snapshot of every process, including idle ones,
with command line and swap use. Between snapshots only processes that did
something are recorded, up to 40 per sample.

Every second it checks load, PSI and `procs_blocked`. When anything is
blocked it scans for threads in D state (uninterruptible sleep) and records,
for each one:

- process and thread name
- `wchan`: the kernel function it is sleeping in
- the syscall it is in, and the file behind the syscall's fd (for
  read/write/fsync/fdatasync/pwrite and the like). Here "pagefault" means
  the thread was blocked outside any syscall, which almost always means it
  was waiting for a page fault to be served from storage.
- during an incident, its kernel stack

An **incident** starts when any of these is true:

- load1 ≥ number of CPUs (`-load`)
- PSI io "some" avg10 ≥ 20% (`-psi-io`)
- `procs_blocked` ≥ 2 on three probes in a row (`-blocked`)

It ends after 30 s with no trigger (`-hold`). During the first 10 minutes of
an incident (`-max-detail`), blackbox scans D-state threads every 500 ms and
captures their stacks. After that it falls back to the normal rate, so a
long incident does not raise the overhead indefinitely.

## Build

You need Go 1.22 or newer on any machine. Nothing is downloaded at build
time.

```sh
cd tools/blackbox
make arm64        # dist/blackbox-linux-arm64 (UniFi Gateway Fiber and other current UniFi OS consoles)
make all          # also armv7 and amd64
make test
```

## Install on the gateway

Enable SSH in UniFi OS, then:

```sh
make arm64
ssh root@GATEWAY mkdir -p /tmp/bb
scp dist/blackbox-linux-arm64 blackbox.service install.sh root@GATEWAY:/tmp/bb/
ssh root@GATEWAY sh /tmp/bb/install.sh
```

This installs the binary and its data to `/data/blackbox` and starts a
systemd service. `/data` survives reboots and firmware updates, but a
firmware update can reset `/etc/systemd/system`. After an update, run
`sh /data/blackbox/install.sh` again, or simply rerun the steps above.

`install.sh` finishes by running `blackbox check`, which reports which kernel
features are present:

```
io_accounting   yes  per-process bytes read/written
task_delayacct  1    per-process time spent waiting on block I/O (BLKIO column)
psi             yes  pressure stall information
stack           yes  kernel stacks of blocked threads
syscall         yes  syscall a blocked thread is in, used to find the file
wchan           yes  kernel function a blocked thread is waiting in
```

A missing feature only removes the matching column or field; the rest still
works. The unit passes `-delayacct`, which sets `kernel.task_delayacct=1` on
kernels that have that switch (5.14 and later). On older kernels, delay
accounting is on whenever it is compiled in.

To record without touching the flash at all, run with `-dir /run/blackbox`
(tmpfs). The history then costs RAM and is lost on reboot.

## Finding the source of disk-wait load spikes

Linux counts threads in D state in the load average, so disk stalls show up
as load even when the CPUs are idle. The questions are which threads are
blocked, and which process is generating the I/O they are waiting behind.
Often these are different processes.

**1. List the incidents:**

```sh
blackbox incidents -list
blackbox incidents            # summary of the last 10
```

Each summary shows the window, peak load, iowait, PSI, memory, per-disk
throughput and utilisation, then:

- **Processes by stall**: per process, `BLKIO s` (time its threads spent
  waiting on block I/O, from delay accounting), `D thr·s` (thread-seconds
  observed in D state), bytes read and written to storage, major faults,
  CPU, and peak RSS.
- **D-state breakdown**: by thread, by kernel wait point, and by
  syscall + file, plus the most common kernel stack.
- **Notes**: conclusions drawn from the numbers, such as page-cache
  thrashing, journal (jbd2) contention, dirty-page throttling, or one
  process doing most of the writing.

**2. Narrow to a window:**

```sh
blackbox top -since "2026-10-08 14:00" -until "14:20" -by write
blackbox blocked -since 2h -stacks 5
blackbox show -since 14:00 -until 14:10 -procs 5 -blocked
```

`show` prints one line per sample:

```
TIME            LOAD1  RUN  BLK   D   USR  SYS  IOW  IDLE  PSIio  AVAIL  DIRTY   SWAP  MAJF/s  DISKS
10-08 13:05:34   0.19    2    3   4     1   16   65    17   70.2  15.3G   5.7M      0       0  vda 91% r2.0K w287.8M
```

`D` is the largest number of D-state threads seen since the previous line.

**3. Look live:**

```sh
blackbox now          # 1-second measurement: top processes, every D-state thread with its stack
```

### Reading the results

- **One process has high WRITE and other processes are blocked in fsync, or
  in `jbd2_log_wait_commit`.** The writer fills the ext4 journal or the
  page cache, and everyone who fsyncs waits behind it. The syscall + file
  section shows exactly which files are being synced.
- **The blocked syscall is "pagefault", major faults are high, and
  MemAvailable is low.** The system is short of memory. Code and mmapped
  files are evicted and re-read from flash, so the disk is busy with reads
  nobody asked for explicitly. The fix is memory, not I/O: find the largest
  RSS and swap users in `top -by rss`.
- **Threads wait in `balance_dirty_pages`.** Dirty data is produced faster
  than the flash can write it back, and the kernel is throttling the
  writers. The top writer is the cause.
- **The blocked threads are `kworker/...+flush-*` or `jbd2/*`.** Those are
  the kernel writing back data. Writeback is charged to the process that
  dirtied the pages, so use `top -by write` to find whose data it is.

## Data and offline analysis

Data lives in `/data/blackbox/data/bb-<UTC time>.jsonl.gz`. Each file is a
gzip stream of JSON lines; a new file starts every hour or 4 MB. Files cut
short by a crash or power loss are read up to the last flush (every 30 s,
and on every sample during an incident).

```sh
blackbox export -since 24h > dump.jsonl                   # everything
blackbox export -kinds blk,inc | jq -c 'select(.k=="blk")'
zcat /data/blackbox/data/*.gz | head                     # files are plain gzip
```

The record schema is the `Record` type in `record.go`. Process counters are
amounts during the record's `dt` seconds, not rates.

## Limitations

- **D-state is sampled.** Stalls shorter than the probe interval (1 s, or
  500 ms during an incident) are seen statistically, not exactly.
  `BLKIO` from delay accounting is exact; use it to rank processes.
- **Short-lived processes.** A process that starts and exits between two
  samples is never seen. The kernel adds a reaped child's I/O counters to
  its parent, so that I/O shows up on the parent, e.g. a shell or cron job
  that ran `dd`. For children that were seen, blackbox subtracts what it
  had already counted, so nothing is reported twice.
- **Threads that exit** take their unsampled block-I/O delay with them.
- **Time**: records use wall-clock time. If the gateway boots without a
  valid clock and NTP corrects it later, timestamps jump.
- **Syscall names** are built in for arm64 and x86_64. On other
  architectures syscalls are shown by number.

## Recorder flags

`blackbox record -h` lists them all. The main ones:

| flag | default | |
|---|---|---|
| `-dir` | `/data/blackbox/data` | data directory |
| `-interval` | 5s | sample interval |
| `-probe` | 1s | load / blocked-thread check interval |
| `-load` | #CPUs | load1 incident threshold |
| `-psi-io` | 20 | PSI io avg10 % threshold (0 = off) |
| `-blocked` | 2 | procs_blocked threshold |
| `-max-mb` | 64 | data directory size cap |
| `-delayacct` | off | turn on `kernel.task_delayacct` |
| `-mlock` | off (on in the unit) | lock the program code in RAM (2.5 MB) so memory pressure cannot evict it |
