package audiencepackage

import "github.com/bububa/oceanengine/marketing-api/enum"

// AudiencePackage 定向包信息
type AudiencePackage struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AudiencePackageID 定向包id
	AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	// Name 定向包名称
	Name string `json:"name,omitempty"`
	// Description 定向包描述
	Description string `json:"description,omitempty"`
	// LandingType 定向包推广类型
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// AdType 广告类型，允许值：
	// ALL 所有广告（默认值）
	// SEARCH 搜索广告
	// 搜索定向包仅支持落地页、应用推广、抖音号、直播间，不支持商品、电商店铺、快应用、小游戏
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// HideIfConverted 过滤已转化用户; 注意点：过滤已转化不支持投放目标为展示、点击计划，不支持受众预估; 过滤已转化用户：选择过滤已转化用户之后，可以避免该广告再次投放给已转化过的用户，可以选择过滤5个广告层级的已转化用户（不限/广告计划/广告组/广告账户/公司账户/APP）
	HideIfConverted string `json:"hide_if_converted,omitempty"`
	// ConvertedTimeDuration 过滤时间范围，详见， 【附录-过滤时间范围】，当过滤已转化用户类型选择"公司账户"时填写，教育行业专属，非教育行业不会生效
	ConvertedTimeDuration string `json:"converted_time_duration,omitempty"`
	// DeliveryRange 广告投放范围【附录：广告投放范围】
	DeliverRange enum.AdDeliveryRange `json:"delivery_range,omitempty"`
	// HideIfExists 已安装用户，0表示不限，1表示过滤，2表示定向；过滤表示投放时不给安装客户展示广告，支持应用推广；定向表示投放时给安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；定向仅对Android链接生效。
	HideIfExists int `json:"hide_if_exists,omitempty"`
	// Audience
	Audience *Audience `json:"audience,omitempty"`
}
