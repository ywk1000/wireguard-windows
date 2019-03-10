/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// SID_AND_ATTRIBUTES defined in winnt.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/winnt/ns-winnt-_sid_and_attributes).
type wtSidAndAttributes struct {
	Sid        *wtSid
	Attributes uint32 // Windows type: DWORD
}
