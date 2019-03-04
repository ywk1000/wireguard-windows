/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ui

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/lxn/walk"
	"github.com/lxn/win"
	"golang.org/x/crypto/curve25519"
	"golang.zx2c4.com/wireguard/windows/conf"
	"golang.zx2c4.com/wireguard/windows/service"
	"golang.zx2c4.com/wireguard/windows/ui/syntax"
)

const demoConfig = `[Interface]
PrivateKey = 6KpcbNFK4tKBciKBT2Rj6Z/sHBqxdV+p+nuNA5AlWGI=
Address = 192.168.4.84/24
DNS = 8.8.8.8, 8.8.4.4, 1.1.1.1, 1.0.0.1

[Peer]
PublicKey = JRI8Xc0zKP9kXk8qP84NdUQA04h6DLfFbwJn4g+/PFs=
Endpoint = demo.wireguard.com:12912
AllowedIPs = 0.0.0.0/0
`

var (
	mw            *walk.MainWindow
	tray          *walk.NotifyIcon
	icon          *walk.Icon
	runningTunnel *service.Tunnel
)

func RunUI() {
	mw, _ = walk.NewMainWindowWithName("WireGuard")

	tray, _ = walk.NewNotifyIcon(mw)
	defer tray.Dispose()

	icon, _ = walk.NewIconFromResourceId(8)
	defer icon.Dispose()

	setupTray()
	setupTunnelsList()
	bindService()

	mw.Run()
}

func setupTray() {
	tray.SetToolTip("WireGuard: Deactivated")
	tray.SetVisible(true)
	tray.SetIcon(icon)

	tray.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button == walk.LeftButton {
			mw.Show()
			win.SetForegroundWindow(mw.Handle())
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
		{label: "&Manage tunnels...", handler: onManage, enabled: true},
		{label: "&Import tunnel(s) from file...", handler: onImport, enabled: true},
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
}

// Handlers

func onAbout() {
	walk.MsgBox(mw, "About WireGuard", "TODO", walk.MsgBoxOK)
}

func onQuit() {
	tray.Dispose()
	_, err := service.IPCClientQuit(true)
	if err != nil {
		walk.MsgBox(nil, "Error Exiting WireGuard", fmt.Sprintf("Unable to exit service due to: %s. You may want to stop WireGuard from the service manager.", err), walk.MsgBoxIconError)
		os.Exit(1)
	}
}

func onManage() {
	mw.Show()
	win.SetForegroundWindow(mw.Handle())
}

func onImport() {
	dlg := &walk.FileDialog{}
	// dlg.InitialDirPath
	dlg.Filter = "Configuration Files (*.zip, *.conf)|*.zip;*.conf|All Files (*.*)|*.*"
	dlg.Title = "Import tunnel(s) from file..."

	if ok, _ := dlg.ShowOpenMultiple(mw); !ok {
		return
	}

	walk.MsgBox(mw, "Must import", strings.Join(dlg.FilePaths, ", "), walk.MsgBoxOK)
}

func setupTunnelsList() {
	mw.SetSize(walk.Size{900, 800})
	mw.SetLayout(walk.NewHBoxLayout())
	mw.SetIcon(icon)
	mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		// "Close to tray" instead of exiting application
		// *canceled = true
		// mw.Hide()
	})

	listBoxContainer, _ := walk.NewComposite(mw)
	listBoxContainer.SetLayout(walk.NewVBoxLayout())

	// Left side of main window: listbox, controls

	// TODO: not greedy vertically
	walk.NewListBox(listBoxContainer)
	// TODO: Use a Rebar on the ListBox to tie them together
	controls := mw.ToolBar()

	importAction := walk.NewAction()
	importAction.SetText("Import tunnels from file...")
	importAction.Triggered().Attach(onImport)

	addAction := walk.NewAction()
	addAction.SetText("Add empty tunnel")
	// TODO: How to tell it's a new tunnel
	addAction.Triggered().Attach(onEdit)

	deleteAction := walk.NewAction()
	deleteAction.SetText("-")
	deleteAction.Triggered().Attach(onDelete)

	exportLogAction := walk.NewAction()
	exportLogAction.SetText("Export log to file...")

	exportTunnelAction := walk.NewAction()
	exportTunnelAction.SetText("Export tunnels to zip...")

	addMenu, _ := walk.NewMenu()
	addMenu.Actions().Add(addAction)
	addMenu.Actions().Add(importAction)

	addMenuAction, _ := controls.Actions().AddMenu(addMenu)
	addMenuAction.SetText("+")

	settingsMenu, _ := walk.NewMenu()
	settingsMenu.Actions().Add(exportLogAction)
	settingsMenu.Actions().Add(exportTunnelAction)

	settingsMenuAction, _ := controls.Actions().AddMenu(settingsMenu)
	settingsMenuAction.SetText("S")

	// Right side of main window: currently selected tunnel, edit

	currentTunnelContainer, _ := walk.NewComposite(mw)
	currentTunnelContainer.SetLayout(walk.NewVBoxLayout())

	currentTunnel, _ := walk.NewTextEdit(currentTunnelContainer)
	currentTunnel.SetReadOnly(true)

	editTunnel, _ := walk.NewPushButton(currentTunnelContainer)
	editTunnel.SetText("Edit")
	editTunnel.Clicked().Attach(onEdit)
}

