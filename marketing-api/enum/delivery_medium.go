package enum

// DeliveryMedium 单投放载体
type DeliveryMedium string

const (
	// DeliveryMedium_WECHAT_GAME：微信小游戏
	DeliveryMedium_WECHAT_GAME DeliveryMedium = "WECHAT_GAME"
	// DeliveryMedium_WECHAT_APP：微信小程序
	DeliveryMedium_WECHAT_APP DeliveryMedium = "WECHAT_APP"
	// DeliveryMedium_BYTE_GAME：字节小游戏
	DeliveryMedium_BYTE_GAME DeliveryMedium = "BYTE_GAME"
	// DeliveryMedium_BYTE_APP：字节小程序
	DeliveryMedium_BYTE_APP DeliveryMedium = "BYTE_APP"
	// DeliveryMedium_PRODUCT：商品
	DeliveryMedium_PRODUCT DeliveryMedium = "PRODUCT"
	// DeliveryMedium_ORANGE： 橙子落地页
	DeliveryMedium_ORANGE DeliveryMedium = "ORANGE"
	// DeliveryMedium_THIRDPARTY ：自研落地页
	DeliveryMedium_THIRDPARTY DeliveryMedium = "THIRDPARTY"
	// DeliveryMedium_ENTERPRISE ：企业号落地页
	DeliveryMedium_ENTERPRISE DeliveryMedium = "ENTERPRISE"
	// DeliveryMedium_AWEME： 抖音号
	DeliveryMedium_AWEME DeliveryMedium = "AWEME"
	// DeliveryMedium_QUICK_APP：快应用
	DeliveryMedium_QUICK_APP DeliveryMedium = "QUICK_APP"
	// DeliveryMedium_APP：应用
	DeliveryMedium_APP DeliveryMedium = "APP"
	// DeliveryMedium_LANDING_PAGE_LINK：落地页
	DeliveryMedium_LANDING_PAGE_LINK DeliveryMedium = "LANDING_PAGE_LINK"
)
