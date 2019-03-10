/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "unsafe"

// FWP_BYTE_BLOB defined in fwptypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwptypes/ns-fwptypes-fwp_byte_blob_)
type wtFwpByteBlob struct {
	size uint32
	data *uint8
}

func (bb *wtFwpByteBlob) toSlice() []uint8 {

	if bb == nil {
		return nil
	}

	bbSlice := make([]uint8, bb.size, bb.size)

	if bb.size < 1 {
		return bbSlice
	}

	pFirst := uintptr(unsafe.Pointer(bb.data))
	uint8Size := uintptr(1) // Size of uint8

	for i := uint32(0); i < bb.size; i++ {
		bbSlice[i] = *(*uint8)(unsafe.Pointer(pFirst + uint8Size*uintptr(i)))
	}

	return bbSlice
}
