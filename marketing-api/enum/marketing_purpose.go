package enum

// MarketingPurpose 营销目的
type MarketingPurpose string

const (
	// MarketingPurpose_UNLIMITED 不限
	MarketingPurpose_UNLIMITED = "UNLIMITED"
	// MarketingPurpose_CONVERSION 行动转化
	MarketingPurpose_CONVERSION MarketingPurpose = "CONVERSION"
	// MarketingPurpose_INTENTION 用户意向
	MarketingPurpose_INTENTION MarketingPurpose = "INTENTION"
	// MarketingPurpose_ACKNOWLEDGE 品牌认知
	MarketingPurpose_ACKNOWLEDGE MarketingPurpose = "ACKNOWLEDGE"
)
