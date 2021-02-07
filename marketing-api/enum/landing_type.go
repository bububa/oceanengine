package enum

// 推广目的类型
type LandingType string

const (
	LINK    LandingType = "LINK"    // 销售线索收集
	APP     LandingType = "APP"     // 应用推广
	DPA     LandingType = "DPA"     // 商品目录推广
	GOODS   LandingType = "GOODS"   // 商品推广（鲁班）
	STORE   LandingType = "STORE"   // 门店推广
	AWEME   LandingType = "AWEME"   // 抖音号推广
	SHOP    LandingType = "SHOP"    // 电商店铺推广
	ARTICLE LandingType = "ARTICLE" // 头条文章推广，目前API暂不支持该推广目的的设定，可在平台侧选择该推广目的类型
)
