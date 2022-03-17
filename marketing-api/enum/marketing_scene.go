package enum

// MarketingScene 游戏预约场景
type MarketingScene string

const (
	// MarketingScene_GAME_PROMOTION 游戏大推
	MarketingScene_GAME_PROMOTION MarketingScene = "GAME_PROMOTION"
	// MarketingScene_GAME_SUBSCRIBE 游戏预约
	MarketingScene_GAME_SUBSCRIBE MarketingScene = "GAME_SUBSCRIBE"
	// MarketingScene_NORMAL 普通场景
	MarketingScene_NORMAL MarketingScene = "NORMAL"
)
