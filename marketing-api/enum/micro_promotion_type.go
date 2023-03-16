package enum

// MicroPromotionType 小程序类型
type MicroPromotionType string

const (
	// MicroPromotionType_WECHAT_GAME 微信小游戏
	MicroPromotionType_WECHAT_GAME MicroPromotionType = "WECHAT_GAME"
	// MicroPromotionType_WECHAT_APP 微信小程序
	MicroPromotionType_WECHAT_APP MicroPromotionType = "WECHAT_APP"
	// MicroPromotionType_BYTE_GAME 字节小游戏
	MicroPromotionType_BYTE_GAME MicroPromotionType = "BYTE_GAME"
	// MicroPromotionType_BYTE_APP 字节小程序
	MicroPromotionType_BYTE_APP MicroPromotionType = "BYTE_APP"
)
