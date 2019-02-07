/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2019 WireGuard LLC. All Rights Reserved.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"git.zx2c4.com/wireguard-go/tun"
)

const (
	ExitSetupSuccess = 0
	ExitSetupFailed  = 1
)

func main() {
	//TODO: ensure we're running as a service

	if len(os.Args) != 2 {
		os.Exit(ExitSetupFailed)
	}
	interfaceName := os.Args[1] //TODO: infer from file instead

	logger := NewLogger(
		LogLevelDebug,
		fmt.Sprintf("(%s) ", interfaceName),
	)
	logger.Info.Println("Starting wireguard-go version", WireGuardGoVersion)
	logger.Debug.Println("Debug log enabled")

	tun, err := tun.CreateTUN(interfaceName)
	if err == nil {
		realInterfaceName, err2 := tun.Name()
		if err2 == nil {
			interfaceName = realInterfaceName
		}
	} else {
		logger.Error.Println("Failed to create TUN device:", err)
		os.Exit(ExitSetupFailed)
	}

	device := NewDevice(tun, logger)
	device.Up()
	logger.Info.Println("Device started")

	uapi, err := UAPIListen(interfaceName)
	if err != nil {
		logger.Error.Println("Failed to listen on uapi socket:", err)
		os.Exit(ExitSetupFailed)
	}

	errs := make(chan error)
	term := make(chan os.Signal, 1)

	go func() {
		for {
			conn, err := uapi.Accept()
			if err != nil {
				errs <- err
				return
			}
			go ipcHandle(device, conn)
		}
	}()
	logger.Info.Println("UAPI listener started")

	//TODO: read from encrypted wg-quick(8) file instead
	phonyConfig := `private_key=e8aa5c6cd14ae2d2817222814f6463e99fec1c1ab1755fa9fa7b8d0390255862
listen_port=0
fwmark=0
replace_peers=true
public_key=25123c5dcd3328ff645e4f2a3fce0d754400d3887a0cb7c56f0267e20fbf3c5b
endpoint=163.172.161.0:12912
replace_allowed_ips=true
allowed_ip=0.0.0.0/0`
	ipcSetOperation(device, bufio.NewReader(strings.NewReader(phonyConfig)))

	//TODO: set address,routes,dns with winipcfg module

	// wait for program to terminate

	signal.Notify(term, os.Interrupt)
	signal.Notify(term, os.Kill)
	signal.Notify(term, syscall.SIGTERM)

	select {
	case <-term:
	case <-errs:
	case <-device.Wait():
	}

	// clean up

	uapi.Close()
	device.Close()

	logger.Info.Println("Shutting down")
}
