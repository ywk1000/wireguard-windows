/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ui

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/lxn/walk"
	"github.com/lxn/win"
	"golang.zx2c4.com/wireguard/windows/conf"
	"golang.zx2c4.com/wireguard/windows/service"
	"golang.zx2c4.com/wireguard/windows/ui/syntax"
)

type ManageTunnelsWindow struct {
	*walk.MainWindow

	icon *walk.Icon

	// Currently selected tunnel index in the tunnels list, or -1
	currentIndex int

	// Currently selected tunnel (not the running tunnel)
	currentTunnel *service.Tunnel
}

func NewManageTunnelsWindow(icon *walk.Icon) (*ManageTunnelsWindow, error) {
	var err error

	mtw := &ManageTunnelsWindow{
		icon: icon,
	}
	mtw.MainWindow, err = walk.NewMainWindowWithName("WireGuard")
	if err != nil {
		return nil, err
	}

	return mtw, mtw.setup()
}

func (mtw *ManageTunnelsWindow) setup() error {
	mtw.SetIcon(mtw.icon)
	mtw.SetSize(walk.Size{900, 800})
	mtw.SetLayout(walk.NewHBoxLayout())
	mtw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		// "Close to tray" instead of exiting application
		// *canceled = true
		// mtw.Hide()
		onQuit()
	})

	listBoxContainer, _ := walk.NewComposite(mtw)
	listBoxContainer.SetLayout(walk.NewVBoxLayout())

	// Left side of main window: listbox, controls

	// TODO: not greedy vertically
	walk.NewListBox(listBoxContainer)

	importAction := walk.NewAction()
	importAction.SetText("Import tunnels from file...")
	importAction.Triggered().Attach(mtw.onImport)

	addAction := walk.NewAction()
	addAction.SetText("Add empty tunnel")
	// TODO: How to tell it's a new tunnel
	addAction.Triggered().Attach(mtw.onAddTunnel)

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
	deleteButton.Clicked().Attach(mtw.onDelete)

	// TODO: Trigger the menu on standard button click
	settingsButton, _ := walk.NewSplitButton(listBoxButtonBar)
	settingsButton.SetText("S")
	settingsButton.Menu().Actions().Add(exportLogAction)
	settingsButton.Menu().Actions().Add(exportTunnelAction)

	walk.NewHSpacer(listBoxButtonBar)

	// Right side of main window: currently selected tunnel, edit

	currentTunnelContainer, _ := walk.NewComposite(mtw)
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
	editTunnel.Clicked().Attach(mtw.onEditTunnel)

	return nil
}

func (mtw *ManageTunnelsWindow) Run() int {
	return mtw.MainWindow.Run()
}

func (mtw *ManageTunnelsWindow) Show() {
	mtw.Show()
	win.SetForegroundWindow(mtw.Handle())
	win.BringWindowToTop(mtw.Handle())
}

func (mtw *ManageTunnelsWindow) setTunnelState(tunnel *service.Tunnel, state service.TunnelState) {
}

func (mtw *ManageTunnelsWindow) runTunnelEdit(tunnel *service.Tunnel) *conf.Config {
	var (
		title  string
		config conf.Config
	)

	if tunnel == nil {
		// Creating a new tunnel, create a new private key and use the default template
		title = "Create new tunnel"
		pk, _ := conf.NewPrivateKey()
		config = conf.Config{Interface: conf.Interface{PrivateKey: *pk}}
	} else {
		title = "Edit tunnel"
		config, _ = tunnel.RuntimeConfig()
	}

	dlg, _ := walk.NewDialog(mtw)
	dlg.SetIcon(mtw.icon)
	dlg.SetTitle(title)
	dlg.SetLayout(walk.NewGridLayout())
	dlg.Layout().(*walk.GridLayout).SetColumnStretchFactor(1, 3)
	dlg.Layout().SetSpacing(6)
	dlg.Layout().SetMargins(walk.Margins{18, 18, 18, 18})

	nameLabel, _ := walk.NewTextLabel(dlg)
	dlg.Layout().(*walk.GridLayout).SetRange(nameLabel, walk.Rectangle{0, 0, 1, 1})
	nameLabel.SetTextAlignment(walk.AlignHFarVCenter)
	nameLabel.SetText("Name:")

	nameEdit, _ := walk.NewLineEdit(dlg)
	dlg.Layout().(*walk.GridLayout).SetRange(nameEdit, walk.Rectangle{1, 0, 1, 1})
	// TODO: compute the next available tunnel name ?
	// nameEdit.SetText("")

	pubkeyLabel, _ := walk.NewTextLabel(dlg)
	dlg.Layout().(*walk.GridLayout).SetRange(pubkeyLabel, walk.Rectangle{0, 1, 1, 1})
	pubkeyLabel.SetTextAlignment(walk.AlignHFarVCenter)
	pubkeyLabel.SetText("Public key:")

	pubkeyEdit, _ := walk.NewLineEdit(dlg)
	dlg.Layout().(*walk.GridLayout).SetRange(pubkeyEdit, walk.Rectangle{1, 1, 1, 1})
	pubkeyEdit.SetReadOnly(true)
	pubkeyEdit.SetText("(unknown)")

	syntaxEdit, _ := syntax.NewSyntaxEdit(dlg)
	dlg.Layout().(*walk.GridLayout).SetRange(syntaxEdit, walk.Rectangle{0, 2, 2, 1})
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
	syntaxEdit.SetText(config.ToWgQuick())

	buttonsContainer, _ := walk.NewComposite(dlg)
	dlg.Layout().(*walk.GridLayout).SetRange(buttonsContainer, walk.Rectangle{0, 3, 2, 1})
	buttonsContainer.SetLayout(walk.NewHBoxLayout())
	buttonsContainer.Layout().SetMargins(walk.Margins{})

	walk.NewHSpacer(buttonsContainer)

	cancelButton, _ := walk.NewPushButton(buttonsContainer)
	cancelButton.SetText("Cancel")
	cancelButton.Clicked().Attach(dlg.Cancel)

	saveButton, _ := walk.NewPushButton(buttonsContainer)
	saveButton.SetText("Save")
	saveButton.Clicked().Attach(func() {
		// TODO: Save the current config
		dlg.Accept()
	})

	dlg.SetCancelButton(cancelButton)
	dlg.SetDefaultButton(saveButton)

	if dlg.Run() == walk.DlgCmdOK {
		// Save
	}

	return &conf.Config{}
}

