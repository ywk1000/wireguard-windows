/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"fmt"
	"strings"
	"testing"
)

func TestGuidHelpers(t *testing.T) {

	const guidString = "{60B32EF7-340B-4225-B198-0DA6AA350171}"

	guid, err := stringToGuid(guidString)

	if err != nil {
		t.Errorf("stringToGuid() returned an error: %v", err)
		return
	}

	str := guidToString(guid)

	str = strings.ToUpper(str)

	if str != guidString {
		t.Errorf("GUID / string conversion failed. Expected: %s; Actual: %s.", guidString, str)
	}
}

func TestGetModuleFileNameWWrapper(t *testing.T) {

	currentFile, err := GetModuleFileNameWWrapper(0)

	if err != nil {
		t.Errorf("GetModuleFileNameWWrapper() returned an error: %v", err)
		return
	}

	if len(currentFile) < 1 {
		t.Error("Current file name is an empty string although GetModuleFileNameWWrapper() has executed successfully.")
		return
	}

	fmt.Printf("Executable: %s\n", currentFile)
}

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
