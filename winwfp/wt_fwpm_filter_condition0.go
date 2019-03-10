/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "golang.org/x/sys/windows"

// FWPM_FILTER_CONDITION0 defined in fwpmtypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwpmtypes/ns-fwpmtypes-fwpm_filter_condition0_).
type wtFwpmFilterCondition0 struct {
	fieldKey       windows.GUID // Windows type: GUID
	matchType      FwpMatchType
	conditionValue wtFwpConditionValue0
}
