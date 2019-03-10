/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"fmt"
)

type FwpmDisplayData struct {
	Name        string
	Description string
}

func (dd *FwpmDisplayData) toWtFwpmDisplayData() (*wtFwpmDisplayData0, error) {
	if dd == nil {
		return nil, nil
	} else {
		return createWtFwpmDisplayData0(dd.Name, dd.Description)
	}
}

func (dd *FwpmDisplayData) String() string {
	if dd == nil {
		return "<nil>"
	} else {
		return fmt.Sprintf("Name: %s; Description: %s", dd.Name, dd.Description)
	}
}
