/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"golang.org/x/sys/windows"
	"os"
	"unsafe"
)

// Returns FwpAppId of current process.
func GetCurrentAppId() (*FwpByteBlob, error) {

	currentFile, err := os.Executable()

	if err != nil {
		return nil, err
	}

	curFilePtr, err := windows.UTF16PtrFromString(currentFile)

	if err != nil {
		return nil, err
	}

	var appId *FwpByteBlob = nil

	result := fwpmGetAppIdFromFileName0(curFilePtr, unsafe.Pointer(&appId))

	if result == 0 {
		return appId, nil
	} else {
		return nil, os.NewSyscallError("fwpuclnt.FwpmGetAppIdFromFileName0", windows.Errno(result))
	}
}
