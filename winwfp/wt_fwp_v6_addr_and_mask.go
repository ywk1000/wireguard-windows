/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// FWP_V6_ADDR_AND_MASK defined in fwptypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwptypes/ns-fwptypes-fwp_v6_addr_and_mask_).
type wtFwpV6AddrAndMask struct {
	addr         [16]uint8
	prefixLength uint8
}
