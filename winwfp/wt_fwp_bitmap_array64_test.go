/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpBitmapArray64Size(t *testing.T) {

	const actualWtFwpBitmapArray64Size = unsafe.Sizeof(wtFwpBitmapArray64{})

	if actualWtFwpBitmapArray64Size != wtFwpBitmapArray64_Size {
		t.Errorf("Size of wtFwpBitmapArray64 is %d, although %d is expected.", actualWtFwpBitmapArray64Size,
			wtFwpBitmapArray64_Size)
	}
}
