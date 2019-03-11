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
	FWP_EMPTY                         wtFwpDataType = 0
	FWP_UINT8                         wtFwpDataType = FWP_EMPTY + 1
	FWP_UINT16                        wtFwpDataType = FWP_UINT8 + 1
	FWP_UINT32                        wtFwpDataType = FWP_UINT16 + 1
	FWP_UINT64                        wtFwpDataType = FWP_UINT32 + 1
	FWP_INT8                          wtFwpDataType = FWP_UINT64 + 1
	FWP_INT16                         wtFwpDataType = FWP_INT8 + 1
	FWP_INT32                         wtFwpDataType = FWP_INT16 + 1
	FWP_INT64                         wtFwpDataType = FWP_INT32 + 1
	FWP_FLOAT                         wtFwpDataType = FWP_INT64 + 1
	FWP_DOUBLE                        wtFwpDataType = FWP_FLOAT + 1
	FWP_BYTE_ARRAY16_TYPE             wtFwpDataType = FWP_DOUBLE + 1
	FWP_BYTE_BLOB_TYPE                wtFwpDataType = FWP_BYTE_ARRAY16_TYPE + 1
	FWP_SID                           wtFwpDataType = FWP_BYTE_BLOB_TYPE + 1
	FWP_SECURITY_DESCRIPTOR_TYPE      wtFwpDataType = FWP_SID + 1
	FWP_TOKEN_INFORMATION_TYPE        wtFwpDataType = FWP_SECURITY_DESCRIPTOR_TYPE + 1
	FWP_TOKEN_ACCESS_INFORMATION_TYPE wtFwpDataType = FWP_TOKEN_INFORMATION_TYPE + 1
	FWP_UNICODE_STRING_TYPE           wtFwpDataType = FWP_TOKEN_ACCESS_INFORMATION_TYPE + 1
	FWP_BYTE_ARRAY6_TYPE              wtFwpDataType = FWP_UNICODE_STRING_TYPE + 1
	FWP_BITMAP_INDEX_TYPE             wtFwpDataType = FWP_BYTE_ARRAY6_TYPE + 1
	FWP_BITMAP_ARRAY64_TYPE           wtFwpDataType = FWP_BITMAP_INDEX_TYPE + 1
	FWP_SINGLE_DATA_TYPE_MAX          wtFwpDataType = 0xff
	FWP_V4_ADDR_MASK                  wtFwpDataType = FWP_SINGLE_DATA_TYPE_MAX + 1
	FWP_V6_ADDR_MASK                  wtFwpDataType = FWP_V4_ADDR_MASK + 1
	FWP_RANGE_TYPE                    wtFwpDataType = FWP_V6_ADDR_MASK + 1
	FWP_DATA_TYPE_MAX                 wtFwpDataType = FWP_RANGE_TYPE + 1
)

func (dt wtFwpDataType) String() string {
	switch dt {
	case FWP_EMPTY:
		return "FWP_EMPTY"
	case FWP_UINT8:
		return "FWP_UINT8"
	case FWP_UINT16:
		return "FWP_UINT16"
	case FWP_UINT32:
		return "FWP_UINT32"
	case FWP_UINT64:
		return "FWP_UINT64"
	case FWP_INT8:
		return "FWP_INT8"
	case FWP_INT16:
		return "FWP_INT16"
	case FWP_INT32:
		return "FWP_INT32"
	case FWP_INT64:
		return "FWP_INT64"
	case FWP_FLOAT:
		return "FWP_FLOAT"
	case FWP_DOUBLE:
		return "FWP_DOUBLE"
	case FWP_BYTE_ARRAY16_TYPE:
		return "FWP_BYTE_ARRAY16_TYPE"
	case FWP_BYTE_BLOB_TYPE:
		return "FWP_BYTE_BLOB_TYPE"
	case FWP_SID:
		return "FWP_SID"
	case FWP_SECURITY_DESCRIPTOR_TYPE:
		return "FWP_SECURITY_DESCRIPTOR_TYPE"
	case FWP_TOKEN_INFORMATION_TYPE:
		return "FWP_TOKEN_INFORMATION_TYPE"
	case FWP_TOKEN_ACCESS_INFORMATION_TYPE:
		return "FWP_TOKEN_ACCESS_INFORMATION_TYPE"
	case FWP_UNICODE_STRING_TYPE:
		return "FWP_UNICODE_STRING_TYPE"
	case FWP_BYTE_ARRAY6_TYPE:
		return "FWP_BYTE_ARRAY6_TYPE"
	case FWP_BITMAP_INDEX_TYPE:
		return "FWP_BITMAP_INDEX_TYPE"
	case FWP_BITMAP_ARRAY64_TYPE:
		return "FWP_BITMAP_ARRAY64_TYPE"
	case FWP_SINGLE_DATA_TYPE_MAX:
		return "FWP_SINGLE_DATA_TYPE_MAX"
	case FWP_V4_ADDR_MASK:
		return "FWP_V4_ADDR_MASK"
	case FWP_V6_ADDR_MASK:
		return "FWP_V6_ADDR_MASK"
	case FWP_RANGE_TYPE:
		return "FWP_RANGE_TYPE"
	case FWP_DATA_TYPE_MAX:
		return "FWP_DATA_TYPE_MAX"
	default:
		return fmt.Sprintf("FwpDataType_UNKNOWN(%d)", dt)
	}
}
