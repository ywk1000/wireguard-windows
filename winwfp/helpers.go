/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"bytes"
	"fmt"
	"golang.org/x/sys/windows"
	"strconv"
	"unsafe"
)

const maxCStringLength = 100000

func wcharToString(wchar *uint16) string {
	return windows.UTF16ToString((*(*[maxCStringLength]uint16)(unsafe.Pointer(wchar)))[:])
}

func charToString(char *uint8) string {
	slice := (*(*[maxCStringLength]uint8)(unsafe.Pointer(char)))[:]
	null := bytes.IndexByte(slice, 0)
	if null != -1 {
		slice = slice[:null]
	}
	return string(slice)
}

func uint8ToBool(val uint8) bool {
	return val != 0
}

func boolToUint8(val bool) uint8 {
	if val {
		return 1
	} else {
		return 0
	}
}

func getZeroGuid() *windows.GUID {
	return &windows.GUID{Data1: 0, Data2: 0, Data3: 0, Data4: [8]byte{0, 0, 0, 0, 0, 0, 0, 0}}
}

func guidToString(guid *windows.GUID) string {
	if guid == nil {
		return "<nil>"
	} else {
		return fmt.Sprintf("{%06X-%04X-%04X-%04X-%012X}", guid.Data1, guid.Data2, guid.Data3, guid.Data4[:2],
			guid.Data4[2:])
	}
}

// TODO: For temporary usage. Delete when not needed anymore.
func stringToGuid(guid string) (*windows.GUID, error) {

	switch len(guid) {
	case 38:
		// {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
		if guid[0] != '{' || guid[37] != '}' {
			return nil, fmt.Errorf("stringToGuid() - invalid format")
		}
		guid = guid[1:]
	case 36:
		// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
		break;
	default:
		return nil, fmt.Errorf("stringToGuid() - invalid format")
	}

	if guid[8] != '-' || guid[13] != '-' || guid[18] != '-' || guid[23] != '-' {
		return nil, fmt.Errorf("stringToGuid() - invalid format")
	}

	data1, err := strconv.ParseUint(guid[0:8], 16, 32)

	if err != nil {
		return nil, err
	}

	data2, err := strconv.ParseUint(guid[9:13], 16, 16)

	if err != nil {
		return nil, err
	}

	data3, err := strconv.ParseUint(guid[14:18], 16, 16)

	if err != nil {
		return nil, err
	}

	var data4 [8]byte

	for i := 0; i < 8; i++ {

		idx := 2*i + 19

		if i > 1 {
			idx++
		}

		bt, err := strconv.ParseUint(guid[idx:idx+2], 16, 8)

		if err != nil {
			return nil, err
		}

		data4[i] = byte(bt)
	}

	return &windows.GUID{
		Data1: uint32(data1),
		Data2: uint16(data2),
		Data3: uint16(data3),
		Data4: data4,
	}, nil
}
