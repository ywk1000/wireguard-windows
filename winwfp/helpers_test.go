/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
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
