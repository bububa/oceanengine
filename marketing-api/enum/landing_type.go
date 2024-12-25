package enum

// LandingType 推广目的类型
type LandingType string

const (
	// EXTERNAL 落地页
	EXTENAL LandingType = "EXTERNAL"
	// LINK 销售线索收集
	LINK LandingType = "LINK"
	// APP 应用推广
	APP LandingType = "APP"
	// DPA 商品目录推广
	DPA LandingType = "DPA"
	// GOODS 商品推广（鲁班）
	GOODS LandingType = "GOODS"
	// STORE 门店推广
	STORE LandingType = "STORE"
	// AWEME 抖音号推广
	AWEME LandingType = "AWEME"
	// SHOP 电商店铺推广
	SHOP LandingType = "SHOP"
	// ARTICLE 头条文章推广，目前API暂不支持该推广目的的设定，可在平台侧选择该推广目的类型
	ARTICLE LandingType = "ARTICLE"
	// MICRO_GAME 小游戏
	MICRO_GAME LandingType = "MICRO_GAME"
	// QUICK_APP 快应用
	QUICK_APP LandingType = "QUICK_APP"
	// NATIVE_ACTION 原生互动
	NATIVE_ACTION LandingType = "NATIVE_ACTION"
	// APP_ANDROID 应用下载-安卓
	APP_ANDROID LandingType = "APP_ANDROID"
	// APP_IOS 应用下载-IOS
	APP_IOS LandingType = "APP_IOS"
	// LIVE 直播间推广
	LIVE LandingType = "LIVE"
)
