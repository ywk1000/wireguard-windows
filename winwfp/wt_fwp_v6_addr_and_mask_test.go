/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpV6AddrAndMaskSize(t *testing.T) {

	const actualWtFwpV6AddrAndMaskSize = unsafe.Sizeof(wtFwpV6AddrAndMask{})

	if actualWtFwpV6AddrAndMaskSize != wtFwpV6AddrAndMask_Size {
		t.Errorf("Size of wtFwpV6AddrAndMask is %d, although %d is expected.", actualWtFwpV6AddrAndMaskSize,
			wtFwpV6AddrAndMask_Size)
	}
}

func TestWtFwpV6AddrAndMaskOffsets(t *testing.T) {

	s := wtFwpV6AddrAndMask{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.prefixLength)) - sp

	if offset != wtFwpV6AddrAndMask_prefixLength_Offset {
		t.Errorf("wtFwpV6AddrAndMask.prefixLength offset is %d although %d is expected", offset,
			wtFwpV6AddrAndMask_prefixLength_Offset)
		return
	}
}
