/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpmAction0Size(t *testing.T) {

	const actualWtFwpmAction0Size = unsafe.Sizeof(wtFwpmAction0{})

	if actualWtFwpmAction0Size != wtFwpmAction0_Size {
		t.Errorf("Size of wtFwpmAction0 is %d, although %d is expected.", actualWtFwpmAction0Size,
			wtFwpmAction0_Size)
	}
}

func Test_wtFwpmAction0_Offsets(t *testing.T) {

	s := wtFwpmAction0{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.filterType)) - sp

	if offset != wtFwpmAction0_filterType_Offset {
		t.Errorf("wtFwpmAction0.filterType offset is %d although %d is expected", offset,
			wtFwpmAction0_filterType_Offset)
		return
	}
}
