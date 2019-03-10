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

// SID_IDENTIFIER_AUTHORITY defined in winnt.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/winnt/ns-winnt-_sid_identifier_authority).
type SidIdentifierAuthority struct {
	Value [6]uint8 // Windows type: [6]BYTE
}

// Corresponds to SID defined in winnt.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/winnt/ns-winnt-_sid).
type Sid struct {
	Revision            uint8
	IdentifierAuthority SidIdentifierAuthority
	SubAuthority        []uint32
}

func (sid *Sid) toWtSid() (*wtSid, error) {

	if sid == nil {
		return nil, nil
	}

	subAuthorityCount := uint8(len(sid.SubAuthority))

	if subAuthorityCount < 1 || subAuthorityCount > 8 {
		return nil, fmt.Errorf("Sid.SubAuthority must contain 1 to 8 items")
	}

	getSubAuthority := func(subAuthorities []uint32, idx int) uint32 {
		if len(subAuthorities) > idx {
			return subAuthorities[idx]
		} else {
			return uint32(0)
		}
	}

	var wt *wtSid = nil

	result := allocateAndInitializeSid(&sid.IdentifierAuthority, subAuthorityCount, sid.SubAuthority[0],
		getSubAuthority(sid.SubAuthority, 1), getSubAuthority(sid.SubAuthority, 2),
		getSubAuthority(sid.SubAuthority, 3), getSubAuthority(sid.SubAuthority, 4),
		getSubAuthority(sid.SubAuthority, 5), getSubAuthority(sid.SubAuthority, 6),
		getSubAuthority(sid.SubAuthority, 7), unsafe.Pointer(&wt))

	if uint8ToBool(result) {
		return wt, nil
	} else {
		return nil, os.NewSyscallError("fwpuclnt.AllocateAndInitializeSid", windows.GetLastError())
	}
}

func (sid *Sid) String() string {

	if sid == nil {
		return "<nil>"
	}

	wt, err := sid.toWtSid()

	if wt != nil {
		defer freeSid(wt)
	}

	if err != nil {
		return fmt.Sprintf("< Error: %v >", err)
	}

	return wt.String()
}
