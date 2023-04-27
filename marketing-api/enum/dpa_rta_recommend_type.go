package enum

// DpaRtaRecommedType RTA推荐逻辑，ONLY仅RTA推荐商品，MORE基于RTA推荐更多商品，开启RTA重定向开关时必填
type DpaRtaRecommendType string

const (
	// DpaRtaRecommendType_ONLY 仅RTA推荐商品
	DpaRtaRecommendType_ONLY DpaRtaRecommendType = "ONLY"
	// DpaRtaRecommendType_MORE 基于RTA推荐更多商品
	DpaRtaRecommendType_MORE DpaRtaRecommendType = "MORE"
)
