/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package ui

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/lxn/walk"
	"github.com/lxn/win"
	"golang.org/x/sys/windows"
	"golang.zx2c4.com/wireguard/windows/conf"
	"golang.zx2c4.com/wireguard/windows/service"
)

type labelTextLine struct {
	label *walk.TextLabel
	text  *walk.LineEdit
}

type interfaceView struct {
	status     *labelTextLine
	publicKey  *labelTextLine
	listenPort *labelTextLine
	mtu        *labelTextLine
	addresses  *labelTextLine
	dns        *labelTextLine
	toggler    *walk.PushButton
}

type peerView struct {
	publicKey           *labelTextLine
	presharedKey        *labelTextLine
	allowedIPs          *labelTextLine
	endpoint            *labelTextLine
	persistentKeepalive *labelTextLine
	latestHandshake     *labelTextLine
	transfer            *labelTextLine
}

type ConfView struct {
	*walk.ScrollView
	name      *walk.GroupBox
	interfaze *interfaceView
	peers     map[conf.Key]*peerView
	spacer    *walk.Spacer

	originalWndProc uintptr
	creatingThread  uint32
}

func (lt *labelTextLine) show(text string) {
	s, e := lt.text.TextSelection()
	lt.text.SetText(text)
	lt.label.SetVisible(true)
	lt.text.SetVisible(true)
	lt.text.SetTextSelection(s, e)
}

func (lt *labelTextLine) hide() {
	lt.text.SetText("")
	lt.label.SetVisible(false)
	lt.text.SetVisible(false)
}

func newLabelTextLine(fieldName string, parent walk.Container) *labelTextLine {
	lt := new(labelTextLine)
	lt.label, _ = walk.NewTextLabel(parent)
	lt.label.SetText(fieldName + ":")
	lt.label.SetTextAlignment(walk.AlignHFarVNear)
	lt.label.SetVisible(false)

	lt.text, _ = walk.NewLineEdit(parent)
	win.SetWindowLong(lt.text.Handle(), win.GWL_EXSTYLE, win.GetWindowLong(lt.text.Handle(), win.GWL_EXSTYLE)&^win.WS_EX_CLIENTEDGE)
	lt.text.SetReadOnly(true)
	lt.text.SetBackground(walk.NullBrush())
	lt.text.SetVisible(false)
	lt.text.FocusedChanged().Attach(func() {
		lt.text.SetTextSelection(0, 0)
	})
	return lt
}

func newInterfaceView(parent walk.Container) *interfaceView {
	iv := &interfaceView{
		status:     newLabelTextLine("Status", parent),
		publicKey:  newLabelTextLine("Public key", parent),
		listenPort: newLabelTextLine("Listen port", parent),
		mtu:        newLabelTextLine("MTU", parent),
		addresses:  newLabelTextLine("Addresses", parent),
		dns:        newLabelTextLine("DNS servers", parent),
	}
	buttonContainer, _ := walk.NewComposite(parent)
	parent.Layout().(*walk.GridLayout).SetRange(buttonContainer, walk.Rectangle{1, 6, 1, 1})
	buttonContainer.SetLayout(walk.NewHBoxLayout())
	buttonContainer.Layout().SetMargins(walk.Margins{})
	iv.toggler, _ = walk.NewPushButton(buttonContainer)
	walk.NewHSpacer(buttonContainer)

	layoutInGrid(iv, parent.Layout().(*walk.GridLayout), true)

	return iv
}

func newPeerView(parent walk.Container) *peerView {
	pv := &peerView{
		newLabelTextLine("Public key", parent),
		newLabelTextLine("Preshared key", parent),
		newLabelTextLine("Allowed IPs", parent),
		newLabelTextLine("Endpoint", parent),
		newLabelTextLine("Persistent keepalive", parent),
		newLabelTextLine("Latest handshake", parent),
		newLabelTextLine("Transfer", parent),
	}
	layoutInGrid(pv, parent.Layout().(*walk.GridLayout), false)
	return pv
}

