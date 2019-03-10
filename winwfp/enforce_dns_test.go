/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import (
	"fmt"
	"golang.org/x/sys/windows"
	"testing"
	"unsafe"
)

const (
	allowedInterfaceLuid = uint64(1689399632855040) // TODO: Set LUID of an interface for which DNS will be allowed
)

func TestEnforceDns(t *testing.T) {

	// Creating session

	sessionDisplayData, err := createWtFwpmDisplayData0("WireGuard DNS session", "")

	if err != nil {
		t.Errorf("createWtFwpmDisplayData0() returned an error: %v", err)
		return
	}

	session := wtFwpmSession0{
		sessionKey:  windows.GUID{Data1: 0, Data2: 0, Data3: 0, Data4: [8]byte{0, 0, 0, 0, 0, 0, 0, 0}},
		displayData: *sessionDisplayData,
		flags:       fwpmSessionFlagDynamic,
	}

	// Opening engine

	engineHandle := uintptr(0)

	result := fwpmEngineOpen0(nil, rpcCAuthN_WINNT, nil, &session, unsafe.Pointer(&engineHandle))

	if engineHandle != 0 {
		defer func() {
			r := fwpmEngineClose0(engineHandle)

			if r != 0 {
				t.Errorf("fwpmEngineClose0() returned %d.", r)
			}
		}()
	}

	if result != 0 {
		t.Errorf("fwpmEngineOpen0() returned %d.", result)
		return
	}

	if engineHandle == 0 {
		t.Error("fwpmEngineOpen0() executed successfully, but engineHandle is still 0.")
		return
	}

	// Creating sublayer

	wireguardDnsSublayer, err := stringToGuid("{D32F19B0-AD53-4263-88A5-E772FC65DC3F}")

	if err != nil {
		t.Errorf("stringToGuid() returned an error: %v", err)
		return
	}

	sublayerDisplayData, err := createWtFwpmDisplayData0("WireGuard DNS block", "")

	if err != nil {
		t.Errorf("createWtFwpmDisplayData0() returned an error: %v", err)
		return
	}

	sublayer := wtFwpmSublayer0{
		subLayerKey: *wireguardDnsSublayer,
		displayData: *sublayerDisplayData,
		flags:       wtFwpmSublayerFlags(0),
		weight:      ^uint16(0),
	}

	res := fwpmSubLayerAdd0(engineHandle, &sublayer, 0)

	if res != 0 {
		t.Errorf("fwpmSubLayerAdd0() returned %d.", res)
		return
	}

	// The first thing to do is to allow DNS access to this process (our app).

	// Get app ID

	appId, err := getCurrentAppId()

	if err != nil {
		t.Errorf("getCurrentAppId() returned an error: %v", err)
		return
	}

	// Define filter

	conditions2 := [2]wtFwpmFilterCondition0{
		{
			fieldKey:  *GetFwpmConditionIpRemotePort(),
			matchType: FWP_MATCH_EQUAL,
			conditionValue: wtFwpConditionValue0{
				_type: fwpUint16,
				value: 53,
			},
		},
		{
			fieldKey:  *GetFwpmConditionAleAppId(),
			matchType: FWP_MATCH_EQUAL,
			conditionValue: wtFwpConditionValue0{
				_type: fwpByteBlobType,
				value: uintptr(unsafe.Pointer(appId)),
			},
		}}

	filter := wtFwpmFilter0{
		filterKey:   windows.GUID{Data1: 0, Data2: 0, Data3: 0, Data4: [8]byte{0, 0, 0, 0, 0, 0, 0, 0}},
		displayData: *sessionDisplayData,
		flags:       0,
		layerKey:    *GetFwpmLayerAleAuthConnectV4(),
		subLayerKey: *wireguardDnsSublayer,
		weight: wtFwpValue0{
			_type: fwpUint8,
			value: 15,
		},
		numFilterConditions: 2,
		filterCondition:     (*wtFwpmFilterCondition0)(unsafe.Pointer(&conditions2[0])),
		action: wtFwpmAction0{
			_type: FWP_ACTION_PERMIT,
		},
	}

	// Adding the filter

	filterId := uint64(0)

	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)

	if res != 0 {
		t.Errorf("fwpmFilterAdd0() returned %d.", res)
		return
	}

	fmt.Printf("IPv4 filter added: %d\n", filterId)

	// Creating and adding identical IPv6 filter

	filter.layerKey = *GetFwpmLayerAleAuthConnectV6()

	res = fwpmFilterAdd0(engineHandle, &filter, 0, &filterId)

	if res != 0 {
		t.Errorf("fwpmFilterAdd0() returned %d.", res)
		return
	}

	fmt.Printf("IPv6 filter added: %d\n", filterId)
}
