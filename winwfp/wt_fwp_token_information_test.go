/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpTokenInformationSize(t *testing.T) {

	const actualWtFwpTokenInformationSize = unsafe.Sizeof(wtFwpTokenInformation{})

	if actualWtFwpTokenInformationSize != wtFwpTokenInformation_Size {
		t.Errorf("Size of wtFwpTokenInformation is %d, although %d is expected.", actualWtFwpTokenInformationSize,
			wtFwpTokenInformation_Size)
	}
}

func TestWtFwpTokenInformationOffsets(t *testing.T) {

	s := wtFwpTokenInformation{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.sids)) - sp

	if offset != wtFwpTokenInformation_sids_Offset {
		t.Errorf("wtFwpTokenInformation.sids offset is %d although %d is expected", offset,
			wtFwpTokenInformation_sids_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.restrictedSidCount)) - sp

	if offset != wtFwpTokenInformation_restrictedSidCount_Offset {
		t.Errorf("wtFwpTokenInformation.restrictedSidCount offset is %d although %d is expected", offset,
			wtFwpTokenInformation_restrictedSidCount_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.restrictedSids)) - sp

	if offset != wtFwpTokenInformation_restrictedSids_Offset {
		t.Errorf("wtFwpTokenInformation.restrictedSids offset is %d although %d is expected", offset,
			wtFwpTokenInformation_restrictedSids_Offset)
		return
	}
}
