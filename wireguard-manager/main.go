/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package main

import (
	"encoding/base64"
	"git.zx2c4.com/wireguard-windows/wireguard-manager/walk"
	. "git.zx2c4.com/wireguard-windows/wireguard-manager/walk/declarative"
	"golang.org/x/crypto/curve25519"
	"log"
)

func main() {
	var se *SyntaxEdit
	var tl *walk.TextLabel
	lastPrivate := ""
	icon, err := walk.Resources.Icon("./icon/icon.ico")
	if err != nil {
		log.Fatal(err)
	}

	demo_config := `[Interface]
PrivateKey = cJM9wXUVXc0fa/t5b/Lm0BHNx6jh5UiTLsO+oJhyQUU=
Address = 192.168.4.84/24, 2001:abcd:33::/120 # this is a comment
DNS = 8.8.8.8, 8.8.4.4, 1.1.1.1, 1.0.0.1
PreUp = calc.exe
ListenPort = notarealnumber

[Peer]
PublicKey = JRI8Xc0zKP9kXk8qP84NdUQA04h6DLfFbwJn4g+/PFs=
Endpoint = demo.wireguard.com:12912
AllowedIPs = 0.0.0.0/0

[Peer]
PublicKey = QCssGR6joqOIEQW6j2AR7nMcXJIVI9E9PCcbwrVXhU8=
Endpoint = intranet.wireguard.com:51820
AllowedIPs = 192.168.22.0/24, fd00:3001::/64
`

	MainWindow{
		Title:   "WireGuard for Windows",
		Icon:    icon,
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			TextLabel{
				AssignTo: &tl,
				Text:     "Public key: (unknown)",
			},
			SyntaxEditDecl{
				AssignTo: &se,
				Text:     demo_config,
				OnPrivateKeyChange: func(privateKey string) {
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
				},
			},
		},
	}.Run()
}
