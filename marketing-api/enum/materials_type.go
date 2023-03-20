package enum

// MaterialsType  素材类型，直播场景必填，
type MaterialsType string

const (
	// LIVE_MATERIALS 直播素材
	LIVE_MATERIALS MaterialsType = "LIVE_MATERIALS"
	// PROMOTION_MATERIALS 广告素材
	PROMOTION_MATERIALS MaterialsType = "PROMOTION_MATERIALS"
)
