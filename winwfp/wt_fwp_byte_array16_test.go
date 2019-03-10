/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpByteArray16Size(t *testing.T) {

	const actualWtFwpByteArray16Size = unsafe.Sizeof(wtFwpByteArray16{})

	if actualWtFwpByteArray16Size != wtFwpByteArray16_Size {
		t.Errorf("Size of wtFwpByteArray16 is %d, although %d is expected.", actualWtFwpByteArray16Size,
			wtFwpByteArray16_Size)
	}
}
