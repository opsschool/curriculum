# Static binaries, no cgo. The UniFi Gateway Fiber and other current UniFi OS
# consoles are arm64; arm (v7) is for older 32-bit devices.
GOFLAGS := -trimpath
LDFLAGS := -s -w

.PHONY: all arm64 arm amd64 test clean
all: arm64 arm amd64

arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(GOFLAGS) -ldflags='$(LDFLAGS)' -o dist/blackbox-linux-arm64 .
arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build $(GOFLAGS) -ldflags='$(LDFLAGS)' -o dist/blackbox-linux-armv7 .
amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -ldflags='$(LDFLAGS)' -o dist/blackbox-linux-amd64 .

test:
	go vet ./...
	go test ./...

clean:
	rm -rf dist
