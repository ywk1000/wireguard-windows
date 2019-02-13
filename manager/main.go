/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package main

import (
	"time"

	"git.zx2c4.com/wireguard-windows/manager/conf"
	"git.zx2c4.com/wireguard-windows/manager/walk"
	. "git.zx2c4.com/wireguard-windows/manager/walk/declarative"
)

type MainWindowModel struct {
	*walk.MainWindow
	model        *InterfacesModel
	lb           *walk.ListBox
	pv           *ConfView
	refreshtimer *time.Timer
}

func main() {

	mw := &MainWindowModel{model: &InterfacesModel{}}
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
	mw.model.items = append(mw.model.items, tunnel)

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "WireGuard for Windows",
		MinSize:  Size{650, 350},
		Icon:     "icon/icon.ico",
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
}