func layoutInGrid(view interface{}, layout *walk.GridLayout, layoutLast bool) {
	v := reflect.ValueOf(view).Elem()
	j := v.NumField()
	if !layoutLast {
		j -= 1
	}
	for i := 0; i < j; i++ {
		lt := (*labelTextLine)(unsafe.Pointer(v.Field(i).Pointer()))
		layout.SetRange(lt.label, walk.Rectangle{0, i, 1, 1})
		layout.SetRange(lt.text, walk.Rectangle{1, i, 1, 1})
	}
}

func (iv *interfaceView) apply(c *conf.Interface, status service.TunnelState) {
	switch status {
	case service.TunnelUnknown:
		iv.status.show("Unknown")
		iv.toggler.SetText("Activate")
		iv.toggler.SetEnabled(false)
	case service.TunnelStarted:
		iv.status.show("Active")
		iv.toggler.SetText("Deactivate")
		iv.toggler.SetEnabled(true)
	case service.TunnelStopped:
		iv.status.show("Inactive")
		iv.toggler.SetText("Activate")
		iv.toggler.SetEnabled(true)
	case service.TunnelStarting:
		iv.status.show("Activating")
		iv.toggler.SetText("Activating...")
		iv.toggler.SetEnabled(false)
	case service.TunnelStopping:
		iv.status.show("Deactivating")
		iv.toggler.SetText("Deactivating...")
		iv.toggler.SetEnabled(false)
	}

	iv.publicKey.show(c.PrivateKey.Public().String())

	if c.ListenPort > 0 {
		iv.listenPort.show(strconv.Itoa(int(c.ListenPort)))
	} else {
		iv.listenPort.hide()
	}

	if c.Mtu > 0 {
		iv.mtu.show(strconv.Itoa(int(c.Mtu)))
	} else {
		iv.mtu.hide()
	}

	if len(c.Addresses) > 0 {
		addrStrings := make([]string, len(c.Addresses))
		for i, address := range c.Addresses {
			addrStrings[i] = address.String()
		}
		iv.addresses.show(strings.Join(addrStrings[:], ", "))
	} else {
		iv.addresses.hide()
	}

	if len(c.Dns) > 0 {
		addrStrings := make([]string, len(c.Dns))
		for i, address := range c.Dns {
			addrStrings[i] = address.String()
		}
		iv.dns.show(strings.Join(addrStrings[:], ", "))
	} else {
		iv.dns.hide()
	}
}

func (pv *peerView) apply(c *conf.Peer) {
	pv.publicKey.show(c.PublicKey.String())

	if !c.PresharedKey.IsZero() {
		pv.presharedKey.show("enabled")
	} else {
		pv.presharedKey.hide()
	}

	if len(c.AllowedIPs) > 0 {
		addrStrings := make([]string, len(c.AllowedIPs))
		for i, address := range c.AllowedIPs {
			addrStrings[i] = address.String()
		}
		pv.allowedIPs.show(strings.Join(addrStrings[:], ", "))
	} else {
		pv.allowedIPs.hide()
	}

	if !c.Endpoint.IsEmpty() {
		pv.endpoint.show(c.Endpoint.String())
	} else {
		pv.endpoint.hide()
	}

	if c.PersistentKeepalive > 0 {
		pv.persistentKeepalive.show(strconv.Itoa(int(c.PersistentKeepalive)))
	} else {
		pv.persistentKeepalive.hide()
	}

	if !c.LastHandshakeTime.IsEmpty() {
		pv.latestHandshake.show(c.LastHandshakeTime.String())
	} else {
		pv.latestHandshake.hide()
	}

	if c.RxBytes > 0 || c.TxBytes > 0 {
		pv.transfer.show(fmt.Sprintf("%s received, %s sent", c.RxBytes.String(), c.TxBytes.String()))
	} else {
		pv.transfer.hide()
	}
}

