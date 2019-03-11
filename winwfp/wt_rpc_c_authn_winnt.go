/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

// Defined in rpcdce.h
type wtRpcCAuthN uint32

const (
	RPC_C_AUTHN_NONE          wtRpcCAuthN = 0
	RPC_C_AUTHN_DCE_PRIVATE   wtRpcCAuthN = 1
	RPC_C_AUTHN_DCE_PUBLIC    wtRpcCAuthN = 2
	RPC_C_AUTHN_DEC_PUBLIC    wtRpcCAuthN = 4
	RPC_C_AUTHN_GSS_NEGOTIATE wtRpcCAuthN = 9
	RPC_C_AUTHN_WINNT         wtRpcCAuthN = 10
	RPC_C_AUTHN_GSS_SCHANNEL  wtRpcCAuthN = 14
	RPC_C_AUTHN_GSS_KERBEROS  wtRpcCAuthN = 16
	RPC_C_AUTHN_DPA           wtRpcCAuthN = 17
	RPC_C_AUTHN_MSN           wtRpcCAuthN = 18
	RPC_C_AUTHN_DIGEST        wtRpcCAuthN = 21
	RPC_C_AUTHN_KERNEL        wtRpcCAuthN = 20
	RPC_C_AUTHN_NEGO_EXTENDER wtRpcCAuthN = 30
	RPC_C_AUTHN_PKU2U         wtRpcCAuthN = 31
	RPC_C_AUTHN_LIVE_SSP      wtRpcCAuthN = 32
	RPC_C_AUTHN_LIVEXP_SSP    wtRpcCAuthN = 35
	RPC_C_AUTHN_CLOUD_AP      wtRpcCAuthN = 36
	RPC_C_AUTHN_MSONLINE      wtRpcCAuthN = 82
	RPC_C_AUTHN_MQ            wtRpcCAuthN = 100
	RPC_C_AUTHN_DEFAULT       wtRpcCAuthN = 0xFFFFFFFF
)
