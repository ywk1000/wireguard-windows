/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtSecWinntAuthIdentityWSize(t *testing.T) {

	const actualWtSecWinntAuthIdentityWSize = unsafe.Sizeof(wtSecWinntAuthIdentityW{})

	if actualWtSecWinntAuthIdentityWSize != wtSecWinntAuthIdentityW_Size {
		t.Errorf("Size of wtSecWinntAuthIdentityW is %d, although %d is expected.",
			actualWtSecWinntAuthIdentityWSize, wtSecWinntAuthIdentityW_Size)
	}
}

func TestWtSecWinntAuthIdentityWOffsets(t *testing.T) {

	s := wtSecWinntAuthIdentityW{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.UserLength)) - sp

	if offset != wtSecWinntAuthIdentityW_UserLength_Offset {
		t.Errorf("wtSecWinntAuthIdentityW.UserLength offset is %d although %d is expected", offset,
			wtSecWinntAuthIdentityW_UserLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Domain)) - sp

	if offset != wtSecWinntAuthIdentityW_Domain_Offset {
		t.Errorf("wtSecWinntAuthIdentityW.Domain offset is %d although %d is expected", offset,
			wtSecWinntAuthIdentityW_Domain_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.DomainLength)) - sp

	if offset != wtSecWinntAuthIdentityW_DomainLength_Offset {
		t.Errorf("wtSecWinntAuthIdentityW.DomainLength offset is %d although %d is expected", offset,
			wtSecWinntAuthIdentityW_DomainLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Password)) - sp

	if offset != wtSecWinntAuthIdentityW_Password_Offset {
		t.Errorf("wtSecWinntAuthIdentityW.Password offset is %d although %d is expected", offset,
			wtSecWinntAuthIdentityW_Password_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.PasswordLength)) - sp

	if offset != wtSecWinntAuthIdentityW_PasswordLength_Offset {
		t.Errorf("wtSecWinntAuthIdentityW.PasswordLength offset is %d although %d is expected", offset,
			wtSecWinntAuthIdentityW_PasswordLength_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.Flags)) - sp

	if offset != wtSecWinntAuthIdentityW_Flags_Offset {
		t.Errorf("wtSecWinntAuthIdentityW.Flags offset is %d although %d is expected", offset,
			wtSecWinntAuthIdentityW_Flags_Offset)
		return
	}
}
