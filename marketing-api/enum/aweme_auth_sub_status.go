package enum

type AwemeAuthSubStatus string

const (
	// AwemeAuthSubStatus_INVALID_CANCEL 主动操作解除授权
	AwemeAuthSubStatus_INVALID_CANCEL AwemeAuthSubStatus = "INVALID_CANCEL"
	// AwemeAuthSubStatus_INVALID_EXPIRED 授权期限已到
	AwemeAuthSubStatus_INVALID_EXPIRED AwemeAuthSubStatus = "INVALID_EXPIRED"
	// AwemeAuthSubStatus_INVALID_REJECT C端拒绝授权
	AwemeAuthSubStatus_INVALID_REJECT AwemeAuthSubStatus = "INVALID_REJECT"
	// AwemeAuthSubStatus_INVALID_TIME_OUT 超时未确认
	AwemeAuthSubStatus_INVALID_TIME_OUT AwemeAuthSubStatus = "INVALID_TIME_OUT"
	// AwemeAuthSubStatus_RENEWING 续期待确认
	AwemeAuthSubStatus_RENEWING AwemeAuthSubStatus = "RENEWING"
	// AwemeAuthSubStatus_RENEW_FAIL 续期申请失效
	AwemeAuthSubStatus_RENEW_FAIL AwemeAuthSubStatus = "RENEW_FAIL"
	// AwemeAuthSubStatus_RENEW_SUCCESS 续期成功
	AwemeAuthSubStatus_RENEW_SUCCESS AwemeAuthSubStatus = "RENEW_SUCCESS"
)
