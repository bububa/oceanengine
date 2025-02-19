package enum

// AwemeAuthSubStatus 授权子状态
type AwemeAuthSubStatus string

const (
	// AwemeAuthSubStatus_RENEWING 续期待确认
	AwemeAuthSubStatus_RENEWING AwemeAuthSubStatus = "RENEWING"
	// AwemeAuthSubStatus_RENEW_FAIL 续期申请失效
	AwemeAuthSubStatus_RENEW_FAIL AwemeAuthSubStatus = "RENEW_FAIL"
	// AwemeAuthSubStatus_RENEW_SUCCESS 续期成功
	AwemeAuthSubStatus_RENEW_SUCCESS AwemeAuthSubStatus = "RENEW_SUCCESS"
	// AwemeAuthSubStatus_RENEW_FAILED_BY_AWEME 授权续期申请失败
	AwemeAuthSubStatus_RENEW_FAILED_BY_AWEME AwemeAuthSubStatus = "RENEW_FAILED_BY_AWEME"
	// AwemeAuthSubStatus_INVALID_CANCEL 主动操作解除授权
	AwemeAuthSubStatus_INVALID_CANCEL AwemeAuthSubStatus = "INVALID_CANCEL"
	// AwemeAuthSubStatus_INVALID_EXPIRED 授权期限已到
	AwemeAuthSubStatus_INVALID_EXPIRED AwemeAuthSubStatus = "INVALID_EXPIRED"
	// AwemeAuthSubStatus_INVALID_REJECT C端拒绝授权
	AwemeAuthSubStatus_INVALID_REJECT AwemeAuthSubStatus = "INVALID_REJECT"
	// AwemeAuthSubStatus_INVALID_TIME_OUT 超时未确认
	AwemeAuthSubStatus_INVALID_TIME_OUT AwemeAuthSubStatus = "INVALID_TIME_OUT"
	// AwemeAuthSubStatus_INVALID_FAILED_BY_AWEME 授权申请失败
	AwemeAuthSubStatus_INVALID_FAILED_BY_AWEME AwemeAuthSubStatus = "INVALID_FAILED_BY_AWEME"
	// AwemeAuthSubStatus_INVALID_PROCESS_TIME_OUT 超时未处理自动解除
	AwemeAuthSubStatus_INVALID_PROCESS_TIME_OUT AwemeAuthSubStatus = "INVALID_PROCESS_TIME_OUT"
	// AwemeAuthSubStatus_AWEME_REVOKE_REQUEST 创作者发起解除申请
	AwemeAuthSubStatus_AWEME_REVOKE_REQUEST AwemeAuthSubStatus = "AWEME_REVOKE_REQUEST"
	// AwemeAuthSubStatus_CONFIRM_REVOKE_REQUEST 同意解除授权申请
	AwemeAuthSubStatus_CONFIRM_REVOKE_REQUEST AwemeAuthSubStatus = "CONFIRM_REVOKE_REQUEST"
	// AwemeAuthSubStatus_UNQUALIFIED_AUTO_RELEASE 不达要求自动解除
	AwemeAuthSubStatus_UNQUALIFIED_AUTO_RELEASE AwemeAuthSubStatus = "UNQUALIFIED_AUTO_RELEASE"
	// AwemeAuthSubStatus_ENTERPRISE_AUTH_RELEASE 身份变更，不达要求
	AwemeAuthSubStatus_ENTERPRISE_AUTH_RELEASE AwemeAuthSubStatus = "ENTERPRISE_AUTH_RELEASE"
)
