/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// FWP_V4_ADDR_AND_MASK defined in fwptypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwptypes/ns-fwptypes-fwp_v4_addr_and_mask_).
type wtFwpV4AddrAndMask struct {
	addr uint32
	mask uint32
}
