package enum

// AppPromotionType 子目标
type AppPromotionType string

const (
	// AppPromotionType_DOWNLOAD 应用下载
	AppPromotionType_DOWNLOAD AppPromotionType = "DOWNLOAD"
	// AppPromotionType_LAUNCH 应用调用
	AppPromotionType_LAUNCH AppPromotionType = "LAUNCH"
	// AppPromotionType_RESERVE 预约下载
	AppPromotionType_RESERVE AppPromotionType = "RESERVE"
)
