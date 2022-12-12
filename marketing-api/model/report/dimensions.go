package report

import "github.com/bububa/oceanengine/marketing-api/enum"

// Dimensions 维度数据
type Dimensions struct {
	// StatDatetime 数据起始时间，分组条件包含 STAT_GROUP_BY_FIELD_STAT_TIME 时返回，格式为：yyyy-MM-dd HH:mm:ss
	StatDatetime string `json:"stat_datetime,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告组id，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 广告组名称，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	CampaignName string `json:"campaign_name,omitempty"`
	// AdID 计划id，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	AdID uint64 `json:"ad_id,omitempty"`
	// AdName 计划名称，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	AdName string `json:"ad_name,omitempty"`
	// CreativeID 创意id，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	CreativeID uint64 `json:"creative_id,omitempty"`
	// Bidword 关键词名称。如果分组条件中包括STAT_GROUP_BY_BIDWORD_ID,此字段有值
	Bidword string `json:"bidword,omitempty"`
	// BidwordID 关键词id。如果分组条件中包括STAT_GROUP_BY_BIDWORD_ID,此字段有值
	BidwordID uint64 `json:"bidword_id,omitempty"`
	// Query 搜索词。如果分组条件中包括STAT_GROUP_BY_QUERY,此字段有值
	Query string `json:"query,omitempty"`
	// Pricing 出价类型，分组条件包含STAT_GROUP_BY_PRICING时返回
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// ImageMode 素材类型，分组条件STAT_GROUP_BY_IMAGE_MODE返回
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// Inventory 投放广告位，分组条件包含STAT_GROUP_BY_INVENTORY时返回
	Inventory enum.StatInventoryType `json:"inventory,omitempty"`
	// CampaignType 广告组类型。如果分组条件中包括STAT_GROUP_BY_CAMPAIGN_TYPE,此字段有值
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// CreativeMaterialMode 创意类型，分组条件包含STAT_GROUP_BY_CREATIVE_MATERIAL_MODE时返回，允许值：STATIC_ASSEMBLE（程序化创意）、INTERVENE（自定义创意）
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// ExternalAction 转化类型。如果分组条件中包括STAT_GROUP_BY_EXTERNAL_ACTION,此字段有值
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// LandingType 推广目的类型，分组条件包含STAT_GROUP_BY_LANDING_TYPE时返回
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// PricingCategory 广告类型。如果分组条件中包括STAT_GROUP_BY_PRICING_CATEGORY,此字段有值
	PricingCategory string `json:"pricing_category,omitempty"`
	// ProvinceName 省份。如果分组条件中包括 STAT_GROUP_BY_PROVINCE_NAME 时返回
	ProvinceName string `json:"province_name,omitempty"`
	// CityName 城市。如果分组条件中包括 STAT_GROUP_BY_CITY_NAME 时返回。
	CityName string `json:"city_name,omitempty"`
	// Gender 性别。如果分组条件中包括 STAT_GROUP_BY_GENDER 时返回。
	Gender string `json:"gender,omitempty"`
	// Age 年龄。如果分组条件中包括 STAT_GROUP_BY_AGE
	Age string `json:"age,omitempty"`
	// Platform 平台。如果分组条件中包括 STAT_GROUP_BY_PLATFORM 时返回。
	Platform string `json:"platform,omitempty"`
	// Ac 网络类型。如果分组条件中包括 STAT_GROUP_BY_AC 时返回。
	Ac string `json:"ac,omitempty"`
	// AdTag 兴趣分类。如果分组条件中包括 STAT_GROUP_BY_AD_TAG 时返回。
	AdTag string `json:"ad_tag,omitempty"`
	// InterestTag 兴趣标签。如果分组条件中包括STAT_GROUP_BY_INTEREST_TAG,此字段有值
	InterestTag string `json:"interest_tag,omitempty"`
	// MaterialID 素材ID。如果分组条件中包括STAT_GROUP_BY_MATERIAL_ID,此字段有值
	MaterialID uint64 `json:"material_id,omitempty"`
	// PlayableID 试玩素材ID。如果分组条件中包括STAT_GROUP_BY_PLAYABLE_ID,此字段有值
	PlayableID uint64 `json:"playable_id,omitempty"`
	// PlayableName 试玩素材名称。如果分组条件中包括STAT_GROUP_BY_PLAYABLE_ID,此字段有值
	PlayableName string `json:"playable_name,omitempty"`
	// PlayableUrl 试玩素材链接。如果分组条件中包括STAT_GROUP_BY_PLAYABLE_ID,此字段有值
	PlayableUrl string `json:"playable_url,omitempty"`
	// PlayableOrientation 试玩素材展示方向。如果分组条件中包括STAT_GROUP_BY_PLAYABLE_ID,此字段有值
	PlayableOrientation enum.PlayableOrientation `json:"playable_orientation,omitempty"`
	// PlayablePreviewUrl 试玩素材预览链接。如果分组条件中包括STAT_GROUP_BY_PLAYABLE_ID,此字段有值
	PlayablePreviewUrl string `json:"playable_preview_url,omitempty"`
}
