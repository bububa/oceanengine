package enum

// UniversalOrderStatus 星图任务订单状态
type UniversalOrderStatus string

const (
	// UniversalOrderStatus_ALL 不限
	UniversalOrderStatus_ALL UniversalOrderStatus = "ALL"
	// UniversalOrderStatus_WAIT_PAYMENT 待付款
	UniversalOrderStatus_WAIT_PAYMENT UniversalOrderStatus = "WAIT_PAYMENT"
	// UniversalOrderStatus_RECEIVEING 待接收
	UniversalOrderStatus_RECEIVEING UniversalOrderStatus = "RECEIVEING"
	// UniversalOrderStatus_ONGOING 进行中
	UniversalOrderStatus_ONGOING UniversalOrderStatus = "ONGOING"
	// UniversalOrderStatus_FINISHED 已完成
	UniversalOrderStatus_FINISHED UniversalOrderStatus = "FINISHED"
	// UniversalOrderStatus_WAIT_EVALUATE 待评价
	UniversalOrderStatus_WAIT_EVALUATE UniversalOrderStatus = "WAIT_EVALUATE"
	// UniversalOrderStatus_CANCELED 已取消
	UniversalOrderStatus_CANCELED UniversalOrderStatus = "CANCELED"
)
