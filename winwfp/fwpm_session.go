/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"fmt"
	"golang.org/x/sys/windows"
)

// Corresponds to FWPM_SESSION0 defined in fwpmtypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwpmtypes/ns-fwpmtypes-fwpm_session0_).
type FwpmSession struct {
	SessionKey           windows.GUID
	DisplayData          FwpmDisplayData
	Dynamic              bool
	TxnWaitTimeoutInMSec uint32
	ProcessId            uint32
	Sid                  *Sid
	Username             string
	KernelMode           bool
}

func (s *FwpmSession) toWtFwpmSession0() (*wtFwpmSession0, error) {

	if s == nil {
		return nil, nil
	}

	displayData, err := s.DisplayData.toWtFwpmDisplayData()

	if err != nil {
		return nil, err
	}

	username, err := windows.UTF16PtrFromString(s.Username)

	if err != nil {
		return nil, err
	}

	wtsid, err := s.Sid.toWtSid()

	if err != nil {
		return nil, err
	}

	wt := wtFwpmSession0{
		sessionKey: s.SessionKey,
		displayData: *displayData,
		flags: fwpmSessionFlagsValue(0),
		txnWaitTimeoutInMSec: s.TxnWaitTimeoutInMSec,
		processId: s.ProcessId,
		sid: wtsid,
		username: username,
		kernelMode: boolToUint8(s.KernelMode),
	}

	if s.Dynamic {
		wt.flags |= fwpmSessionFlagDynamic
	}

	return &wt, nil
}

func (s *FwpmSession) String() string {

	if s == nil {
		return "<nil>"
	}

	return fmt.Sprintf(`SessionKey: %s
DisplayData: %s
Dynamic: %v
TxnWaitTimeoutInMSec: %d
ProcessId: %d
Sid: %s
Username: %s
KernelMode: %v`, guidToString(&s.SessionKey), s.DisplayData.String(), s.Dynamic, s.TxnWaitTimeoutInMSec, s.ProcessId,
		s.Sid.String(), s.Username, s.KernelMode)
}
