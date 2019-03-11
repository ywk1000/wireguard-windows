/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

const (
	fwpByteBlob_Size = 8

	fwpByteBlob_data_Offset = 4

	wtFwpConditionValue0_Size = 8

	wtFwpConditionValue0_uint8_Offset = 4

	wtFwpmDisplayData0_Size = 8

	wtFwpmDisplayData0_description_Offset = 4

	wtFwpmFilter0_Size = 152

	wtFwpmFilter0_displayData_Offset         = 16
	wtFwpmFilter0_flags_Offset               = 24
	wtFwpmFilter0_providerKey_Offset         = 28
	wtFwpmFilter0_providerData_Offset        = 32
	wtFwpmFilter0_layerKey_Offset            = 40
	wtFwpmFilter0_subLayerKey_Offset         = 56
	wtFwpmFilter0_weight_Offset              = 72
	wtFwpmFilter0_numFilterConditions_Offset = 80
	wtFwpmFilter0_filterCondition_Offset     = 84
	wtFwpmFilter0_action_Offset              = 88
	wtFwpmFilter0_providerContextKey_Offset  = 112
	wtFwpmFilter0_reserved_Offset            = 128
	wtFwpmFilter0_filterId_Offset            = 136
	wtFwpmFilter0_effectiveWeight_Offset     = 144

	wtFwpmFilterCondition0_Size = 28

	wtFwpmFilterCondition0_matchType_Offset      = 16
	wtFwpmFilterCondition0_conditionValue_Offset = 20

	wtFwpmSession0_Size = 48

	wtFwpmSession0_displayData_Offset          = 16
	wtFwpmSession0_flags_Offset                = 24
	wtFwpmSession0_txnWaitTimeoutInMSec_Offset = 28
	wtFwpmSession0_processId_Offset            = 32
	wtFwpmSession0_sid_Offset                  = 36
	wtFwpmSession0_username_Offset             = 40
	wtFwpmSession0_kernelMode_Offset           = 44

	wtFwpmSublayer0_Size = 44

	wtFwpmSublayer0_displayData_Offset  = 16
	wtFwpmSublayer0_flags_Offset        = 24
	wtFwpmSublayer0_providerKey_Offset  = 28
	wtFwpmSublayer0_providerData_Offset = 32
	wtFwpmSublayer0_weight_Offset       = 40

	wtFwpProvider0_Size = 40

	wtFwpProvider0_displayData_Offset  = 16
	wtFwpProvider0_flags_Offset        = 24
	wtFwpProvider0_providerData_Offset = 28
	wtFwpProvider0_serviceName_Offset  = 36

	wtFwpTokenInformation_Size = 16

	wtFwpTokenInformation_sids_Offset               = 4
	wtFwpTokenInformation_restrictedSidCount_Offset = 8
	wtFwpTokenInformation_restrictedSids_Offset     = 12

	wtFwpValue0_Size = 8

	wtFwpValue0_value_Offset = 4

	wtSecWinntAuthIdentityW_Size = 28

	wtSecWinntAuthIdentityW_UserLength_Offset     = 4
	wtSecWinntAuthIdentityW_Domain_Offset         = 8
	wtSecWinntAuthIdentityW_DomainLength_Offset   = 12
	wtSecWinntAuthIdentityW_Password_Offset       = 16
	wtSecWinntAuthIdentityW_PasswordLength_Offset = 20
	wtSecWinntAuthIdentityW_Flags_Offset          = 24

	wtSidAndAttributes_Size = 8

	wtSidAndAttributes_Attributes_Offset = 4
)
