package enum

// CompensateEndReason 成本保障结束原因
type CompensateEndReason string

const (
	// CompensateEndReason_UN_OBERCOST 超成本比例没有达到1.2倍
	CompensateEndReason_UN_OBERCOST CompensateEndReason = "UN_OBERCOST"
	// CompensateEndReason_ROI_REAL_EXPECTED实际roi大于目标roi的80%
	CompensateEndReason_ROI_REAL_EXPECTED CompensateEndReason = "ROI_REAL_EXPECTED"
	// CompensateEndReason_CONVERT_UNDER_THRESHOLD转化数没有达到门槛
	CompensateEndReason_CONVERT_UNDER_THRESHOLD CompensateEndReason = "CONVERT_UNDER_THRESHOLD"
	// CompensateEndReason_CURRENCY_PRECISION赔付金额小于0.01元
	CompensateEndReason_CURRENCY_PRECISION CompensateEndReason = "CURRENCY_PRECISION"
)
