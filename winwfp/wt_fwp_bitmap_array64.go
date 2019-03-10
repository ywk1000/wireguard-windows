/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// FWP_BITMAP_ARRAY64 defined in fwtypes.h
type wtFwpBitmapArray64 struct {
	bitmapArray64 [8]uint8 // Windows type: [8]UINT8
}
