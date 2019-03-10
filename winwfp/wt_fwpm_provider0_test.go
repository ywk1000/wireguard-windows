/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpProvider0Size(t *testing.T) {

	const actualWtFwpProvider0Size = unsafe.Sizeof(wtFwpProvider0{})

	if actualWtFwpProvider0Size != wtFwpProvider0_Size {
		t.Errorf("Size of wtFwpProvider0 is %d, although %d is expected.", actualWtFwpProvider0Size,
			wtFwpProvider0_Size)
	}
}

func TestWtFwpProvider0Offsets(t *testing.T) {

	s := wtFwpProvider0{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.displayData)) - sp

	if offset != wtFwpProvider0_displayData_Offset {
		t.Errorf("wtFwpProvider0.displayData offset is %d although %d is expected", offset,
			wtFwpProvider0_displayData_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.flags)) - sp

	if offset != wtFwpProvider0_flags_Offset {
		t.Errorf("wtFwpProvider0.flags offset is %d although %d is expected", offset,
			wtFwpProvider0_flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.providerData)) - sp

	if offset != wtFwpProvider0_providerData_Offset {
		t.Errorf("wtFwpProvider0.providerData offset is %d although %d is expected", offset,
			wtFwpProvider0_providerData_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.serviceName)) - sp

	if offset != wtFwpProvider0_serviceName_Offset {
		t.Errorf("wtFwpProvider0.serviceName offset is %d although %d is expected", offset,
			wtFwpProvider0_serviceName_Offset)
		return
	}
}
