/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "golang.org/x/sys/windows"

// Defined in fwpmu.h. 4cd62a49-59c3-4969-b7f3-bda5d32890a4
var FWPM_CONDITION_IP_LOCAL_INTERFACE = windows.GUID{
	Data1: 0x4cd62a49,
	Data2: 0x59c3,
	Data3: 0x4969,
	Data4: [8]byte{0xb7, 0xf3, 0xbd, 0xa5, 0xd3, 0x28, 0x90, 0xa4},
}

// Defined in fwpmu.h. c35a604d-d22b-4e1a-91b4-68f674ee674b
var FWPM_CONDITION_IP_REMOTE_PORT = windows.GUID{
	Data1: 0xc35a604d,
	Data2: 0xd22b,
	Data3: 0x4e1a,
	Data4: [8]byte{0x91, 0xb4, 0x68, 0xf6, 0x74, 0xee, 0x67, 0x4b},
}

// Defined in fwpmu.h. d78e1e87-8644-4ea5-9437-d809ecefc971
var FWPM_CONDITION_ALE_APP_ID = windows.GUID{
	Data1: 0xd78e1e87,
	Data2: 0x8644,
	Data3: 0x4ea5,
	Data4: [8]byte{0x94, 0x37, 0xd8, 0x09, 0xec, 0xef, 0xc9, 0x71},
}
