/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
)

const (
	allowedInterfaceLuid = uint64(1689399632855040) // TODO: Set LUID of an interface for which DNS will be allowed
)

func TestEnforceDns(t *testing.T) {

	err := BlockDnsExceptInterface(allowedInterfaceLuid, "WireGuard DNS Block",
		"Allows DNS traffic only through wintun interface.")

	if err != nil {
		t.Errorf("BlockDnsExceptInterface() returned an error: %v", err)
	}
}
