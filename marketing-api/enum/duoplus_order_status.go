package enum

// DuoplusOrderStatus DOU+订单状态
type DuoplusOrderStatus string

const (
	// DuoplusOrderStatus_UNPAID 未支付
	DuoplusOrderStatus_UNPAID DuoplusOrderStatus = "UNPAID"
	// DuoplusOrderStatus_AUDITING 审核中
	DuoplusOrderStatus_AUDITING DuoplusOrderStatus = "AUDITING"
	// DuoplusOrderStatus_OFFLINE_AUDIT 审核不通过
	DuoplusOrderStatus_OFFLINE_AUDIT DuoplusOrderStatus = "OFFLINE_AUDIT"
	// DuoplusOrderStatus_TIME_NO_REACH 待开播
	DuoplusOrderStatus_TIME_NO_REACH DuoplusOrderStatus = "TIME_NO_REACH"
	// DuoplusOrderStatus_DELIVERING 投放中
	DuoplusOrderStatus_DELIVERING DuoplusOrderStatus = "DELIVERING"
	// DuoplusOrderStatus_UNDELIVERIED 投放终止
	DuoplusOrderStatus_UNDELIVERIED DuoplusOrderStatus = "UNDELIVERIED"
	// DuoplusOrderStatus_DELIVERIED 投放完成/结束
	DuoplusOrderStatus_DELIVERIED DuoplusOrderStatus = "DELIVERIED"
	// DuoplusOrderStatus_AUDIT_PAUSE 审核暂停
	DuoplusOrderStatus_AUDIT_PAUSE DuoplusOrderStatus = "AUDIT_PAUSE"
	// DuopluseOrderStatus_WAIT_TO_START 等待开始
	DuopluseOrderStatus_WAIT_TO_START DuoplusOrderStatus = "WAIT_TO_START"
)
