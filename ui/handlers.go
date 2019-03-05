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
	"golang.zx2c4.com/wireguard/windows/service"
)

const aboutText = `
WireGuard Windows.
TODO.

Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
`

func onAbout() {
	walk.MsgBox(mw, "About WireGuard", aboutText, walk.MsgBoxOK)
}

func onQuit() {
	tray.Dispose()
	_, err := service.IPCClientQuit(true)
	if err != nil {
		walk.MsgBox(nil, "Error Exiting WireGuard", fmt.Sprintf("Unable to exit service due to: %s. You may want to stop WireGuard from the service manager.", err), walk.MsgBoxIconError)
		os.Exit(1)
	}
}

func onManageTunnels() {
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

func onDelete() {
	// result is either walk.IDNO or walk.IDYES
	walk.MsgBox(mw, fmt.Sprintf(`Delete "%s"?`, "tunnel name"), fmt.Sprintf(`Are you sure you want to delete "%s"`, "tunnel name"), walk.MsgBoxYesNo|walk.MsgBoxIconWarning)
}

func onEdit() {
	// result is either walk.DlgCmdOK or walk.DlgCmdCancel
	getTunnelEdit().Run()
}