// importFiles tries to import a list of configurations.
func (mtw *ManageTunnelsWindow) importFiles(paths []string) {
	type unparsedConfig struct {
		Name   string
		Config string
	}

	var (
		unparsedConfigs []unparsedConfig
		lastErr         error
	)

	// Note: other versions of WireGuard start with all .zip files, then all .conf files.
	// To reproduce that if needed, inverse-sort the array.
	for _, path := range paths {
		switch filepath.Ext(path) {
		case ".conf":
			textConfig, err := ioutil.ReadFile(path)
			if err != nil {
				lastErr = err
				continue
			}
			unparsedConfigs = append(unparsedConfigs, unparsedConfig{Name: strings.TrimSuffix(filepath.Base(path), ".conf"), Config: string(textConfig)})
		case ".zip":
			// 1 .conf + 1 error .zip edge case?
			r, err := zip.OpenReader(path)
			if err != nil {
				lastErr = err
				continue
			}

			for _, f := range r.File {
				if filepath.Ext(f.Name) != ".conf" {
					continue
				}

				rc, err := f.Open()
				if err != nil {
					lastErr = err
					continue
				}
				textConfig, err := ioutil.ReadAll(rc)
				rc.Close()
				if err != nil {
					lastErr = err
					continue
				}
				unparsedConfigs = append(unparsedConfigs, unparsedConfig{Name: strings.TrimSuffix(filepath.Base(f.Name), ".conf"), Config: string(textConfig)})
			}

			r.Close()
		}
	}

	if lastErr != nil || unparsedConfigs == nil {
		walk.MsgBox(mtw, "Error", fmt.Sprintf("Could not parse some files: %v", lastErr), walk.MsgBoxIconWarning)
		return
	}

	var (
		configs []*conf.Config
	)

	for _, unparsedConfig := range unparsedConfigs {
		config, err := conf.FromWgQuick(unparsedConfig.Config, unparsedConfig.Name)
		if err != nil {
			lastErr = err
			continue
		}
		configs = append(configs, config)
	}

	m, n := len(configs), len(unparsedConfigs)
	switch {
	case n == 1 && m != n:
		walk.MsgBox(mtw, "Error", fmt.Sprintf("Could not parse some files: %v", lastErr), walk.MsgBoxIconWarning)
	case n == 1 && m == n:
		// Select tunnel
	case m == n:
		walk.MsgBox(mtw, "Imported tunnels", fmt.Sprintf("Imported %d tunnels", m), walk.MsgBoxOK)
	case m != n:
		walk.MsgBox(mtw, "Imported tunnels", fmt.Sprintf("Imported %d of %d tunnels", m, n), walk.MsgBoxIconWarning)
	default:
		panic("unreachable case")
	}
}

// Handlers

func (mtw *ManageTunnelsWindow) onEditTunnel() {
	oldName := mtw.currentTunnel.Name
	if config := mtw.runTunnelEdit(mtw.currentTunnel); config != nil {
		// TODO: is there a tunnel rename call?
		if oldName != config.Name {
			// Delete old one
			// mtw.currentTunnel.Delete()
		}
		// Save new one
		// config.Create()
		// Update the currentTunnel to use the new config (careful to keep ordering)
	}
}

func (mtw *ManageTunnelsWindow) onAddTunnel() {
	if config := mtw.runTunnelEdit(nil); config != nil {
		// Save new
	}
}

func (mtw *ManageTunnelsWindow) onDelete() {
	if mtw.currentTunnel == nil {
		return
	}

	if walk.MsgBox(mtw, fmt.Sprintf(`Delete "%s"?`, mtw.currentTunnel.Name), fmt.Sprintf(`Are you sure you want to delete "%s"`, "tunnel name"), walk.MsgBoxYesNo|walk.MsgBoxIconWarning) != walk.DlgCmdOK {
		return
	}

	mtw.currentTunnel.Delete()
}

func (mtw *ManageTunnelsWindow) onImport() {
	dlg := &walk.FileDialog{}
	// dlg.InitialDirPath
	dlg.Filter = "Configuration Files (*.zip, *.conf)|*.zip;*.conf|All Files (*.*)|*.*"
	dlg.Title = "Import tunnel(s) from file..."

	if ok, _ := dlg.ShowOpenMultiple(mtw); !ok {
		return
	}

	mtw.importFiles(dlg.FilePaths)
}
