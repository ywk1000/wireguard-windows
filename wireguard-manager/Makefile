CFLAGS ?= -O3
CFLAGS += -Wall -std=gnu11
export CGO_ENABLED := 1
export CGO_CFLAGS := $(CFLAGS)
export CC := x86_64-w64-mingw32-gcc
export GOOS := windows

wireguard-manager.exe: $(wildcard *.go *.c *.h)
	go build -ldflags="-H windowsgui" -v

run: wireguard-manager.exe
	wine wireguard-manager.exe

clean:
	rm -f wireguard-manager.exe

.PHONY: run clean
