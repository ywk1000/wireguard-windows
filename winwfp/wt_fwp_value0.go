/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// FWP_VALUE0 defined in fwptypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwptypes/ns-fwptypes-fwp_value0_)
type wtFwpValue0 struct {
	_type wtFwpDataType
	value uintptr
}
