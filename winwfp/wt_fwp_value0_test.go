/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpValue0Size(t *testing.T) {

	const actualWtFwpValue0Size = unsafe.Sizeof(wtFwpValue0{})

	if actualWtFwpValue0Size != wtFwpValue0_Size {
		t.Errorf("Size of wtFwpValue0 is %d, although %d is expected.", actualWtFwpValue0Size, wtFwpValue0_Size)
	}
}

func Test_wtFwpValue0_Offsets(t *testing.T) {

	s := wtFwpValue0{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.value)) - sp

	if offset != wtFwpValue0_value_Offset {
		t.Errorf("wtFwpValue0.value offset is %d although %d is expected", offset, wtFwpValue0_value_Offset)
		return
	}
}
