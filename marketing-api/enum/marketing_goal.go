package enum

// MarketingGoal 营销目标
type MarketingGoal string

const (
	// MarketingGoal_VIDEO_PROM_GOODS 短视频带货
	MarketingGoal_VIDEO_PROM_GOODS MarketingGoal = "VIDEO_PROM_GOODS"
	// MarketingGoal_LIVE_PROM_GOODS 直播带货
	MarketingGoal_LIVE_PROM_GOODS MarketingGoal = "LIVE_PROM_GOODS"
	// MarketingGoal_VIDEO_AND_IMAGE 短视频/图片
	MarketingGoal_VIDEO_AND_IMAGE MarketingGoal = "VIDEO_AND_IMAGE"
	// MarketingGoal_LIVE 直播
	MarketingGoal_LIVE MarketingGoal = "LIVE"
	// MarketingGoal_VIDEO 短视频
	MarketingGoal_VIDEO MarketingGoal = "VIDEO"
	// MarketingGoal_ALL 不限
	MarketingGoal_ALL MarketingGoal = "ALL"
)
