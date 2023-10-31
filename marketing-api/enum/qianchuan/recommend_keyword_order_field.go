package qianchuan

// RecommendKeywordOrderField 排序字段
type RecommendKeywordOrderField string

const (
	// RecommendKeywordOrderField_DEFAULT 推荐度（默认）
	RecommendKeywordOrderField_DEFAULT RecommendKeywordOrderField = "DEFAULT"
	// RecommendKeywordOrderField_COMPETITION 竞争程度
	RecommendKeywordOrderField_COMPETITION RecommendKeywordOrderField = "COMPETITION"
	// RecommendKeywordOrderField_PV 月搜索量
	RecommendKeywordOrderField_PV RecommendKeywordOrderField = "PV"
	// RecommendKeywordOrderField_RELEVANCE 相关性
	RecommendKeywordOrderField_RELEVANCE RecommendKeywordOrderField = "RELEVANCE"
)
