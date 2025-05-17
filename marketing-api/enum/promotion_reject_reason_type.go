package enum

// PromotionRejectReasonType 审核建议类型，
type PromotionRejectReasonType string

const (
	// PromotionRejectReasonType_NONE 无建议
	PromotionRejectReasonType_NONE PromotionRejectReasonType = "NONE"
	// PromotionRejectReasonType_REVIEW_REJECT 审核不通过
	PromotionRejectReasonType_REVIEW_REJECT PromotionRejectReasonType = "REVIEW_REJECT"
	// PromotionRejectReasonType_LOW_MATERAIL 低质素材
	PromotionRejectReasonType_LOW_MATERAIL PromotionRejectReasonType = "LOW_MATERIAL"
	// PromotionRejectReasonType_DISCOMFORT 引人不适
	PromotionRejectReasonType_DISCOMFORT PromotionRejectReasonType = "DISCOMFORT"
	// PromotionRejectReasonType_QUALITY_POOR 素材质量低
	PromotionRejectReasonType_QUALITY_POOR PromotionRejectReasonType = "QUALITY_POOR"
	// PromotionRejectReasonType_EXAGGERATION 夸大宣传
	PromotionRejectReasonType_EXAGGERATION PromotionRejectReasonType = "EXAGGERATION"
	// PromotionRejectReasonType_ELSE 其他
	PromotionRejectReasonType_ELSE PromotionRejectReasonType = "ELSE"
	// PromotionRejectReasonType_ALL 不限
	PromotionRejectReasonType_ALL PromotionRejectReasonType = "ALL"
)
