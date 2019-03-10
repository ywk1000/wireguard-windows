/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "golang.org/x/sys/windows"

// FWPM_PROVIDER0 defined in fwpmtypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwpmtypes/ns-fwpmtypes-fwpm_provider0_)
type wtFwpProvider0 struct {
	providerKey  windows.GUID // Windows type: GUID
	displayData  wtFwpmDisplayData0
	flags        uint32
	providerData wtFwpByteBlob
	serviceName  *uint16 // Windows type: *wchar_t
}
