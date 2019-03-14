/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpV4AddrAndMaskSize(t *testing.T) {

	const actualWtFwpV4AddrAndMaskSize = unsafe.Sizeof(wtFwpV4AddrAndMask{})

	if actualWtFwpV4AddrAndMaskSize != wtFwpV4AddrAndMask_Size {
		t.Errorf("Size of wtFwpV4AddrAndMask is %d, although %d is expected.", actualWtFwpV4AddrAndMaskSize,
			wtFwpV4AddrAndMask_Size)
	}
}

func TestWtFwpV4AddrAndMaskOffsets(t *testing.T) {

	s := wtFwpV4AddrAndMask{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.mask)) - sp

	if offset != wtFwpV4AddrAndMask_mask_Offset {
		t.Errorf("wtFwpV4AddrAndMask.mask offset is %d although %d is expected", offset,
			wtFwpV4AddrAndMask_mask_Offset)
		return
	}
}
