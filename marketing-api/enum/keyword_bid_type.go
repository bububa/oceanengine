package enum

// KeywordBidType 关键词出价类型
type KeywordBidType string

const (
	// KeywordBidType_CUSTOM 自定义
	KeywordBidType_CUSTOM KeywordBidType = "CUSTOM"
	// KeywordBidType_WITH_AD 随计划出价
	KeywordBidType_WITH_AD KeywordBidType = "WITH_AD"
	// KeywordBidType_SUGGEST 推荐出价
	KeywordBidType_SUGGEST KeywordBidType = "SUGGEST"
	// KeywordBidType_FEED_TO_SEARCH 搜索快投
	KeywordBidType_FEED_TO_SEARCH KeywordBidType = "FEED_TO_SEARCH"
	// KeywordBidType_BRAND_AD 品牌广告专用
	KeywordBidType_BRAND_AD KeywordBidType = "BRAND_AD"
)
