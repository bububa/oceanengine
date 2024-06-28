package enum

// CompensateInvalidReason 成本保障失效原因
type CompensateInvalidReason string

const (
	// CompensateInvalidReason_AUD_CHANGES 单日修改定向超过2次
	CompensateInvalidReason_AUD_CHANGES CompensateInvalidReason = "AUD_CHANGES"
	// CompensateInvalidReason_BID_CHANGES 单日修改出价超过2次
	CompensateInvalidReason_BID_CHANGES CompensateInvalidReason = "BID_CHANGES"
	// CompensateInvalidReason_ROI_CHANGES 单日修改roi_goal超过2次
	CompensateInvalidReason_ROI_CHANGES CompensateInvalidReason = "ROI_CHANGES"
	// CompensateInvalidReason_ANTI_SPAM 命中反作弊
	CompensateInvalidReason_ANTI_SPAM CompensateInvalidReason = "ANTI_SPAM"
	// CompensateInvalidReason_BID_TYPE_EXPIRED 选择的出价产品暂不支持成本保障
	CompensateInvalidReason_BID_TYPE_EXPIRED CompensateInvalidReason = "BID_TYPE_EXPIRED"
	// CompensateInvalidReason_MANUAL_JUDGE_SPAM 有异常的作弊行为
	CompensateInvalidReason_MANUAL_JUDGE_SPAM CompensateInvalidReason = "MANUAL_JUDGE_SPAM"
	// CompensateInvalidReason_AUD_BID_CHANGES 单日修改定向/出价超过2次
	CompensateInvalidReason_AUD_BID_CHANGES CompensateInvalidReason = "AUD_BID_CHANGES"
	// CompensateInvalidReason_AUD_ROI_CHANGES 单日修改定向/ROI目标超过2次
	CompensateInvalidReason_AUD_ROI_CHANGES CompensateInvalidReason = "AUD_ROI_CHANGES"
	// CompensateInvalidReason_ACCOUNT_TRANSFER_APPLICATION 申请转户
	CompensateInvalidReason_ACCOUNT_TRANSFER_APPLICATION CompensateInvalidReason = "ACCOUNT_TRANSFER_APPLICATION"
)
