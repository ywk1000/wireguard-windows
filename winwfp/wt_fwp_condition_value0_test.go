/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpConditionValue0Size(t *testing.T) {

	const actualWtFwpConditionValue0Size = unsafe.Sizeof(wtFwpConditionValue0{})

	if actualWtFwpConditionValue0Size != wtFwpConditionValue0_Size {
		t.Errorf("Size of wtFwpConditionValue0 is %d, although %d is expected.", actualWtFwpConditionValue0Size,
			wtFwpConditionValue0_Size)
	}
}

func TestWtFwpConditionValue0Offsets(t *testing.T) {

	s := wtFwpConditionValue0{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.value)) - sp

	if offset != wtFwpConditionValue0_uint8_Offset {
		t.Errorf("wtFwpConditionValue0.value offset is %d although %d is expected", offset, wtFwpConditionValue0_uint8_Offset)
		return
	}
}
