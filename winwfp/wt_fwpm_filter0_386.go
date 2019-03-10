/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "golang.org/x/sys/windows"

// FWPM_FILTER0 defined in fwpmtypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwpmtypes/ns-fwpmtypes-fwpm_filter0_).
type wtFwpmFilter0 struct {
	filterKey           windows.GUID // Windows type: GUID
	displayData         wtFwpmDisplayData0
	flags               uint32
	providerKey         *windows.GUID // Windows type: *GUID
	providerData        wtFwpByteBlob
	layerKey            windows.GUID // Windows type: GUID
	subLayerKey         windows.GUID // Windows type: GUID
	weight              wtFwpValue0
	numFilterConditions uint32
	filterCondition     *wtFwpmFilterCondition0
	action              wtFwpmAction0
	offset1             [4]byte // Layout correction field
	providerContextKey  windows.GUID  // Windows type: GUID
	reserved            *windows.GUID // Windows type: *GUID
	offset2             [4]byte // Layout correction field
	filterId            uint64
	effectiveWeight     wtFwpValue0
}
