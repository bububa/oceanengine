package enum

// AnchorType 锚点类型
type AnchorType string

const (
	// AnchorType_APP_GAME 应用下载-游戏
	AnchorType_APP_GAME AnchorType = "APP_GAME"
	// AnchorType_APP_INTERNET_SERVICE 应用下载-网服
	AnchorType_APP_INTERNET_SERVICE AnchorType = "APP_INTERNET_SERVICE"
	// AnchorType_APP_SHOP 应用下载-电商
	AnchorType_APP_SHOP AnchorType = "APP_SHOP"
	// AnchorType_ONLINE_SUBSCRIBE 高级在线预约
	AnchorType_ONLINE_SUBSCRIBE AnchorType = "ONLINE_SUBSCRIBE"
	// AnchorType_SHOPPING_CART 购物车
	AnchorType_SHOPPING_CART AnchorType = "SHOPPING_CART"
	// AnchorType_PRIVATE_CHAT  私信锚点
	AnchorType_PRIVATE_CHAT AnchorType = "PRIVATE_CHAT"
	// AnchorType_INSURANCE 保险锚点
	AnchorType_INSURANCE AnchorType = "INSURANCE"
)
