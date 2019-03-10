/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"testing"
	"unsafe"
)

func TestWtFwpmSession0Size(t *testing.T) {

	const actualWtFwpmSession0Size = unsafe.Sizeof(wtFwpmSession0{})

	if actualWtFwpmSession0Size != wtFwpmSession0_Size {
		t.Errorf("Size of wtFwpmSession0 is %d, although %d is expected.", actualWtFwpmSession0Size,
			wtFwpmSession0_Size)
	}
}

func TestWtFwpmSession0Offsets(t *testing.T) {

	s := wtFwpmSession0{}
	sp := uintptr(unsafe.Pointer(&s))

	offset := uintptr(unsafe.Pointer(&s.displayData)) - sp

	if offset != wtFwpmSession0_displayData_Offset {
		t.Errorf("wtFwpmSession0.displayData offset is %d although %d is expected", offset,
			wtFwpmSession0_displayData_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.flags)) - sp

	if offset != wtFwpmSession0_flags_Offset {
		t.Errorf("wtFwpmSession0.flags offset is %d although %d is expected", offset, wtFwpmSession0_flags_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.txnWaitTimeoutInMSec)) - sp

	if offset != wtFwpmSession0_txnWaitTimeoutInMSec_Offset {
		t.Errorf("wtFwpmSession0.txnWaitTimeoutInMSec offset is %d although %d is expected", offset,
			wtFwpmSession0_txnWaitTimeoutInMSec_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.processId)) - sp

	if offset != wtFwpmSession0_processId_Offset {
		t.Errorf("wtFwpmSession0.processId offset is %d although %d is expected", offset,
			wtFwpmSession0_processId_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.sid)) - sp

	if offset != wtFwpmSession0_sid_Offset {
		t.Errorf("wtFwpmSession0.sid offset is %d although %d is expected", offset, wtFwpmSession0_sid_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.username)) - sp

	if offset != wtFwpmSession0_username_Offset {
		t.Errorf("wtFwpmSession0.username offset is %d although %d is expected", offset,
			wtFwpmSession0_username_Offset)
		return
	}

	offset = uintptr(unsafe.Pointer(&s.kernelMode)) - sp

	if offset != wtFwpmSession0_kernelMode_Offset {
		t.Errorf("wtFwpmSession0.kernelMode offset is %d although %d is expected", offset,
			wtFwpmSession0_kernelMode_Offset)
		return
	}
}
