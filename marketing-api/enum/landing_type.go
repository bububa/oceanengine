package enum

// LandingType 推广目的类型
type LandingType string

const (
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
)
