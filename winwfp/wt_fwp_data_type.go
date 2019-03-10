/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winwfp

import "fmt"

// FWP_DATA_TYPE defined in fwptypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwptypes/ne-fwptypes-fwp_data_type_)
type wtFwpDataType uint

const (
	fwpEmpty                      wtFwpDataType = 0                                 // FWP_EMPTY
	fwpUint8                      wtFwpDataType = fwpEmpty + 1                      // FWP_UINT8
	fwpUint16                     wtFwpDataType = fwpUint8 + 1                      // FWP_UINT16
	fwpUint32                     wtFwpDataType = fwpUint16 + 1                     // FWP_UINT32
	fwpUint64                     wtFwpDataType = fwpUint32 + 1                     // FWP_UINT64
	fwpInt8                       wtFwpDataType = fwpUint64 + 1                     // FWP_INT8
	fwpInt16                      wtFwpDataType = fwpInt8 + 1                       // FWP_INT16
	fwpInt32                      wtFwpDataType = fwpInt16 + 1                      // FWP_INT32
	fwpInt64                      wtFwpDataType = fwpInt32 + 1                      // FWP_INT64
	fwpFloat                      wtFwpDataType = fwpInt64 + 1                      // FWP_FLOAT
	fwpDouble                     wtFwpDataType = fwpFloat + 1                      // FWP_DOUBLE
	fwpByteArray16Type            wtFwpDataType = fwpDouble + 1                     // FWP_BYTE_ARRAY16_TYPE
	fwpByteBlobType               wtFwpDataType = fwpByteArray16Type + 1            // FWP_BYTE_BLOB_TYPE
	fwpSid                        wtFwpDataType = fwpByteBlobType + 1               // FWP_SID
	fwpSecurityDescriptorType     wtFwpDataType = fwpSid + 1                        // FWP_SECURITY_DESCRIPTOR_TYPE
	fwpTokenInformationType       wtFwpDataType = fwpSecurityDescriptorType + 1     // FWP_TOKEN_INFORMATION_TYPE
	fwpTokenAccessInformationType wtFwpDataType = fwpTokenInformationType + 1       // FWP_TOKEN_ACCESS_INFORMATION_TYPE
	fwpUnicodeStringType          wtFwpDataType = fwpTokenAccessInformationType + 1 // FWP_UNICODE_STRING_TYPE
	fwpByteArray6Type             wtFwpDataType = fwpUnicodeStringType + 1          // FWP_BYTE_ARRAY6_TYPE
	fwpBitmapIndexType            wtFwpDataType = fwpByteArray6Type + 1             // FWP_BITMAP_INDEX_TYPE
	fwpBitmapArray64Type          wtFwpDataType = fwpBitmapIndexType + 1            // FWP_BITMAP_ARRAY64_TYPE
	fwpSingleDataTypeMax          wtFwpDataType = 0xff                              // FWP_SINGLE_DATA_TYPE_MAX
	fwpV4AddrMask                 wtFwpDataType = fwpSingleDataTypeMax + 1          // FWP_V4_ADDR_MASK
	fwpV6AddrMask                 wtFwpDataType = fwpV4AddrMask + 1                 // FWP_V6_ADDR_MASK
	fwpRangeType                  wtFwpDataType = fwpV6AddrMask + 1                 // FWP_RANGE_TYPE
	fwpDataTypeMax                wtFwpDataType = fwpRangeType + 1                  // FWP_DATA_TYPE_MAX
)

func (dt wtFwpDataType) String() string {
	switch dt {
	case fwpEmpty:
		return "fwpEmpty"
	case fwpUint8:
		return "fwpUint8"
	case fwpUint16:
		return "fwpUint16"
	case fwpUint32:
		return "fwpUint32"
	case fwpUint64:
		return "fwpUint64"
	case fwpInt8:
		return "fwpInt8"
	case fwpInt16:
		return "fwpInt16"
	case fwpInt32:
		return "fwpInt32"
	case fwpInt64:
		return "fwpInt64"
	case fwpFloat:
		return "fwpFloat"
	case fwpDouble:
		return "fwpDouble"
	case fwpByteArray16Type:
		return "fwpByteArray16Type"
	case fwpByteBlobType:
		return "fwpByteBlobType"
	case fwpSid:
		return "fwpSid"
	case fwpSecurityDescriptorType:
		return "fwpSecurityDescriptorType"
	case fwpTokenInformationType:
		return "fwpTokenInformationType"
	case fwpTokenAccessInformationType:
		return "fwpTokenAccessInformationType"
	case fwpUnicodeStringType:
		return "fwpUnicodeStringType"
	case fwpByteArray6Type:
		return "fwpByteArray6Type"
	case fwpBitmapIndexType:
		return "fwpBitmapIndexType"
	case fwpBitmapArray64Type:
		return "fwpBitmapArray64Type"
	case fwpSingleDataTypeMax:
		return "fwpSingleDataTypeMax"
	case fwpV4AddrMask:
		return "fwpV4AddrMask"
	case fwpV6AddrMask:
		return "fwpV6AddrMask"
	case fwpRangeType:
		return "fwpRangeType"
	case fwpDataTypeMax:
		return "fwpDataTypeMax"
	default:
		return fmt.Sprintf("FwpDataType_UNKNOWN(%d)", dt)
	}
}
