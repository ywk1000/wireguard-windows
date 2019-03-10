/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// FWP_TOKEN_INFORMATION defined in fwptypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwptypes/ns-fwptypes-fwp_token_information)
type wtFwpTokenInformation struct {
	sidCount           uint32              // Windows type: ULONG
	sids               *wtSidAndAttributes // Windows type: PSID_AND_ATTRIBUTES
	restrictedSidCount uint32              // Windows type: ULONG
	restrictedSids     *wtSidAndAttributes // Windows type: PSID_AND_ATTRIBUTES
}
