package enum

// StarAdUniteTaskStatus 任务状态
type StarAdUniteTaskStatus string

const (
	// StarAdUniteTaskStatus_ALL 不限
	StarAdUniteTaskStatus_ALL StarAdUniteTaskStatus = "ALL"
	// StarAdUniteTaskStatus_BILLING 计费中
	StarAdUniteTaskStatus_BILLING StarAdUniteTaskStatus = "BILLING"
	// StarAdUniteTaskStatus_BULLETIN 公示中
	StarAdUniteTaskStatus_BULLETIN StarAdUniteTaskStatus = "BULLETIN"
	// StarAdUniteTaskStatus_CANCELED 已取消
	StarAdUniteTaskStatus_CANCELED StarAdUniteTaskStatus = "CANCELED"
	// StarAdUniteTaskStatus_CLOSED 已关闭
	StarAdUniteTaskStatus_CLOSED StarAdUniteTaskStatus = "CLOSED"
	// StarAdUniteTaskStatus_ENROLL 预热中
	StarAdUniteTaskStatus_ENROLL StarAdUniteTaskStatus = "ENROLL"
	// StarAdUniteTaskStatus_FINISHED 已完成
	StarAdUniteTaskStatus_FINISHED StarAdUniteTaskStatus = "FINISHED"
	// StarAdUniteTaskStatus_ONGOING 进行中
	StarAdUniteTaskStatus_ONGOING StarAdUniteTaskStatus = "ONGOING"
	// StarAdUniteTaskStatus_PROVIDER_ACCEPTING 服务商接单中
	StarAdUniteTaskStatus_PROVIDER_ACCEPTING StarAdUniteTaskStatus = "PROVIDER_ACCEPTING"
	// StarAdUniteTaskStatus_STARTED 投稿中
	StarAdUniteTaskStatus_STARTED StarAdUniteTaskStatus = "STARTED"
	// StarAdUniteTaskStatus_RECEIVEING 待接收
	StarAdUniteTaskStatus_RECEIVEING StarAdUniteTaskStatus = "RECEIVEING"
	// StarAdUniteTaskStatus_WAIT_EVALUATE 待评价
	StarAdUniteTaskStatus_WAIT_EVALUATE StarAdUniteTaskStatus = "WAIT_EVALUATE"
	// StarAdUniteTaskStatus_WAIT_PAYMENT 待付款
	StarAdUniteTaskStatus_WAIT_PAYMENT StarAdUniteTaskStatus = "WAIT_PAYMENT"
)
