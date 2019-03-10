/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtSidAndAttributesSize(t *testing.T) {

	const actualWtSidAndAttributesSize = unsafe.Sizeof(wtSidAndAttributes{})

	if actualWtSidAndAttributesSize != wtSidAndAttributes_Size {
		t.Errorf("Size of wtSidAndAttributes is %d, although %d is expected.", actualWtSidAndAttributesSize,
			wtSidAndAttributes_Size)
	}
}

func TestWtSidAndAttributesOffsets(t *testing.T) {

	s := wtSidAndAttributes{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.Attributes)) - sp

	if offset != wtSidAndAttributes_Attributes_Offset {
		t.Errorf("wtSidAndAttributes.Attributes offset is %d although %d is expected", offset,
			wtSidAndAttributes_Attributes_Offset)
		return
	}
}
