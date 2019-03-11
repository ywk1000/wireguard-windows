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

// Wraps GetModuleFileNameW function
// (https://docs.microsoft.com/en-us/windows/desktop/api/libloaderapi/nf-libloaderapi-getmodulefilenamew).
// Input parameter is module handle, which can be 0, in which case the method will return the path of the executable
// file of the current process.
func GetModuleFileNameWWrapper(module uintptr) (string, error) {

	moduleFileNameLength := uint32(10000)

	for {
		buffer := make([]uint16, moduleFileNameLength, moduleFileNameLength)

		result := getModuleFileNameW(module, (*uint16)(unsafe.Pointer(&buffer[0])), moduleFileNameLength)

		if result == 0 {
			return "", os.NewSyscallError("Kernel32.GetModuleFileNameW", windows.GetLastError())
		}

		if result < moduleFileNameLength {

			buffer[result] = 0

			return windows.UTF16ToString(buffer), nil
		}

		if windows.GetLastError() == windows.ERROR_INSUFFICIENT_BUFFER {
			moduleFileNameLength += 10000
		} else {
			return windows.UTF16ToString(buffer), nil
		}
	}
}

// Returns FwpAppId of module identified by the input paramter. The input paramter can be 0, in which case the method
// will return FwpAppId of the current app (process).
// NOTE: Returned FwpAppId (FwpByteBlob struct) should be destroyed when not needed anymore by calling its Free() method.
func GetFwpAppId(module uintptr) (*FwpByteBlob, error) {

	currentFile, err := GetModuleFileNameWWrapper(module)

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

// Returns the same as GetFwpAppId when called with input argument module = 0.
func GetCurrentAppId() (*FwpByteBlob, error) {
	return GetFwpAppId(0)
}
