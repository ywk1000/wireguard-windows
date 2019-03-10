package winwfp

import "fmt"

// FWP_MATCH_TYPE defined in fwptypes.h
// (https://docs.microsoft.com/en-us/windows/desktop/api/fwptypes/ne-fwptypes-fwp_match_type_)
type FwpMatchType uint32

const (
	FWP_MATCH_EQUAL                  FwpMatchType = 0
	FWP_MATCH_GREATER                FwpMatchType = FWP_MATCH_EQUAL + 1
	FWP_MATCH_LESS                   FwpMatchType = FWP_MATCH_GREATER + 1
	FWP_MATCH_GREATER_OR_EQUAL       FwpMatchType = FWP_MATCH_LESS + 1
	FWP_MATCH_LESS_OR_EQUAL          FwpMatchType = FWP_MATCH_GREATER_OR_EQUAL + 1
	FWP_MATCH_RANGE                  FwpMatchType = FWP_MATCH_LESS_OR_EQUAL + 1
	FWP_MATCH_FLAGS_ALL_SET          FwpMatchType = FWP_MATCH_RANGE + 1
	FWP_MATCH_FLAGS_ANY_SET          FwpMatchType = FWP_MATCH_FLAGS_ALL_SET + 1
	FWP_MATCH_FLAGS_NONE_SET         FwpMatchType = FWP_MATCH_FLAGS_ANY_SET + 1
	FWP_MATCH_EQUAL_CASE_INSENSITIVE FwpMatchType = FWP_MATCH_FLAGS_NONE_SET + 1
	FWP_MATCH_NOT_EQUAL              FwpMatchType = FWP_MATCH_EQUAL_CASE_INSENSITIVE + 1
	FWP_MATCH_PREFIX                 FwpMatchType = FWP_MATCH_NOT_EQUAL + 1
	FWP_MATCH_NOT_PREFIX             FwpMatchType = FWP_MATCH_PREFIX + 1
	FWP_MATCH_TYPE_MAX               FwpMatchType = FWP_MATCH_NOT_PREFIX + 1
)

func (m FwpMatchType) String() string {
	switch m {
	case FWP_MATCH_EQUAL:
		return "FWP_MATCH_EQUAL"
	case FWP_MATCH_GREATER:
		return "FWP_MATCH_GREATER"
	case FWP_MATCH_LESS:
		return "FWP_MATCH_LESS"
	case FWP_MATCH_GREATER_OR_EQUAL:
		return "FWP_MATCH_GREATER_OR_EQUAL"
	case FWP_MATCH_LESS_OR_EQUAL:
		return "FWP_MATCH_LESS_OR_EQUAL"
	case FWP_MATCH_RANGE:
		return "FWP_MATCH_RANGE"
	case FWP_MATCH_FLAGS_ALL_SET:
		return "FWP_MATCH_FLAGS_ALL_SET"
	case FWP_MATCH_FLAGS_ANY_SET:
		return "FWP_MATCH_FLAGS_ANY_SET"
	case FWP_MATCH_FLAGS_NONE_SET:
		return "FWP_MATCH_FLAGS_NONE_SET"
	case FWP_MATCH_EQUAL_CASE_INSENSITIVE:
		return "FWP_MATCH_EQUAL_CASE_INSENSITIVE"
	case FWP_MATCH_NOT_EQUAL:
		return "FWP_MATCH_NOT_EQUAL"
	case FWP_MATCH_PREFIX:
		return "FWP_MATCH_PREFIX"
	case FWP_MATCH_NOT_PREFIX:
		return "FWP_MATCH_NOT_PREFIX"
	case FWP_MATCH_TYPE_MAX:
		return "FWP_MATCH_TYPE_MAX"
	default:
		return fmt.Sprintf("FwpMatchType_UNKNOWN(%d)", m)
	}
}
