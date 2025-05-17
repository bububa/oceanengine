package local

// FollowStateName 线索通话状态
type FollowStateName string

const (
	// FollowStateName_NOT_CALLED 待联系
	FollowStateName_NOT_CALLED FollowStateName = "NOT_CALLED"
	// FollowStateName_NOT_ANSWERED 未接通
	FollowStateName_NOT_ANSWERED FollowStateName = "NOT_ANSWERED"
	// FollowStateName_SHORT_ANSWERED 已接通
	FollowStateName_SHORT_ANSWERED FollowStateName = "SHORT_ANSWERED"
	// FollowStateName_ANSWERED 有效沟通
	FollowStateName_ANSWERED FollowStateName = "ANSWERED"
	// FollowStateName_DEEP_ANSWERED 深度沟通
	FollowStateName_DEEP_ANSWERED FollowStateName = "DEEP_ANSWERED"
)
