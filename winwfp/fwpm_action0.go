/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "golang.org/x/sys/windows"

// FWPM_ACTION0 defined in fwpmtypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwpmtypes/ns-fwpmtypes-fwpm_action0_)
type wtFwpmAction0 struct {
	_type      FwpActionType
	filterType windows.GUID // Windows type: GUID
}
