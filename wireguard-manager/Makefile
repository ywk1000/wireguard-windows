CFLAGS ?= -O3
CFLAGS += -Wall -std=gnu11
export CGO_ENABLED := 1
export CGO_CFLAGS := $(CFLAGS)
export CC := x86_64-w64-mingw32-gcc
export GOOS := windows
export PATH := $(PATH):$(GOPATH)/bin

wireguard-manager.exe: $(wildcard *.go *.c *.h) resources.syso
	go build -ldflags="-H windowsgui" -v

resources.syso: icon/icon.ico wireguard-manager.exe.manifest
	rsrc -manifest wireguard-manager.exe.manifest -ico icon/icon.ico -arch amd64 -o $@

run: wireguard-manager.exe
	wine wireguard-manager.exe

clean:
	rm -f wireguard-manager.exe resources.syso

.PHONY: run clean
