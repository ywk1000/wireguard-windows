/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpmSublayer0Size(t *testing.T) {

	const actualWtFwpmSublayer0Size = unsafe.Sizeof(wtFwpmSublayer0{})

	if actualWtFwpmSublayer0Size != wtFwpmSublayer0_Size {
		t.Errorf("Size of wtFwpmSublayer0 is %d, although %d is expected.", actualWtFwpmSublayer0Size,
			wtFwpmSublayer0_Size)
	}
}

func Test_wtFwpmSublayer0_Offsets(t *testing.T) {

	s := wtFwpmSublayer0{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.displayData)) - sp

	if offset != wtFwpmSublayer0_displayData_Offset {
		t.Errorf("wtFwpmSublayer0.displayData offset is %d although %d is expected", offset,
			wtFwpmSublayer0_displayData_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.flags)) - sp

	if offset != wtFwpmSublayer0_flags_Offset {
		t.Errorf("wtFwpmSublayer0.flags offset is %d although %d is expected", offset,
			wtFwpmSublayer0_flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.providerKey)) - sp

	if offset != wtFwpmSublayer0_providerKey_Offset {
		t.Errorf("wtFwpmSublayer0.providerKey offset is %d although %d is expected", offset,
			wtFwpmSublayer0_providerKey_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.providerData)) - sp

	if offset != wtFwpmSublayer0_providerData_Offset {
		t.Errorf("wtFwpmSublayer0.providerData offset is %d although %d is expected", offset,
			wtFwpmSublayer0_providerData_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.weight)) - sp

	if offset != wtFwpmSublayer0_weight_Offset {
		t.Errorf("wtFwpmSublayer0.weight offset is %d although %d is expected", offset,
			wtFwpmSublayer0_weight_Offset)
		return
	}
}
