package qianchuan

// KeywordSuggestReason 推荐理由类型
type KeywordSuggestReason string

const (
	// KeywordSuggestReason_CLICK 高点击词
	KeywordSuggestReason_CLICK KeywordSuggestReason = "CLICK"
	// KeywordSuggestReason_DARKHORSE 黑马词
	KeywordSuggestReason_DARKHORSE KeywordSuggestReason = "DARKHORSE"
	// KeywordSuggestReason_LOW_COST 低成本词
	KeywordSuggestReason_LOW_COST KeywordSuggestReason = "LOW_COST"
	// KeywordSuggestReason_POTENTIAL 蓝海词
	KeywordSuggestReason_POTENTIAL KeywordSuggestReason = "POTENTIAL"
	// KeywordSuggestReason_PREFERENCE 同行买词
	KeywordSuggestReason_PREFERENCE KeywordSuggestReason = "PREFERENCE"
)
