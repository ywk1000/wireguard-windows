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

// Defined in fwpmu.h. b235ae9a-1d64-49b8-a44c-5ff3d9095045
var FWPM_CONDITION_IP_REMOTE_ADDRESS = windows.GUID{
	Data1: 0xb235ae9a,
	Data2: 0x1d64,
	Data3: 0x49b8,
	Data4: [8]byte{0xa4, 0x4c, 0x5f, 0xf3, 0xd9, 0x09, 0x50, 0x45},
}

// Defined in fwpmu.h. daf8cd14-e09e-4c93-a5ae-c5c13b73ffca
var FWPM_CONDITION_INTERFACE_TYPE = windows.GUID{
	Data1: 0xdaf8cd14,
	Data2: 0xe09e,
	Data3: 0x4c93,
	Data4: [8]byte{0xa5, 0xae, 0xc5, 0xc1, 0x3b, 0x73, 0xff, 0xca},
}

// Defined in fwpmu.h. 3971ef2b-623e-4f9a-8cb1-6e79b806b9a7
var FWPM_CONDITION_IP_PROTOCOL = windows.GUID{
	Data1: 0x3971ef2b,
	Data2: 0x623e,
	Data3: 0x4f9a,
	Data4: [8]byte{0x8c, 0xb1, 0x6e, 0x79, 0xb8, 0x06, 0xb9, 0xa7},
}

// Defined in fwpmu.h. 0c1ba1af-5765-453f-af22-a8f791ac775b
var FWPM_CONDITION_IP_LOCAL_PORT = windows.GUID{
	Data1: 0x0c1ba1af,
	Data2: 0x5765,
	Data3: 0x453f,
	Data4: [8]byte{0xaf, 0x22, 0xa8, 0xf7, 0x91, 0xac, 0x77, 0x5b},
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
