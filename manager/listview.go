package main

import (
	"time"

	"git.zx2c4.com/wireguard-windows/manager/conf"
	"git.zx2c4.com/wireguard-windows/manager/walk"
	. "git.zx2c4.com/wireguard-windows/manager/walk/declarative"
)

type InterfacesModel struct {
	walk.ListModelBase
	items   []*conf.Config
	Current *conf.Config
}

func (m *InterfacesModel) ItemCount() int {
	return len(m.items)
}

func (m *InterfacesModel) Value(index int) interface{} {
	return m.items[index].Name
}

func (mw *MainWindowModel) listView() []Widget {
	return []Widget{
		ListBox{
			AssignTo:              &mw.lb,
			Model:                 mw.model,
			OnCurrentIndexChanged: mw.currentIndexChanged,
		},
		Composite{
			Layout: HBox{MarginsZero: true, SpacingZero: true},
			Children: []Widget{
				ToolBar{
					ButtonStyle: ToolBarButtonTextOnly,
					Items: []MenuItem{
						Menu{
							Text: "+",
							Items: []MenuItem{
								Action{
									Text:        "Add empty tunnel",
									OnTriggered: mw.newTunnel,
								},
								Action{
									Text: "Import tunnel(s) from file...",
									OnTriggered: func() {
										if err := mw.importTunnels(); err != nil {
											walk.MsgBox(mw, err.Error(), "Error", walk.MsgBoxIconError)
										}
									},
								},
							},
						},
						Action{
							Text:        "−", //special character. normal "-" doesn't display??
							OnTriggered: mw.removeTunnel,
						},
						Menu{
							Text: "⚙",
							Items: []MenuItem{
								Action{
									Text:        "Export log to file...",
									OnTriggered: mw.exportLog,
								},
								Action{
									Text: "Export tunnels to zip...",
									OnTriggered: func() {
										if err := mw.exportTunnels(); err != nil {
											walk.MsgBox(mw, err.Error(), "Error", walk.MsgBoxIconError)
										}
									},
								},
								Action{
									Text: "Export selected tunnel...",
									OnTriggered: func() {
										if err := mw.exportCurrentTunnel(); err != nil {
											walk.MsgBox(mw, err.Error(), "Error", walk.MsgBoxIconError)
										}
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (mw *MainWindowModel) currentIndexChanged() {

	if mw.refreshtimer == nil {
		mw.refreshtimer = time.AfterFunc(time.Second, func() {
			mw.pv.SetConfiguration(mw.model.Current)
			mw.refreshtimer.Reset(time.Second)
		})
	}

	i := mw.lb.CurrentIndex()
	var current *conf.Config
	if i >= 0 && i < len(mw.model.items) {
		current = mw.model.items[i]
	}
	mw.model.Current = current

	mw.pv.SetConfiguration(current)

	if current != nil && len(current.Peers) > 0 {
		current.Peers[0].LastHandshakeTime = conf.HandshakeTime(time.Duration(time.Now().Unix()) * time.Second)
	}
}

func (mw *MainWindowModel) removeTunnel() {
	if i := mw.lb.CurrentIndex(); i >= 0 &&
		walk.MsgBox(mw,
			"Confirm deletion",
			"Are you sure you want to remove this interface?",
			walk.MsgBoxYesNo|walk.MsgBoxIconQuestion) == walk.DlgCmdYes {
		mw.model.items = append(mw.model.items[:i], mw.model.items[i+1:]...)
		mw.model.PublishItemsReset()
	}
}

func (mw *MainWindowModel) newTunnel() {
	mw.model.items = append(mw.model.items, &conf.Config{Name: "unnamed"})
	mw.model.PublishItemsReset()
}
