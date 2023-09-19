package enum

// SubscribeAdvStatus RDS订阅状态
type SubscribeAdvStatus string

const (
	// SubscribeAdvStatus_OK 推送中
	SubscribeAdvStatus_OK SubscribeAdvStatus = "OK"
	// SubscribeAdvStatus_PENDING 新增状态
	SubscribeAdvStatus_PENDING SubscribeAdvStatus = "PENDING"
	// SubscribeAdvStatus_UNAUTHORIZED 无权限
	SubscribeAdvStatus_UNAUTHORIZED SubscribeAdvStatus = "UNAUTHORIZED"
	// SubscribeAdvStatus_UNKNOWN 未知
	SubscribeAdvStatus_UNKNOWN SubscribeAdvStatus = "UNKNOWN"
)
