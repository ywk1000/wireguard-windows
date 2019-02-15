package main

import (
	"encoding/base64"

	"git.zx2c4.com/wireguard-windows/manager/conf"
	"git.zx2c4.com/wireguard-windows/manager/walk"
	. "git.zx2c4.com/wireguard-windows/manager/walk/declarative"
	"golang.org/x/crypto/curve25519"
)

func (mw *MainWindowModel) detailView() []Widget {
	return []Widget{
		ConfViewDecl{
			AssignTo: &mw.pv,
		},
		Composite{
			Layout: HBox{},
			Children: []Widget{

				HSpacer{},
				PushButton{
					Text: "Edit",
					OnClicked: func() {
						if mw.model.Current == nil {
							return
						}
						if ok, err := mw.runEditDialog(); err == nil && ok == walk.DlgCmdOK {
							mw.pv.SetConfiguration(mw.model.Current)
							mw.model.items[mw.lb.CurrentIndex()] = mw.model.Current
							mw.model.PublishItemChanged(mw.lb.CurrentIndex())
						}
					},
				},
			},
		},
	}
}

func (mw *MainWindowModel) runEditDialog() (int, error) {
	var dlg *walk.Dialog
	var se *SyntaxEdit
	var tl *walk.TextLabel
	var le *walk.LineEdit

	lastPrivate := ""
	return Dialog{
		AssignTo: &dlg,
		Icon:     wgicon,
		Title:    "Edit Tunnel",
		Layout:   Grid{Columns: 2},
		MinSize:  Size{500, 400},
		Children: []Widget{
			Label{
				Text: "Name:",
			},
			LineEdit{
				AssignTo: &le,
				Text:     mw.model.Current.Name,
			},
			TextLabel{
				ColumnSpan: 2,
				AssignTo:   &tl,
				Text:       "Public key: (unknown)",
			},
			Composite{
				ColumnSpan: 2,
				Layout:     HBox{MarginsZero: true},
				Children: []Widget{
					SyntaxEditDecl{
						AssignTo: &se,
						Text:     mw.model.Current.ToWgQuick(),
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
			},

			Composite{
				ColumnSpan: 2,
				Layout:     HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: "Save",
						OnClicked: func() {
							tun, err := conf.FromWgQuick(se.Text(), le.Text())
							if err != nil {
								walk.MsgBox(dlg, "Invalid Configuration", err.Error(), walk.MsgBoxIconError)
								return
							}
							mw.model.Current = tun
							dlg.Accept()
						},
					},
					PushButton{
						Text: "Cancel",
						OnClicked: func() {
							dlg.Cancel()
						},
					},
				},
			},
		},
	}.Run(mw.MainWindow)
}