func newPaddedGroupGrid(parent walk.Container) (group *walk.GroupBox, err error) {
	group, err = walk.NewGroupBox(parent)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			group.Dispose()
		}
	}()
	layout := walk.NewGridLayout()
	layout.SetMargins(walk.Margins{10, 15, 10, 5})
	err = group.SetLayout(layout)
	if err != nil {
		return nil, err
	}
	spacer, err := walk.NewSpacerWithCfg(group, &walk.SpacerCfg{walk.GrowableHorz | walk.GreedyHorz, walk.Size{}, false})
	if err != nil {
		return nil, err
	}
	layout.SetRange(spacer, walk.Rectangle{1, 0, 1, 1})
	return group, nil
}

func NewConfView(parent walk.Container) (*ConfView, error) {
	cv := new(ConfView)
	cv.ScrollView, _ = walk.NewScrollView(parent)
	cv.SetLayout(walk.NewVBoxLayout())
	cv.name, _ = newPaddedGroupGrid(cv)
	cv.interfaze = newInterfaceView(cv.name)
	cv.peers = make(map[conf.Key]*peerView)
	cv.spacer, _ = walk.NewVSpacer(cv)
	cv.creatingThread = windows.GetCurrentThreadId()
	win.SetWindowLongPtr(cv.Handle(), win.GWLP_USERDATA, uintptr(unsafe.Pointer(cv)))
	cv.originalWndProc = win.SetWindowLongPtr(cv.Handle(), win.GWL_WNDPROC, crossThreadMessageHijack)
	return cv, nil
}

//TODO: choose actual good value for this
const crossThreadUpdate = win.WM_APP + 17

var crossThreadMessageHijack = windows.NewCallback(func(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	cv := (*ConfView)(unsafe.Pointer(win.GetWindowLongPtr(hwnd, win.GWLP_USERDATA)))
	if msg == crossThreadUpdate {
		cv.setConfiguration((*conf.Config)(unsafe.Pointer(wParam)), *(*service.TunnelState)(unsafe.Pointer(lParam)))
		return 0
	}
	return win.CallWindowProc(cv.originalWndProc, hwnd, msg, wParam, lParam)
})

func (cv *ConfView) SetConfiguration(c *conf.Config, status service.TunnelState) {
	if cv.creatingThread == windows.GetCurrentThreadId() {
		cv.setConfiguration(c, status)
	} else {
		cv.SendMessage(crossThreadUpdate, uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(&status)))
	}
}

func (cv *ConfView) setConfiguration(c *conf.Config, status service.TunnelState) {
	hasSuspended := false
	suspend := func() {
		if !hasSuspended {
			cv.SetSuspended(true)
			hasSuspended = true
		}
	}
	defer func() {
		if hasSuspended {
			cv.SetSuspended(false)
			cv.SendMessage(win.WM_SIZING, 0, 0) //TODO: FILTHY HACK! And doesn't work when items disappear.
		}
	}()
	title := "Interface: " + c.Name
	if cv.name.Title() != title {
		cv.name.SetTitle(title)
	}
	cv.interfaze.apply(&c.Interface, status)
	inverse := make(map[*peerView]bool, len(cv.peers))
	for _, pv := range cv.peers {
		inverse[pv] = true
	}
	didAddPeer := false
	for _, peer := range c.Peers {
		if pv := cv.peers[peer.PublicKey]; pv != nil {
			pv.apply(&peer)
			inverse[pv] = false
		} else {
			didAddPeer = true
			suspend()
			group, _ := newPaddedGroupGrid(cv)
			group.SetTitle("Peer")
			pv := newPeerView(group)
			pv.apply(&peer)
			cv.peers[peer.PublicKey] = pv
		}
	}
	for pv, remove := range inverse {
		if !remove {
			continue
		}
		k, e := conf.NewPrivateKeyFromString(pv.publicKey.text.Text())
		if e != nil {
			continue
		}
		suspend()
		delete(cv.peers, *k)
		groupBox := pv.publicKey.label.Parent().AsContainerBase().Parent().(*walk.GroupBox)
		groupBox.Parent().Children().Remove(groupBox)
		groupBox.Dispose()
	}
	if didAddPeer {
		cv.Children().Remove(cv.spacer)
		cv.Children().Add(cv.spacer)
	}
}
