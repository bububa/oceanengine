package enum

// RebateApprovalStatus 返点核算单审批状态
type RebateApprovalStatus string

const (
	// RebateApprovalStatus_AGENT_APPROVING 待代理商审批
	RebateApprovalStatus_AGENT_APPROVING RebateApprovalStatus = "AGENT_APPROVING"
	// RebateApprovalStatus_BUSINESS_APPROVING 待平台复核终审
	RebateApprovalStatus_BUSINESS_APPROVING RebateApprovalStatus = "BUSINESS_APPROVING"
	// RebateApprovalStatus_APPROVED 审批完成
	RebateApprovalStatus_APPROVED RebateApprovalStatus = "APPROVED"
)
