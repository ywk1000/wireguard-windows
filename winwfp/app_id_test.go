/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
)

func TestGetCurrentAppId(t *testing.T) {

	appId, err := GetCurrentAppId()

	if err != nil {
		t.Errorf("GetCurrentAppId() returned an error: %v", err)
		return
	}

	if appId == nil {
		t.Error("GetCurrentAppId() returned nil.")
	}
}
