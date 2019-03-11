/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ui

import (
	"fmt"

	"github.com/lxn/walk"
	"golang.zx2c4.com/wireguard/windows/service"
)

type Tray struct {
	*walk.NotifyIcon

	parent *ManageTunnelsWindow
	icon   *walk.Icon
}

func NewTray(parent *ManageTunnelsWindow, icon *walk.Icon) (*Tray, error) {
	var err error

	tray := &Tray{
		parent: parent,
		icon:   icon,
	}
	tray.NotifyIcon, err = walk.NewNotifyIcon(parent.MainWindow)
	if err != nil {
		return nil, err
	}

	return tray, tray.setup()
}

func (tray *Tray) setup() error {
	tray.SetToolTip("WireGuard: Deactivated")
	tray.SetVisible(true)
	tray.SetIcon(tray.icon)

	tray.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button == walk.LeftButton {
			tray.parent.Show()
		}
	})

	// configure initial menu items
	for _, item := range [...]struct {
		label     string
		handler   walk.EventHandler
		enabled   bool
		separator bool
	}{
		{label: "Status: unknown"},
		// TODO: Currently enabled tunnels CIDRs
		{separator: true},
		// TODO: Tunnels go here
		{separator: true},
		{label: "&Manage tunnels...", handler: tray.parent.Show, enabled: true},
		{label: "&Import tunnel(s) from file...", handler: tray.parent.onImport, enabled: true},
		{separator: true},
		{label: "&About WireGuard", handler: onAbout, enabled: true},
		{label: "&Quit", handler: onQuit, enabled: true},
	} {
		var action *walk.Action
		if item.separator {
			action = walk.NewSeparatorAction()
		} else {
			action = walk.NewAction()
			action.SetText(item.label)
			action.SetEnabled(item.enabled)
			if item.handler != nil {
				action.Triggered().Attach(item.handler)
			}
		}

		tray.ContextMenu().Actions().Add(action)
	}

	return nil
}

func (tray *Tray) setTunnelState(tunnel *service.Tunnel, state service.TunnelState) {
	tray.setTunnelStateWithNotification(tunnel, state, true)
}

func (tray *Tray) setTunnelStateWithNotification(tunnel *service.Tunnel, state service.TunnelState, showNotifications bool) {
	//TODO: also set tray icon to reflect state
	switch state {
	case service.TunnelStarting:
		tray.SetToolTip("WireGuard: Activating...")
	case service.TunnelStarted:
		tray.SetToolTip("WireGuard: Activated")
		if showNotifications {
			tray.ShowInfo("WireGuard Activated", fmt.Sprintf("The %s tunnel has been activated.", tunnel.Name))
		}
	case service.TunnelStopping:
		tray.SetToolTip("WireGuard: Deactivating...")
	case service.TunnelStopped:
		tray.SetToolTip("WireGuard: Deactivated")
		if showNotifications {
			tray.ShowInfo("WireGuard Deactivated", fmt.Sprintf("The %s tunnel has been deactivated.", tunnel.Name))
		}
	}
}
