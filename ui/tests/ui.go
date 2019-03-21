package main

import (
	"golang.zx2c4.com/wireguard/windows/service"
	"golang.zx2c4.com/wireguard/windows/ui"
)

func main() {
	service.InitializeIPCClient(nil, nil, nil)
	ui.RunUI()
}
