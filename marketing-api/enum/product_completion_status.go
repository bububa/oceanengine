package enum

// ProductCompletionStatus 字段填充状态
type ProductCompletionStatus string

const (
	// ProductCompletionStatus_AD_COMPLETED 广告场景已完善
	ProductCompletionStatus_AD_COMPLETED ProductCompletionStatus = "AD_COMPLETED"
	// ProductCompletionStatus_ALL_COMPLETED 必填字段已完善
	ProductCompletionStatus_ALL_COMPLETED ProductCompletionStatus = "ALL_COMPLETED"
	// ProductCompletionStatus_LEADS_COMPLETED 经营场景已完善
	ProductCompletionStatus_LEADS_COMPLETED ProductCompletionStatus = "LEADS_COMPLETED"
	// ProductCompletionStatus_TO_BE_COMPLETED 必填字段待完善
	ProductCompletionStatus_TO_BE_COMPLETED ProductCompletionStatus = "TO_BE_COMPLETED"
)
