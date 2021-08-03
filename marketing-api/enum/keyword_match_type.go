package enum

// KeywordMatchType 关键词匹配类型
type KeywordMatchType string

const (
	// KeywordMatchType_PHRASE 短语匹配
	KeywordMatchType_PHRASE KeywordMatchType = "PHRASE"
	// KeywordMatchType_EXTENSIVE 广泛匹配
	KeywordMatchType_EXTENSIVE KeywordMatchType = "EXTENSIVE"
	// KeywordMatchType_PRECISION 精准匹配
	KeywordMatchType_PRECISION KeywordMatchType = "PRECISION"
)
