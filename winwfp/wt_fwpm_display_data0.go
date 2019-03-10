/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "golang.org/x/sys/windows"

// FWPM_DISPLAY_DATA0 defined in fwptypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwptypes/ns-fwptypes-fwpm_display_data0_).
type wtFwpmDisplayData0 struct {
	name        *uint16 // Windows type: *wchar_t
	description *uint16 // Windows type: *wchar_t
}

func createWtFwpmDisplayData0(name, description string) (*wtFwpmDisplayData0, error) {

	namePtr, err := windows.UTF16PtrFromString(name)

	if err != nil {
		return nil, err
	}

	descriptionPtr, err := windows.UTF16PtrFromString(description)

	if err != nil {
		return nil, err
	}

	return &wtFwpmDisplayData0{
		name:        namePtr,
		description: descriptionPtr,
	}, nil
}

func (dd *wtFwpmDisplayData0) toFwpmDisplayData() *FwpmDisplayData {

	if dd == nil {
		return nil
	}

	return &FwpmDisplayData{
		Name:        wcharToString(dd.name),
		Description: wcharToString(dd.description),
	}
}
