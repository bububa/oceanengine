package aweme

// OrderQuota 随心推使用中订单配额信息
type OrderQuota struct {
	// UsedQuota 使用中订单数
	UsedQuota int64 `json:"used_quota,omitempty"`
	// SumQuota 订单配额
	SumQuota int64 `json:"sum_quota,omitempty"`
	// MaxCost 最高月消耗
	MaxCost float64 `json:"max_cost,omitempty"`
	// QuotaGear 当前所在配额档位：1,2,3,4,5,6
	QuotaGear int `json:"quota_gear,omitempty"`
	// TerminateQuotaInfo 终止订单配额信息
	TerminateQuotaInfo *TerminateQuotaInfo `json:"terminate_quota_info,omitempty"`
}

// TerminateQuotaInfo 终止订单配额信息
type TerminateQuotaInfo struct {
	// TerminateQuotaUsed 已终止订单数
	TerminateQuotaUsed int64 `json:"terminate_quota_used,omitempty"`
	// TerminateQuotaSum 终止订单配额
	TerminateQuotaSum int64 `json:"terminate_quota_sum,omitempty"`
}
