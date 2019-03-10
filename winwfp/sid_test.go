/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestSidIdentifierAuthoritySize(t *testing.T) {

	const actualSidIdentifierAuthoritySize = unsafe.Sizeof(SidIdentifierAuthority{})

	if actualSidIdentifierAuthoritySize != sidIdentifierAuthority_Size {
		t.Errorf("Size of SidIdentifierAuthority is %d, although %d is expected.",
			actualSidIdentifierAuthoritySize, sidIdentifierAuthority_Size)
	}
}
