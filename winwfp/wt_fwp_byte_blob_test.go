/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpByteBlobSize(t *testing.T) {

	const actualWtFwpByteBlobSize = unsafe.Sizeof(wtFwpByteBlob{})

	if actualWtFwpByteBlobSize != wtFwpByteBlob_Size {
		t.Errorf("Size of wtFwpByteBlob is %d, although %d is expected.", actualWtFwpByteBlobSize,
			wtFwpByteBlob_Size)
	}
}

func TestWtFwpByteBlobOffsets(t *testing.T) {

	s := wtFwpByteBlob{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.data)) - sp

	if offset != wtFwpByteBlob_data_Offset {
		t.Errorf("wtFwpByteBlob.data offset is %d although %d is expected", offset, wtFwpByteBlob_data_Offset)
		return
	}
}
