package enum

// AwemeAuthStatus 授权状态
type AwemeAuthStatus string

const (
	// AwemeAuthStatus_AUTHRIZED 授权中
	AwemeAuthStatus_AUTHRIZED AwemeAuthStatus = "AUTHRIZED"
	// AwemeAuthStatus_AUTHRIZING 待授权确认
	AwemeAuthStatus_AUTHRIZING AwemeAuthStatus = "AUTHRIZING"
	// AwemeAuthStatus_INVALID 授权失效
	AwemeAuthStatus_INVALID AwemeAuthStatus = "INVALID"
)
