/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"golang.org/x/sys/windows"
	"os"
	"unsafe"
)

// This is a random GUID, but once we launch we should not change it.
// {D32F19B0-AD53-4263-88A5-E772FC65DC3F}
var wireguardDnsSublayer = windows.GUID{
	Data1: 3543079344,
	Data2: 44371,
	Data3: 16995,
	Data4: [8]byte{136, 165, 231, 114, 252, 101, 220, 63},
}

func BlockDnsExceptInterface(ifcLuid uint64, name, description string) error {

	// Creating session

	sessionDisplayData, err := createWtFwpmDisplayData0(name, description)

	if err != nil {
		return err
	}

	session := wtFwpmSession0{
		sessionKey:  *getZeroGuid(),
		displayData: *sessionDisplayData,
		flags:       fwpmSessionFlagDynamic,
	}

	// Opening engine

	engineHandle := uintptr(0)

	result := fwpmEngineOpen0(nil, RPC_C_AUTHN_WINNT, nil, &session, unsafe.Pointer(&engineHandle))

	if engineHandle != 0 {
		defer fwpmEngineClose0(engineHandle)
	}

	if result != 0 {
		return os.NewSyscallError("fwpuclnt.FwpmEngineOpen0", windows.Errno(result))
	}

	// Creating sublayer

	sublayerDisplayData, err := createWtFwpmDisplayData0("WireGuard DNS block", "")

	if err != nil {
		return err
	}

	sublayer := wtFwpmSublayer0{
		subLayerKey: wireguardDnsSublayer,
		displayData: *sublayerDisplayData,
		flags:       wtFwpmSublayerFlags(0),
		weight:      ^uint16(0),
	}

	res := fwpmSubLayerAdd0(engineHandle, &sublayer, 0)

	if res != 0 {
		return os.NewSyscallError("fwpuclnt.FwpmSubLayerAdd0", windows.Errno(res))
	}

	// The first thing to do is to allow DNS access through any interface to this process (our app).

	// Get app ID

	appId, err := GetCurrentAppId()
	defer appId.Free()

	if err != nil {
		return err
	}

	// Define filter

	conditions2 := [2]wtFwpmFilterCondition0{
		{
			fieldKey:  FWPM_CONDITION_IP_REMOTE_PORT,
			matchType: FWP_MATCH_EQUAL,
			conditionValue: wtFwpConditionValue0{
				_type: FWP_UINT16,
				value: 53,
			},
		},
		{
			fieldKey:  FWPM_CONDITION_ALE_APP_ID,
			matchType: FWP_MATCH_EQUAL,
			conditionValue: wtFwpConditionValue0{
				_type: FWP_BYTE_BLOB_TYPE,
				value: uintptr(unsafe.Pointer(appId)),
			},
		}}

	filter := wtFwpmFilter0{
		filterKey:   *getZeroGuid(),
		displayData: *sessionDisplayData,
		flags:       0,
		layerKey:    FWPM_LAYER_ALE_AUTH_CONNECT_V4,
		subLayerKey: wireguardDnsSublayer,
		weight: wtFwpValue0{
			_type: FWP_UINT8,
			value: 15,
		},
		numFilterConditions: 2,
		filterCondition:     (*wtFwpmFilterCondition0)(unsafe.Pointer(&conditions2[0])),
		action: wtFwpmAction0{
			_type: FWP_ACTION_PERMIT,
		},
	}

	filter.cleanup()

	// Adding the filter

	filterId := uint64(0)

	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)

	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}

	filter.cleanup()

	// Creating and adding identical IPv6 filter

	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V6

	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)

	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}

	filter.cleanup()

	// Now we need to allow DNS through the allowed interface

	// We'll reuse the first condition, and rewrite the second:
	conditions2[1].fieldKey = FWPM_CONDITION_IP_LOCAL_INTERFACE
	conditions2[1].conditionValue._type = FWP_UINT64
	conditions2[1].conditionValue.value = uintptr(unsafe.Pointer(&ifcLuid))

	filter.weight.value = 14 // Second filter has smaller weight than the first one
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V4

	// Adding the filter

	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)

	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}

	filter.cleanup()

	// Creating and adding identical IPv6 filter

	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V6

	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)

	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}

	filter.cleanup()

	// Finally we can add "deny all DNS" rule at the bottom, with the lowest priority. For this rule we'll use only the
	// first condition.
	filter.numFilterConditions = 1  // Use only the first condition
	filter.weight._type = FWP_EMPTY // Lowest weight
	filter.weight.value = 0
	filter.action._type = FWP_ACTION_BLOCK // Block
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V4

	// Adding the filter

	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)

	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}

	filter.cleanup()

	// Creating and adding identical IPv6 filter

	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V6

	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)

	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}

	return nil
}
