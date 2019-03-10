/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpByteArray6Size(t *testing.T) {

	const actualWtFwpByteArray6Size = unsafe.Sizeof(wtFwpByteArray6{})

	if actualWtFwpByteArray6Size != wtFwpByteArray6_Size {
		t.Errorf("Size of wtFwpByteArray6 is %d, although %d is expected.", actualWtFwpByteArray6Size,
			wtFwpByteArray6_Size)
	}
}
