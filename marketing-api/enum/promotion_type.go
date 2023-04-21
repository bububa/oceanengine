package enum

// PromotionType 投放内容，允许值：AWEME_HOME_PAGE：抖音主页（默认）LANDING_PAGE_LINK：落地页
type PromotionType string

const (
	// PromotionType_AWEME_HOME_PAGE 抖音主页
	PromotionType_AWEME_HOME_PAGE PromotionType = "AWEME_HOME_PAGE"
	// PromotionType_LANDING_PAGE_LINK 落地页
	PromotionType_LANDING_PAGE_LINK PromotionType = "LANDING_PAGE_LINK"
)
