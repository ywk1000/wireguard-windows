/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// https://docs.microsoft.com/en-us/windows/desktop/api/sddl/nf-sddl-convertsidtostringsidw
//sys	convertSidToStringSidW(Sid *wtSid, StringSid *uint16) (result uint8) = fwpuclnt.ConvertSidToStringSidW

// https://docs.microsoft.com/en-us/windows/desktop/api/securitybaseapi/nf-securitybaseapi-allocateandinitializesid
//sys	allocateAndInitializeSid(pIdentifierAuthority *SidIdentifierAuthority, nSubAuthorityCount uint8, nSubAuthority0 uint32, nSubAuthority1 uint32, nSubAuthority2 uint32, nSubAuthority3 uint32, nSubAuthority4 uint32, nSubAuthority5 uint32, nSubAuthority6 uint32, nSubAuthority7 uint32, pSid unsafe.Pointer) (result uint8) = fwpuclnt.AllocateAndInitializeSid

// https://docs.microsoft.com/en-us/windows/desktop/api/securitybaseapi/nf-securitybaseapi-freesid
//sys	freeSid(Sid *wtSid) (result uint8) = fwpuclnt.FreeSid

// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmu/nf-fwpmu-fwpmengineopen0
//sys	fwpmEngineOpen0(serverName *uint16, authnService wtRpcCAuthN, authIdentity *wtSecWinntAuthIdentityW, session *wtFwpmSession0, engineHandle unsafe.Pointer) (result uint) = fwpuclnt.FwpmEngineOpen0

// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmu/nf-fwpmu-fwpmengineclose0
//sys	fwpmEngineClose0(engineHandle uintptr) (result uint32) = fwpuclnt.FwpmEngineClose0

// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmu/nf-fwpmu-fwpmsublayeradd0
//sys	fwpmSubLayerAdd0(engineHandle uintptr, subLayer *wtFwpmSublayer0, sd uintptr) (result uint32) = fwpuclnt.FwpmSubLayerAdd0

// https://docs.microsoft.com/en-us/windows/desktop/api/libloaderapi/nf-libloaderapi-getmodulefilenamew
//sys	getModuleFileNameW(hModule uintptr, lpFilename *uint16, nSize uint32) (result uint32) = Kernel32.GetModuleFileNameW

// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmu/nf-fwpmu-fwpmgetappidfromfilename0
//sys	fwpmGetAppIdFromFileName0(fileName *uint16, appId unsafe.Pointer) (result uint32) = fwpuclnt.FwpmGetAppIdFromFileName0

// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmu/nf-fwpmu-fwpmfreememory0
//sys	fwpmFreeMemory0(p unsafe.Pointer) = fwpuclnt.FwpmFreeMemory0

// https://docs.microsoft.com/en-us/windows/desktop/api/fwpmu/nf-fwpmu-fwpmfilteradd0
//sys	fwpmFilterAdd0(engineHandle uintptr, filter *wtFwpmFilter0, sd uintptr, id *uint64) (result uint32) = fwpuclnt.FwpmFilterAdd0
