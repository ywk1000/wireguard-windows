/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpmDisplayData0Size(t *testing.T) {

	const actualWtFwpmDisplayData0Size = unsafe.Sizeof(wtFwpmDisplayData0{})

	if actualWtFwpmDisplayData0Size != wtFwpmDisplayData0_Size {
		t.Errorf("Size of wtFwpmDisplayData0 is %d, although %d is expected.", actualWtFwpmDisplayData0Size,
			wtFwpmDisplayData0_Size)
	}
}

func TestWtFwpmDisplayData0Offsets(t *testing.T) {

	s := wtFwpmDisplayData0{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.description)) - sp

	if offset != wtFwpmDisplayData0_description_Offset {
		t.Errorf("wtFwpmDisplayData0.description offset is %d although %d is expected", offset,
			wtFwpmDisplayData0_description_Offset)
		return
	}
}
