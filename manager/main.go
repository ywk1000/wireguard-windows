/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package main

import (
	"log"
	"time"

	"git.zx2c4.com/wireguard-windows/manager/conf"
	"git.zx2c4.com/wireguard-windows/manager/walk"
	. "git.zx2c4.com/wireguard-windows/manager/walk/declarative"
	"git.zx2c4.com/wireguard-windows/manager/walk/win"
)

type MainWindowModel struct {
	MainWindow   *walk.MainWindow
	model        *InterfacesModel
	lb           *walk.ListBox
	pv           *ConfView
	refreshtimer *time.Timer
}

var wgicon *walk.Icon

func main() {

	mwm := &MainWindowModel{model: &InterfacesModel{}}

	demo_config := `[Interface]
PrivateKey = cJM9wXUVXc0fa/t5b/Lm0BHNx6jh5UiTLsO+oJhyQUU=
Address = 192.168.4.84/24, 2001:abcd:33::/120 # this is a comment
DNS = 8.8.8.8, 8.8.4.4, 1.1.1.1, 1.0.0.1

[Peer]
PublicKey = JRI8Xc0zKP9kXk8qP84NdUQA04h6DLfFbwJn4g+/PFs=
Endpoint = demo.wireguard.com:12912
AllowedIPs = 0.0.0.0/0

[Peer]
PublicKey = QCssGR6joqOIEQW6j2AR7nMcXJIVI9E9PCcbwrVXhU8=
Endpoint = intranet.wireguard.com:51820
AllowedIPs = 192.168.22.0/24, fd00:3001::/64
`

	tunnel, _ := conf.FromWgQuick(demo_config, "demo")
	mwm.model.items = append(mwm.model.items, tunnel)

	mw, err := walk.NewMainWindow() //just for the tray icon msg loop, never shown
	if err != nil {
		log.Fatal(err)
	}

	wgicon, err = walk.NewIconFromResourceId(8) //TODO: will this stay at id=8??
	if err != nil {
		log.Fatal(err)
	}

	ni, err := walk.NewNotifyIcon()
	if err != nil {
		log.Fatal(err)
	}
	defer ni.Dispose()

	if err := ni.SetIcon(wgicon); err != nil {
		log.Fatal(err)
	}
	if err := ni.SetToolTip("WireGuard"); err != nil {
		log.Fatal(err)
	}

	menus := []MenuItem{
		Action{
			Text:    "Status: Inactive",
			Enabled: false,
		},
		Action{
			Text:    "Networks: 0.0.0.0/0",
			Enabled: false,
		},
		Separator{},
		Action{
			Text: "Manage tunnels",
			OnTriggered: func() {
				mwm.runManagerWindow()
			},
		},
		Action{
			Text: "Import tunnel(s) from file...",
			OnTriggered: func() {

			},
		},
		Separator{},
		Action{
			Text: "About WireGuard",
			OnTriggered: func() {
				runAboutWindow(mw)
			},
		},
		Action{
			Text: "Quit",
			OnTriggered: func() {
				if mwm.MainWindow != nil {
					//todo: close window/dialogs
				}
				walk.App().Exit(0)
			},
		},
	}
	addMenus(ni, menus)

	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}
	mw.Run()

}

func addMenus(ni *walk.NotifyIcon, menus []MenuItem) error {
	for _, mi := range menus {
		var action *walk.Action
		switch mi.(type) {
		case Action:
			action = walk.NewAction()
			if err := action.SetText(mi.(Action).Text); err != nil {
				return err
			}
			if mi.(Action).Enabled != nil {
				action.SetEnabled(mi.(Action).Enabled.(bool))
			}
			action.Triggered().Attach(mi.(Action).OnTriggered)
		case Separator:
			action = walk.NewSeparatorAction()
		}
		if err := ni.ContextMenu().Actions().Add(action); err != nil {
			return err
		}
	}
	return nil
}

func (mw *MainWindowModel) runManagerWindow() {
	if mw.MainWindow != nil {
		win.SetForegroundWindow(mw.MainWindow.Handle())
		return
	}
	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "WireGuard for Windows",
		MinSize:  Size{650, 350},
		Icon:     wgicon,
		Layout:   HBox{},
		Children: []Widget{
			Composite{
				Layout:   VBox{MarginsZero: true, SpacingZero: true},
				Children: mw.listView(),
			},
			Composite{
				StretchFactor: 2,
				Layout:        VBox{MarginsZero: true},
				Children:      mw.detailView(),
			},
		},
	}.Run()
	mw.MainWindow = nil
}

func runAboutWindow(owner *walk.MainWindow) {

	Dialog{
		Size:      Size{300, 300},
		FixedSize: true,
		Icon:      wgicon,
		Layout:    VBox{},
		Children: []Widget{
			ImageView{
				Image:  "icon/128.png",
				Margin: 10,
				Mode:   ImageViewModeIdeal,
			},
			TextLabel{
				Text: "WireGuard",
				Font: Font{
					Family:    "MS Shell Dlg 2",
					PointSize: 14,
					Bold:      true,
				},
				TextAlignment: AlignHCenterVNear,
			},
			TextLabel{
				Text:          "App version: 12345\nGo backend version: 012345\n\nCopyright© 2018-2019 WireGuard LLC\nAll Rights Reserved",
				TextAlignment: AlignHCenterVNear,
			},
		},
	}.Run(nil)
}
