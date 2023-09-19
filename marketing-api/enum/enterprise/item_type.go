package enterprise

// ItemType 视频类型
type ItemType string

const (
	// ITEM_AD 广告视频
	ITEM_AD ItemType = "ITEM_AD"
	// ITEM_CONTENT 非广告视频(抖音视频)
	ITEM_CONTENT ItemType = "ITEM_CONTENT"
)
