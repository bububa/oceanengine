package enum

// DmpPushStatus 人群包推送状态
type DmpPushStatus int

const (
	// DmpPushStatus_NOT_START 未推送
	DmpPushStatus_NOT_START DmpPushStatus = 0
	// DmpPushStatus_PUSHING 推送中
	DmpPushStatus_PUSHING DmpPushStatus = 1
	// DmpPushStatus_PUSHED 已推送
	DmpPushStatus_PUSHED DmpPushStatus = 2
	// DmpPushStatus_FAILED 推送失败
	DmpPushStatus_FAILED DmpPushStatus = 3
)
