/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpmFilter0Size(t *testing.T) {

	const actualWtFwpmFilter0Size = unsafe.Sizeof(wtFwpmFilter0{})

	if actualWtFwpmFilter0Size != wtFwpmFilter0_Size {
		t.Errorf("Size of wtFwpmFilter0 is %d, although %d is expected.", actualWtFwpmFilter0Size,
			wtFwpmFilter0_Size)
	}
}

func Test_wtFwpmFilter0_Offsets(t *testing.T) {

	s := wtFwpmFilter0{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.displayData)) - sp

	if offset != wtFwpmFilter0_displayData_Offset {
		t.Errorf("wtFwpmFilter0.displayData offset is %d although %d is expected", offset,
			wtFwpmFilter0_displayData_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.flags)) - sp

	if offset != wtFwpmFilter0_flags_Offset {
		t.Errorf("wtFwpmFilter0.flags offset is %d although %d is expected", offset, wtFwpmFilter0_flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.providerKey)) - sp

	if offset != wtFwpmFilter0_providerKey_Offset {
		t.Errorf("wtFwpmFilter0.providerKey offset is %d although %d is expected", offset,
			wtFwpmFilter0_providerKey_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.providerData)) - sp

	if offset != wtFwpmFilter0_providerData_Offset {
		t.Errorf("wtFwpmFilter0.providerData offset is %d although %d is expected", offset,
			wtFwpmFilter0_providerData_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.layerKey)) - sp

	if offset != wtFwpmFilter0_layerKey_Offset {
		t.Errorf("wtFwpmFilter0.layerKey offset is %d although %d is expected", offset,
			wtFwpmFilter0_layerKey_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.subLayerKey)) - sp

	if offset != wtFwpmFilter0_subLayerKey_Offset {
		t.Errorf("wtFwpmFilter0.subLayerKey offset is %d although %d is expected", offset,
			wtFwpmFilter0_subLayerKey_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.weight)) - sp

	if offset != wtFwpmFilter0_weight_Offset {
		t.Errorf("wtFwpmFilter0.weight offset is %d although %d is expected", offset,
			wtFwpmFilter0_weight_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.numFilterConditions)) - sp

	if offset != wtFwpmFilter0_numFilterConditions_Offset {
		t.Errorf("wtFwpmFilter0.numFilterConditions offset is %d although %d is expected", offset,
			wtFwpmFilter0_numFilterConditions_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.filterCondition)) - sp

	if offset != wtFwpmFilter0_filterCondition_Offset {
		t.Errorf("wtFwpmFilter0.filterCondition offset is %d although %d is expected", offset,
			wtFwpmFilter0_filterCondition_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.action)) - sp

	if offset != wtFwpmFilter0_action_Offset {
		t.Errorf("wtFwpmFilter0.action offset is %d although %d is expected", offset,
			wtFwpmFilter0_action_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.providerContextKey)) - sp

	if offset != wtFwpmFilter0_providerContextKey_Offset {
		t.Errorf("wtFwpmFilter0.providerContextKey offset is %d although %d is expected", offset,
			wtFwpmFilter0_providerContextKey_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.reserved)) - sp

	if offset != wtFwpmFilter0_reserved_Offset {
		t.Errorf("wtFwpmFilter0.reserved offset is %d although %d is expected", offset,
			wtFwpmFilter0_reserved_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.filterId)) - sp

	if offset != wtFwpmFilter0_filterId_Offset {
		t.Errorf("wtFwpmFilter0.filterId offset is %d although %d is expected", offset,
			wtFwpmFilter0_filterId_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.effectiveWeight)) - sp

	if offset != wtFwpmFilter0_effectiveWeight_Offset {
		t.Errorf("wtFwpmFilter0.effectiveWeight offset is %d although %d is expected", offset,
			wtFwpmFilter0_effectiveWeight_Offset)
		return
	}
}
