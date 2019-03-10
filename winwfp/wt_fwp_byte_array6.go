/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// FWP_BYTE_ARRAY6 defined in fwtypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwptypes/ns-fwptypes-fwp_byte_array6_)
type wtFwpByteArray6 struct {
	byteArray6 [6]uint8 // Windows type: [6]UINT8
}
