package qianchuan

// OrderStatus 订单状态
type OrderStatus string

const (
	// OrderStatus_AUDIT 审核中
	OrderStatus_AUDIT OrderStatus = "AUDIT"
	// OrderStatus_BOOK 未开播
	OrderStatus_BOOK OrderStatus = "BOOK"
	// OrderStatus_CREATING 已支付，计划创建中
	OrderStatus_CREATING OrderStatus = "CREATING"
	// OrderStatus_DELIVERY_OK 投放中
	OrderStatus_DELIVERY_OK OrderStatus = "DELIVERY_OK"
	// OrderStatus_FAILED 计划创建失败
	OrderStatus_FAILED OrderStatus = "FAILED"
	// OrderStatus_FINISHED 投放完成
	OrderStatus_FINISHED OrderStatus = "FINISHED"
	// OrderStatus_FROZEN 投放终止
	OrderStatus_FROZEN OrderStatus = "FROZEN"
	// OrderStatus_OFFLINE_AUDIT 审核不通过
	OrderStatus_OFFLINE_AUDIT OrderStatus = "OFFLINE_AUDIT"
	// OrderStatus_OVER 投放结束
	OrderStatus_OVER OrderStatus = "OVER"
	// OrderStatus_TIMEOUT 未支付超过15分钟，订单关闭
	OrderStatus_TIMEOUT OrderStatus = "TIMEOUT"
	// OrderStatus_UNPAID 未支付
	OrderStatus_UNPAID OrderStatus = "UNPAID"
	// OrderStatus_UNPAIDCANCEL 未支付取消订单
	OrderStatus_UNPAIDCANCEL OrderStatus = "UNPAIDCANCEL"
)