func onDelete() {
}

func onEdit() {
	dlg, _ := walk.NewDialog(mw)

	// TODO: Size does not seem to apply
	dlg.SetSize(walk.Size{900, 800})
	dlg.SetLayout(walk.NewVBoxLayout())
	dlg.SetIcon(icon)
	dlg.SetTitle("Edit tunnel")

	tl, _ := walk.NewTextLabel(dlg)
	tl.SetText("Public key: (unknown)")

	se, _ := syntax.NewSyntaxEdit(dlg)
	lastPrivate := ""
	se.PrivateKeyChanged().Attach(func(privateKey string) {
		if privateKey == lastPrivate {
			return
		}
		lastPrivate = privateKey
		key := func() string {
			if privateKey == "" {
				return ""
			}
			decoded, err := base64.StdEncoding.DecodeString(privateKey)
			if err != nil {
				return ""
			}
			if len(decoded) != 32 {
				return ""
			}
			var p [32]byte
			var s [32]byte
			copy(s[:], decoded[:32])
			curve25519.ScalarBaseMult(&p, &s)
			return base64.StdEncoding.EncodeToString(p[:])
		}()
		if key != "" {
			tl.SetText("Public key: " + key)
		} else {
			tl.SetText("Public key: (unknown)")
		}
	})
	se.SetText(demoConfig)

	pb, _ := walk.NewPushButton(dlg)
	pb.SetText("Start")
	pb.Clicked().Attach(func() {
		restoreState := true
		pbE := pb.Enabled()
		seE := se.Enabled()
		pbT := pb.Text()
		defer func() {
			if restoreState {
				pb.SetEnabled(pbE)
				se.SetEnabled(seE)
				pb.SetText(pbT)
			}
		}()
		pb.SetEnabled(false)
		se.SetEnabled(false)
		pb.SetText("Requesting..")
		if runningTunnel != nil {
			err := runningTunnel.Stop()
			if err != nil {
				walk.MsgBox(mw, "Unable to stop tunnel", err.Error(), walk.MsgBoxIconError)
				return
			}
			restoreState = false
			runningTunnel = nil
			return
		}
		c, err := conf.FromWgQuick(se.Text(), "test")
		if err != nil {
			walk.MsgBox(mw, "Invalid configuration", err.Error(), walk.MsgBoxIconError)
			return
		}
		tunnel, err := service.IPCClientNewTunnel(c)
		if err != nil {
			walk.MsgBox(mw, "Unable to create tunnel", err.Error(), walk.MsgBoxIconError)
			return
		}
		err = tunnel.Start()
		if err != nil {
			walk.MsgBox(mw, "Unable to start tunnel", err.Error(), walk.MsgBoxIconError)
			return
		}
		restoreState = false
		runningTunnel = &tunnel
	})

	dlg.Run()
}

func bindService() {

	setServiceState := func(tunnel *service.Tunnel, state service.TunnelState, showNotifications bool) {
		if tunnel.Name != "test" {
			return
		}
		//TODO: also set tray icon to reflect state
		switch state {
		case service.TunnelStarting:
			// se.SetEnabled(false)
			// pb.SetText("Starting...")
			// pb.SetEnabled(false)
			tray.SetToolTip("WireGuard: Activating...")
		case service.TunnelStarted:
			// se.SetEnabled(false)
			// pb.SetText("Stop")
			// pb.SetEnabled(true)
			tray.SetToolTip("WireGuard: Activated")
			if showNotifications {
				//TODO: ShowCustom with right icon
				tray.ShowInfo("WireGuard Activated", fmt.Sprintf("The %s tunnel has been activated.", tunnel.Name))
			}
		case service.TunnelStopping:
			// se.SetEnabled(false)
			// pb.SetText("Stopping...")
			// pb.SetEnabled(false)
			tray.SetToolTip("WireGuard: Deactivating...")
		case service.TunnelStopped, service.TunnelDeleting:
			if runningTunnel != nil {
				runningTunnel.Delete()
				runningTunnel = nil
			}
			// se.SetEnabled(true)
			// pb.SetText("Start")
			// pb.SetEnabled(true)
			tray.SetToolTip("WireGuard: Deactivated")
			if showNotifications {
				//TODO: ShowCustom with right icon
				tray.ShowInfo("WireGuard Deactivated", fmt.Sprintf("The %s tunnel has been deactivated.", tunnel.Name))
			}
		}
	}

	service.IPCClientRegisterTunnelChange(func(tunnel *service.Tunnel, state service.TunnelState) {
		setServiceState(tunnel, state, true)
	})
	go func() {
		tunnels, err := service.IPCClientTunnels()
		if err != nil {
			return
		}
		for _, tunnel := range tunnels {
			state, err := tunnel.State()
			if err != nil {
				continue
			}
			runningTunnel = &tunnel
			setServiceState(&tunnel, state, false)
		}
	}()
}
