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

	const actualWtFwpByteBlobSize = unsafe.Sizeof(FwpByteBlob{})

	if actualWtFwpByteBlobSize != fwpByteBlob_Size {
		t.Errorf("Size of FwpByteBlob is %d, although %d is expected.", actualWtFwpByteBlobSize,
			fwpByteBlob_Size)
	}
}

func TestWtFwpByteBlobOffsets(t *testing.T) {

	s := FwpByteBlob{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.data)) - sp

	if offset != fwpByteBlob_data_Offset {
		t.Errorf("FwpByteBlob.data offset is %d although %d is expected", offset, fwpByteBlob_data_Offset)
		return
	}
}
