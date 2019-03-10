/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"fmt"
	"golang.org/x/sys/windows"
)

// Corresponds to SEC_WINNT_AUTH_IDENTITY_W defined in rpcdce.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/rpcdce/ns-rpcdce-_sec_winnt_auth_identity_w).
type SecWinntAuthIdentity struct {
	User     string
	Domain   string
	Password string
}

func (id *SecWinntAuthIdentity) toWtSecWinntAuthIdentityW() (*wtSecWinntAuthIdentityW, error) {

	if id == nil {
		return nil, nil
	}

	user, err := windows.UTF16PtrFromString(id.User)

	if err != nil {
		return nil, err
	}

	domain, err := windows.UTF16PtrFromString(id.Domain)

	if err != nil {
		return nil, err
	}

	password, err := windows.UTF16PtrFromString(id.Password)

	if err != nil {
		return nil, err
	}

	return &wtSecWinntAuthIdentityW{
		User:           user,
		UserLength:     int32(len(id.User)),
		Domain:         domain,
		DomainLength:   int32(len(id.Domain)),
		Password:       password,
		PasswordLength: int32(len(id.Password)),
	}, nil
}

func (id *SecWinntAuthIdentity) String() string {
	if id == nil {
		return "<nil>"
	} else {
		return fmt.Sprintf("%s\\%s", id.Domain, id.User)
	}
}
