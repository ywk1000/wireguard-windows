/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ui

import "github.com/lxn/walk"

type TrayIcon struct {
	*walk.NotifyIcon
}

func (tray *TrayIcon) Bind() {}
