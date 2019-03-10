/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"unsafe"
)

// SID defined in winnt.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/winnt/ns-winnt-_sid).
type wtSid struct {
	Revision            uint8 // Windows type: BYTE
	SubAuthorityCount   uint8 // Windows type: BYTE
	IdentifierAuthority SidIdentifierAuthority
	SubAuthority        [anysizeArray]uint32 // Windows type: [ANYSIZE_ARRAY]DWORD
}

func (sid *wtSid) toSid() *Sid {

	if sid == nil {
		return nil
	}

	subAuthority := make([]uint32, sid.SubAuthorityCount, sid.SubAuthorityCount)

	pFirst := uintptr(unsafe.Pointer(&sid.SubAuthority[0]))
	uint32Size := uintptr(4) // Should be equal to unsafe.Sizeof(sid.SubAuthority[0])

	for i := uint8(0); i < sid.SubAuthorityCount; i++ {
		subAuthority[i] = *(*uint32)(unsafe.Pointer(pFirst + uint32Size*uintptr(i)))
	}

	return &Sid{
		Revision:            sid.Revision,
		IdentifierAuthority: sid.IdentifierAuthority,
		SubAuthority:        subAuthority,
	}
}

func (sid *wtSid) String() string {

	if sid == nil {
		return "<nil>"
	}

	str := make([]uint16, maxCStringLength)

	strPtr := (*uint16)(unsafe.Pointer(&str[0]))

	success := convertSidToStringSidW(sid, strPtr)

	if uint8ToBool(success) {
		return windows.UTF16ToString(str)
	} else {
		return fmt.Sprint(os.NewSyscallError("fwpuclnt.convertSidToStringSidW", windows.GetLastError()))
	}
}
