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

var (
	wireguardTrafficSublayer = windows.GUID{
		Data1: 1301102018,
		Data2: 18918,
		Data3: 20151,
		Data4: [8]byte{135, 32, 14, 80, 104, 75, 115, 61},
	}
	engineHandleUsed = uintptr(0)
	localAddressesV4 = [3]wtFwpV4AddrAndMask{
		{0x0A000000, 0xff000000},
		{0xAC100000, 0xfff00000},
		{0xC0A80000, 0xffff0000},
	}
	localAddressV6 = wtFwpV6AddrAndMask{[16]uint8{0xFC, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},7	}
)

func StartTrafficBlock(ifcLuid uint64, name, description string) error {
	if engineHandleUsed != 0 {
		StopTrafficBlock()
	}
	sessionDisplayData, err := createWtFwpmDisplayData0(name, description)
	if err != nil {
		return err
	}
	session := wtFwpmSession0{
		sessionKey:  *getZeroGuid(),
		displayData: *sessionDisplayData,
		flags:       fwpmSessionFlagDynamic,
	}
	result := fwpmEngineOpen0(nil, RPC_C_AUTHN_WINNT, nil, &session, unsafe.Pointer(&engineHandleUsed))
	if result != 0 {
		StopTrafficBlock()
		return os.NewSyscallError("fwpuclnt.FwpmEngineOpen0", windows.Errno(result))
	}
	res := fwpmTransactionBegin0(engineHandleUsed, 0)
	if res != 0 {
		StopTrafficBlock()
		return os.NewSyscallError("fwpuclnt.FwpmTransactionBegin0", windows.Errno(res))
	}
	err = blockTraffic(engineHandleUsed, ifcLuid, name, description)
	if err == nil {
		res = fwpmTransactionCommit0(engineHandleUsed)
		if res == 0 {
			return nil
		} else {
			return os.NewSyscallError("fwpuclnt.FwpmTransactionCommit0", windows.Errno(res))
		}
	} else {
		fwpmTransactionAbort0(engineHandleUsed)
		return err
	}
}

func blockTraffic(engineHandle uintptr, ifcLuid uint64, name, description string) error {
	sublayerDisplayData, err := createWtFwpmDisplayData0(name, description)
	if err != nil {
		return err
	}
	sublayer := wtFwpmSublayer0{
		subLayerKey: wireguardTrafficSublayer,
		displayData: *sublayerDisplayData,
		flags:       wtFwpmSublayerFlags(0),
		weight:      ^uint16(0),
	}
	res := fwpmSubLayerAdd0(engineHandleUsed, &sublayer, 0)
	if res != 0 {
		return os.NewSyscallError("fwpuclnt.FwpmSubLayerAdd0", windows.Errno(res))
	}
	appId, err := GetCurrentAppId()
	defer appId.Free()
	if err != nil {
		return err
	}
	var conditions [3]wtFwpmFilterCondition0
	conditions[0] = wtFwpmFilterCondition0{
		fieldKey:  FWPM_CONDITION_ALE_APP_ID,
		matchType: FWP_MATCH_EQUAL,
		conditionValue: wtFwpConditionValue0{
			_type: FWP_BYTE_BLOB_TYPE,
			value: uintptr(unsafe.Pointer(appId)),
		},
	}
	filterDisplayData, err := createWtFwpmDisplayData0(name, description)
	if err != nil {
		return err
	}
	filter := wtFwpmFilter0{
		filterKey:   *getZeroGuid(),
		displayData: *filterDisplayData,
		flags:       0,
		layerKey:    FWPM_LAYER_ALE_AUTH_CONNECT_V4,
		subLayerKey: wireguardTrafficSublayer,
		weight: wtFwpValue0{
			_type: FWP_UINT8,
			value: 15,
		},
		numFilterConditions: 1,
		filterCondition:     (*wtFwpmFilterCondition0)(unsafe.Pointer(&conditions[0])),
		action: wtFwpmAction0{
			_type: FWP_ACTION_PERMIT,
		},
	}
	filter.cleanup()
	filterId := uint64(0)
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	filter.cleanup()
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V6
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	filter.cleanup()
	conditions[0].fieldKey = FWPM_CONDITION_IP_LOCAL_INTERFACE
	conditions[0].conditionValue._type = FWP_UINT64
	conditions[0].conditionValue.value = uintptr(unsafe.Pointer(&ifcLuid))
	filter.weight.value = 14
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V4
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	filter.cleanup()
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V6
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	filter.cleanup()
	conditions[0].fieldKey = FWPM_CONDITION_INTERFACE_TYPE
	conditions[0].conditionValue._type = FWP_UINT32
	conditions[0].conditionValue.value = 24
	filter.weight.value = 13
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V4
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	filter.cleanup()
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V6
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	filter.cleanup()
	conditions[0].fieldKey = FWPM_CONDITION_IP_REMOTE_ADDRESS
	conditions[0].conditionValue._type = FWP_V4_ADDR_MASK
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V4
	for idx, addr := range localAddressesV4 {
		conditions[0].conditionValue.value = uintptr(unsafe.Pointer(&addr))
		filter.weight.value = uintptr(12 - idx)
		res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
		if res != 0 {
			return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
		}
		filter.cleanup()
	}
	conditions[0].conditionValue._type = FWP_V6_ADDR_MASK
	conditions[0].conditionValue.value = uintptr(unsafe.Pointer(&localAddressV6))
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V6
	filter.weight.value = 9
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	filter.cleanup()
	conditions[0].fieldKey = FWPM_CONDITION_IP_PROTOCOL
	conditions[0].conditionValue._type = FWP_UINT8
	conditions[0].conditionValue.value = 17
	conditions[1] = wtFwpmFilterCondition0{
		fieldKey:  FWPM_CONDITION_IP_REMOTE_PORT,
		matchType: FWP_MATCH_EQUAL,
		conditionValue: wtFwpConditionValue0{
			_type: FWP_UINT16,
			value: 67,
		},
	}
	conditions[2] = wtFwpmFilterCondition0{
		fieldKey:  FWPM_CONDITION_IP_LOCAL_PORT,
		matchType: FWP_MATCH_EQUAL,
		conditionValue: wtFwpConditionValue0{
			_type: FWP_UINT16,
			value: 68,
		},
	}
	filter.numFilterConditions = 3
	filter.weight.value = 9
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V4
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	filter.cleanup()
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V6
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	filter.cleanup()
	filter.numFilterConditions = 0
	filter.weight._type = FWP_EMPTY
	filter.weight.value = 0
	filter.action._type = FWP_ACTION_BLOCK
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V4
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	filter.cleanup()
	filter.layerKey = FWPM_LAYER_ALE_AUTH_CONNECT_V6
	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)
	if res != 0 {
		return os.NewSyscallError("fwpmSubLayerAdd0.FwpmFilterAdd0", windows.Errno(res))
	}
	return nil
}

func StopTrafficBlock() {
	if engineHandleUsed != 0 {
		fwpmEngineClose0(engineHandleUsed)
		engineHandleUsed = 0
	}
}
