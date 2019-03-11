/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "golang.org/x/sys/windows"

type fwpmSessionFlagsValue uint32

const (
	fwpmSessionFlagDynamic fwpmSessionFlagsValue = 0x00000001 // FWPM_SESSION_FLAG_DYNAMIC defined in fwpmtypes.h
)

func (f fwpmSessionFlagsValue) isDynamic() bool {
	return (f & fwpmSessionFlagDynamic) != 0
}

// FWPM_SESSION0 defined in fwpmtypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwpmtypes/ns-fwpmtypes-fwpm_session0_).
type wtFwpmSession0 struct {
	sessionKey           windows.GUID // Windows type: GUID
	displayData          wtFwpmDisplayData0
	flags                fwpmSessionFlagsValue // Windows type UINT32
	txnWaitTimeoutInMSec uint32
	processId            uint32 // Windows type: DWORD
	sid                  *wtSid
	username             *uint16 // Windows type: *wchar_t
	kernelMode           uint8   // Windows type: BOOL
}
