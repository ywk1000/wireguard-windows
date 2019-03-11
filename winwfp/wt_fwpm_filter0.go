/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// Cleans up some fields so that the same wtFwpmFilter0 struct can be reused.
func (f *wtFwpmFilter0) cleanup() {
	f.providerContextKey = *getZeroGuid()
	f.reserved = nil
	f.filterId = 0
	f.effectiveWeight = wtFwpValue0{
		_type: FWP_EMPTY,
		value: 0,
	}
}
