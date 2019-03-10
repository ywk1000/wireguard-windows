/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

const (
	wtFwpByteBlob_Size = 16

	wtFwpByteBlob_data_Offset = 8

	wtFwpConditionValue0_Size = 16

	wtFwpConditionValue0_uint8_Offset = 8

	wtFwpmDisplayData0_Size = 16

	wtFwpmDisplayData0_description_Offset = 8

	wtFwpmFilter0_Size = 200

	wtFwpmFilter0_displayData_Offset         = 16
	wtFwpmFilter0_flags_Offset               = 32
	wtFwpmFilter0_providerKey_Offset         = 40
	wtFwpmFilter0_providerData_Offset        = 48
	wtFwpmFilter0_layerKey_Offset            = 64
	wtFwpmFilter0_subLayerKey_Offset         = 80
	wtFwpmFilter0_weight_Offset              = 96
	wtFwpmFilter0_numFilterConditions_Offset = 112
	wtFwpmFilter0_filterCondition_Offset     = 120
	wtFwpmFilter0_action_Offset              = 128
	wtFwpmFilter0_providerContextKey_Offset  = 152
	wtFwpmFilter0_reserved_Offset            = 168
	wtFwpmFilter0_filterId_Offset            = 176
	wtFwpmFilter0_effectiveWeight_Offset     = 184

	wtFwpmFilterCondition0_Size = 40

	wtFwpmFilterCondition0_matchType_Offset      = 16
	wtFwpmFilterCondition0_conditionValue_Offset = 24

	wtFwpmSession0_Size = 72

	wtFwpmSession0_displayData_Offset          = 16
	wtFwpmSession0_flags_Offset                = 32
	wtFwpmSession0_txnWaitTimeoutInMSec_Offset = 36
	wtFwpmSession0_processId_Offset            = 40
	wtFwpmSession0_sid_Offset                  = 48
	wtFwpmSession0_username_Offset             = 56
	wtFwpmSession0_kernelMode_Offset           = 64

	wtFwpmSublayer0_Size = 72

	wtFwpmSublayer0_displayData_Offset  = 16
	wtFwpmSublayer0_flags_Offset        = 32
	wtFwpmSublayer0_providerKey_Offset  = 40
	wtFwpmSublayer0_providerData_Offset = 48
	wtFwpmSublayer0_weight_Offset       = 64

	wtFwpProvider0_Size = 64

	wtFwpProvider0_displayData_Offset  = 16
	wtFwpProvider0_flags_Offset        = 32
	wtFwpProvider0_providerData_Offset = 40
	wtFwpProvider0_serviceName_Offset  = 56

	wtFwpTokenInformation_Size = 32

	wtFwpTokenInformation_sids_Offset               = 8
	wtFwpTokenInformation_restrictedSidCount_Offset = 16
	wtFwpTokenInformation_restrictedSids_Offset     = 24

	wtFwpValue0_Size = 16

	wtFwpValue0_value_Offset = 8

	wtSecWinntAuthIdentityW_Size = 48

	wtSecWinntAuthIdentityW_UserLength_Offset     = 8
	wtSecWinntAuthIdentityW_Domain_Offset         = 16
	wtSecWinntAuthIdentityW_DomainLength_Offset   = 24
	wtSecWinntAuthIdentityW_Password_Offset       = 32
	wtSecWinntAuthIdentityW_PasswordLength_Offset = 40
	wtSecWinntAuthIdentityW_Flags_Offset          = 44

	wtSidAndAttributes_Size = 16

	wtSidAndAttributes_Attributes_Offset = 8
)
