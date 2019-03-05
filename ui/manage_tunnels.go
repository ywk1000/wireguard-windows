/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ui

import (
	"github.com/lxn/walk"
	"golang.zx2c4.com/wireguard/windows/service"
)

type ManageTunnelsWindow struct {
	*walk.MainWindow

	// Currently selected tunnel index in the tunnels list, or -1
	currentIndex int

	// Currently selected tunnel
	currentTunnel *service.Tunnel
}
