/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtSidSize(t *testing.T) {

	const actualWtSidSize = unsafe.Sizeof(wtSid{})

	if actualWtSidSize != wtSid_Size {
		t.Errorf("Size of wtSid is %d, although %d is expected.", actualWtSidSize, wtSid_Size)
	}
}

func TestWtSidOffsets(t *testing.T) {

	s := wtSid{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.SubAuthorityCount)) - sp

	if offset != wtSid_SubAuthorityCount_Offset {
		t.Errorf("wtSid.SubAuthorityCount offset is %d although %d is expected", offset,
			wtSid_SubAuthorityCount_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.IdentifierAuthority)) - sp

	if offset != wtSid_IdentifierAuthority_Offset {
		t.Errorf("wtSid.IdentifierAuthority offset is %d although %d is expected", offset,
			wtSid_IdentifierAuthority_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.SubAuthority)) - sp

	if offset != wtSid_SubAuthority_Offset {
		t.Errorf("wtSid.SubAuthority offset is %d although %d is expected", offset,
			wtSid_SubAuthority_Offset)
		return
	}
}
