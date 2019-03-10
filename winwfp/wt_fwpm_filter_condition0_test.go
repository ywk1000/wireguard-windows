/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpmFilterCondition0Size(t *testing.T) {

	const actualWtFwpmFilterCondition0Size = unsafe.Sizeof(wtFwpmFilterCondition0{})

	if actualWtFwpmFilterCondition0Size != wtFwpmFilterCondition0_Size {
		t.Errorf("Size of wtFwpmFilterCondition0 is %d, although %d is expected.",
			actualWtFwpmFilterCondition0Size, wtFwpmFilterCondition0_Size)
	}
}

func TestWtFwpmFilterCondition0Offsets(t *testing.T) {

	s := wtFwpmFilterCondition0{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.matchType)) - sp

	if offset != wtFwpmFilterCondition0_matchType_Offset {
		t.Errorf("wtFwpmFilterCondition0.matchType offset is %d although %d is expected", offset,
			wtFwpmFilterCondition0_matchType_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.conditionValue)) - sp

	if offset != wtFwpmFilterCondition0_conditionValue_Offset {
		t.Errorf("wtFwpmFilterCondition0.conditionValue offset is %d although %d is expected", offset,
			wtFwpmFilterCondition0_conditionValue_Offset)
		return
	}
}
