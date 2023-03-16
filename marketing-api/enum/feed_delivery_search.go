package enum

// FeedDeliverySearch 搜索快投关键词功能
type FeedDeliverySearch string

const (
	// FeedDeliverySearch_HAS_OPEN 启用
	FeedDeliverySearch_HAS_OPEN FeedDeliverySearch = "HAS_OPEN"
	// FeedDeliverySearch_DISABLED 未启用
	FeedDeliverySearch_DISABLED FeedDeliverySearch = "DISABLED"
)
