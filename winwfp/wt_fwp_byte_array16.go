/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// FWP_BYTE_ARRAY16 defined in fwptypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwptypes/ns-fwptypes-fwp_byte_array16_)
type wtFwpByteArray16 struct {
	byteArray16 [16]uint8 // Windows type [16]UINT8
}
