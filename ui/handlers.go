/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ui

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	"golang.zx2c4.com/wireguard/windows/service"
)

const aboutText = `
WireGuard Windows.
TODO.

Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
`

func onAbout() {
	walk.MsgBox(nil, "About WireGuard", aboutText, walk.MsgBoxOK)
}

func onQuit() {
	_, err := service.IPCClientQuit(true)
	if err != nil {
		walk.MsgBox(nil, "Error Exiting WireGuard", fmt.Sprintf("Unable to exit service due to: %s. You may want to stop WireGuard from the service manager.", err), walk.MsgBoxIconError)
		os.Exit(1)
	}

	walk.App().Exit(0)
}
