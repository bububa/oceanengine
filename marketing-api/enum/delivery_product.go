package enum

// DeliveryProduct 推广产品
type DeliveryProduct string

const (
	// DeliveryProduct_NONE：无产品
	DeliveryProduct_NONE DeliveryProduct = "NONE"
	// DeliveryProduct_APP ：应用
	DeliveryProduct_APP DeliveryProduct = "APP"
	// DeliveryProduct_PRODUCT：商品
	DeliveryProduct_PRODUCT DeliveryProduct = "PRODUCT"
	// DeliveryProduct_WECHAT_GAME：微信小游戏
	DeliveryProduct_WECHAT_GAME DeliveryProduct = "WECHAT_GAM"
	// DeliveryProduct_WECHAT_APP：微信小程序
	DeliveryProduct_WECHAT_APP DeliveryProduct = "WECHAT_APP"
	// DeliveryProduct_BYTE_GAME：字节小游戏
	DeliveryProduct_BYTE_GAME DeliveryProduct = "BYTE_GAME"
	// DeliveryProduct_BYTE_APP：字节小程序
	DeliveryProduct_BYTE_APP DeliveryProduct = "BYTE_APP"
	// DeliveryProduct_QUICK_APP：快应用
	DeliveryProduct_QUICK_APP DeliveryProduct = "QUICK_APP"
	// DeliveryProduct_AWEME：抖音号
	DeliveryProduct_AWEME DeliveryProduct = "AWEME"
)
