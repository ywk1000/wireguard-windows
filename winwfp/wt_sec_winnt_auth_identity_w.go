/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// SEC_WINNT_AUTH_IDENTITY_W defined in rpcdce.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/rpcdce/ns-rpcdce-_sec_winnt_auth_identity_w).
type wtSecWinntAuthIdentityW struct {
	User           *uint16 // Windows type: unsigned short
	UserLength     int32   // Windows type: long
	Domain         *uint16 // Windows type: unsigned short
	DomainLength   int32   // Windows type: long
	Password       *uint16 // Windows type: unsigned short
	PasswordLength int32   // Windows type: long
	Flags          int32   // Windows type: long
}
