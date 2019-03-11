/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "golang.org/x/sys/windows"

// FWPM_LAYER_ALE_AUTH_CONNECT_V4 (c38d57d1-05a7-4c33-904f-7fbceee60e82) defined in fwpmu.h
var FWPM_LAYER_ALE_AUTH_CONNECT_V4 = windows.GUID{
	Data1: 0xc38d57d1,
	Data2: 0x05a7,
	Data3: 0x4c33,
	Data4: [8]byte{0x90, 0x4f, 0x7f, 0xbc, 0xee, 0xe6, 0x0e, 0x82},
}

// FWPM_LAYER_ALE_AUTH_CONNECT_V6 (4a72393b-319f-44bc-84c3-ba54dcb3b6b4) defined in fwpmu.h
var FWPM_LAYER_ALE_AUTH_CONNECT_V6 = windows.GUID{
	Data1: 0x4a72393b,
	Data2: 0x319f,
	Data3: 0x44bc,
	Data4: [8]byte{0x84, 0xc3, 0xba, 0x54, 0xdc, 0xb3, 0xb6, 0xb4},
}
