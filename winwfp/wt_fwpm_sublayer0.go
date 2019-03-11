/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "golang.org/x/sys/windows"

type wtFwpmSublayerFlags uint32

const (
	fwpmSublayerFlagPersistent wtFwpmSublayerFlags = 0x00000001 // FWPM_SUBLAYER_FLAG_PERSISTENT defined in fwpmtypes.h
)

// FWPM_SUBLAYER0 defined in fwpmtypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwpmtypes/ns-fwpmtypes-fwpm_sublayer0_)
type wtFwpmSublayer0 struct {
	subLayerKey  windows.GUID // Windows type: GUID
	displayData  wtFwpmDisplayData0
	flags        wtFwpmSublayerFlags
	providerKey  *windows.GUID // Windows type: *GUID
	providerData FwpByteBlob
	weight       uint16
}
