/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/lxn/walk"
	"github.com/lxn/win"
	"golang.zx2c4.com/wireguard/windows/conf"
	"golang.zx2c4.com/wireguard/windows/service"
	"golang.zx2c4.com/wireguard/windows/ui/syntax"
)

var (
	mw   *walk.MainWindow
	tray *walk.NotifyIcon
	icon *walk.Icon
	// TODO: Only one running tunnel at most atm. Plan for more.
	runningTunnel *service.Tunnel
)

func RunUI() {
	mw, _ = walk.NewMainWindowWithName("WireGuard")

	tray, _ = walk.NewNotifyIcon(mw)
	defer tray.Dispose()

	icon, _ = walk.NewIconFromResourceId(8)
	defer icon.Dispose()

	setupTray()
	setupTunnelList()
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
			// updateConfView()
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

func setupTunnelList() {
	mw.SetSize(walk.Size{900, 800})
	mw.SetLayout(walk.NewHBoxLayout())
	mw.SetIcon(icon)
	mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		// "Close to tray" instead of exiting application
		// *canceled = true
		// mw.Hide()
		onQuit()
	})

	listBoxContainer, _ := walk.NewComposite(mw)
	listBoxContainer.SetLayout(walk.NewVBoxLayout())

	// Left side of main window: listbox, controls

	// TODO: not greedy vertically
	walk.NewListBox(listBoxContainer)

	importAction := walk.NewAction()
	importAction.SetText("Import tunnels from file...")
	importAction.Triggered().Attach(onImport)

	addAction := walk.NewAction()
	addAction.SetText("Add empty tunnel")
	// TODO: How to tell it's a new tunnel
	addAction.Triggered().Attach(onEdit)

	exportLogAction := walk.NewAction()
	exportLogAction.SetText("Export log to file...")
	// TODO: Triggered().Attach()

	exportTunnelAction := walk.NewAction()
	exportTunnelAction.SetText("Export tunnels to zip...")
	// TODO: Triggered().Attach()

	// TODO: Maybe a Rebar looks better
	listBoxButtonBar, _ := walk.NewComposite(listBoxContainer)
	listBoxButtonBar.SetLayout(walk.NewHBoxLayout())
	listBoxButtonBar.Layout().SetMargins(walk.Margins{})
	listBoxButtonBar.Layout().SetSpacing(0)

	// TODO: Trigger the menu on standard button click
	addButton, _ := walk.NewSplitButton(listBoxButtonBar)
	addButton.SetText("+")
	addButton.Menu().Actions().Add(addAction)
	addButton.Menu().Actions().Add(importAction)

	deleteButton, _ := walk.NewPushButton(listBoxButtonBar)
	deleteButton.SetText("-")
	deleteButton.Clicked().Attach(onDelete)

	// TODO: Trigger the menu on standard button click
	settingsButton, _ := walk.NewSplitButton(listBoxButtonBar)
	settingsButton.SetText("S")
	settingsButton.Menu().Actions().Add(exportLogAction)
	settingsButton.Menu().Actions().Add(exportTunnelAction)

	walk.NewHSpacer(listBoxButtonBar)

	// Right side of main window: currently selected tunnel, edit

	currentTunnelContainer, _ := walk.NewComposite(mw)
	currentTunnelContainer.SetLayout(walk.NewVBoxLayout())

	// TODO: Replace with ConfView
	currentTunnel, _ := walk.NewTextEdit(currentTunnelContainer)
	currentTunnel.SetReadOnly(true)

	controlsContainer, _ := walk.NewComposite(currentTunnelContainer)
	controlsContainer.SetLayout(walk.NewHBoxLayout())
	controlsContainer.Layout().SetMargins(walk.Margins{})

	toggleTunnel, _ := walk.NewCheckBox(controlsContainer)
	toggleTunnel.SetText("Status: deactivated")

	walk.NewHSpacer(controlsContainer)

	editTunnel, _ := walk.NewPushButton(controlsContainer)
	editTunnel.SetText("Edit")
	editTunnel.Clicked().Attach(onEdit)
}

func onDelete() {
	// result is either walk.IDNO or walk.IDYES
	walk.MsgBox(mw, fmt.Sprintf(`Delete "%s"?`, "tunnel name"), fmt.Sprintf(`Are you sure you want to delete "%s"`, "tunnel name"), walk.MsgBoxYesNo|walk.MsgBoxIconWarning)
}

func onEdit() {
	// result is either walk.DlgCmdOK or walk.DlgCmdCancel
	getTunnelEdit().Run()
}

// Bind service events to the GUI.
// The tray icon tooltip is defined by the active tunnel (at most one).
//
func bindService() {
	setServiceState := func(tunnel *service.Tunnel, state service.TunnelState, showNotifications bool) {
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
				runningTunnel.Stop()
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

func getTunnelEdit() *walk.Dialog {
	dlg, _ := walk.NewDialog(mw)

	// TODO: Size does not seem to apply
	dlg.SetSize(walk.Size{900, 800})
	dlg.SetLayout(walk.NewVBoxLayout())
	dlg.SetIcon(icon)
	dlg.SetTitle("Edit tunnel")

	nameContainer, _ := walk.NewComposite(dlg)
	nameContainer.SetLayout(walk.NewHBoxLayout())

	nameLabel, _ := walk.NewTextLabel(nameContainer)
	nameLabel.SetText("Name:")

	nameEdit, _ := walk.NewLineEdit(nameContainer)
	_ = nameEdit
	// TODO: compute the next available tunnel name ?
	// nameEdit.SetText("")

	pubkeyContainer, _ := walk.NewComposite(dlg)
	pubkeyContainer.SetLayout(walk.NewHBoxLayout())

	pubkeyLabel, _ := walk.NewTextLabel(pubkeyContainer)
	pubkeyLabel.SetText("Public key:")

	pubkeyEdit, _ := walk.NewLineEdit(pubkeyContainer)
	pubkeyEdit.SetReadOnly(true)
	pubkeyEdit.SetText("(unknown)")

	syntaxEdit, _ := syntax.NewSyntaxEdit(dlg)
	lastPrivate := ""
	syntaxEdit.PrivateKeyChanged().Attach(func(privateKey string) {
		if privateKey == lastPrivate {
			return
		}
		lastPrivate = privateKey
		key, _ := conf.NewPrivateKeyFromString(privateKey)
		if key != nil {
			pubkeyEdit.SetText(key.Public().String())
		} else {
			pubkeyEdit.SetText("(unknown)")
		}
	})

	// Generate new private key
	pk, _ := conf.NewPrivateKey()
	newConfig := &conf.Config{Interface: conf.Interface{PrivateKey: *pk}}
	syntaxEdit.SetText(newConfig.ToWgQuick())

	buttonsContainer, _ := walk.NewComposite(dlg)
	buttonsContainer.SetLayout(walk.NewHBoxLayout())

	walk.NewHSpacer(buttonsContainer)

	cancelButton, _ := walk.NewPushButton(buttonsContainer)
	cancelButton.SetText("Cancel")
	cancelButton.Clicked().Attach(func() { dlg.Cancel() })

	saveButton, _ := walk.NewPushButton(buttonsContainer)
	saveButton.SetText("Save")
	saveButton.Clicked().Attach(func() {
		// TODO: Save the current config
		dlg.Accept()
	})

	dlg.SetCancelButton(cancelButton)
	dlg.SetDefaultButton(saveButton)

	return dlg
}
