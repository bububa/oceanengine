package enum

// Behavior 抖音用户行为类型
type Behavior string

const (
	// BEHAVIOR_ALL全部
	BEHAVIOR_ALL Behavior = "ALL"
	// FOLLOWED_USER 已关注用户
	FOLLOWED_USER Behavior = "FOLLOWED_USER"
	// COMMENTED_USER 视频互动-已评论用户
	COMMENTED_USER Behavior = "COMMENTED_USER"
	// LIKED_USER 视频互动-已点赞用户
	LIKED_USER Behavior = "LIKED_USER"
	// SHARED_USER 视频互动-已分享用户
	SHARED_USER Behavior = "SHARED_USER"
	// GOODS_SHARE 下单过商品
	GOODS_SHARE Behavior = "GOODS_SHARE"
)
