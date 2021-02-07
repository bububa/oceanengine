package report

import "github.com/bububa/oceanengine/marketing-api/enum"

type Dimensions struct {
	StatDatetime         string                   `json:"stat_datetime,omitempty"`          // 数据起始时间，分组条件包含 STAT_GROUP_BY_FIELD_STAT_TIME 时返回，格式为：yyyy-MM-dd HH:mm:ss
	AdvertiserID         uint64                   `json:"advertiser_id,omitempty"`          // 广告主ID
	CampaignID           uint64                   `json:"campaign_id,omitempty"`            // 广告组id，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	CampaignName         string                   `json:"campaign_name,omitempty"`          // 广告组名称，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	AdID                 uint64                   `json:"ad_id,omitempty"`                  // 计划id，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	AdName               string                   `json:"ad_name,omitempty"`                // 计划名称，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	CreativeID           uint64                   `json:"creative_id,omitempty"`            // 创意id，分组条件包含STAT_GROUP_BY_FIELD_ID时，返回具体值
	Bidword              string                   `json:"bidword,omitempty"`                // 关键词名称。如果分组条件中包括STAT_GROUP_BY_BIDWORD_ID,此字段有值
	BidwordID            uint64                   `json:"bidword_id,omitempty"`             // 关键词id。如果分组条件中包括STAT_GROUP_BY_BIDWORD_ID,此字段有值
	Query                string                   `json:"query,omitempty"`                  // 搜索词。如果分组条件中包括STAT_GROUP_BY_QUERY,此字段有值
	Pricing              enum.PricingType         `json:"pricing,omitempty"`                // 出价类型，分组条件包含STAT_GROUP_BY_PRICING时返回
	ImageMode            enum.ImageMode           `json:"image_mode,omitempty"`             // 素材类型，分组条件STAT_GROUP_BY_IMAGE_MODE返回
	Inventory            enum.StatInventoryType   `json:"inventory,omitempty"`              // 投放广告位，分组条件包含STAT_GROUP_BY_INVENTORY时返回
	CampaignType         string                   `json:"campaign_type,omitempty"`          // 广告组类型。如果分组条件中包括STAT_GROUP_BY_CAMPAIGN_TYPE,此字段有值
	CreativeMaterialMode string                   `json:"creative_material_mode,omitempty"` // 创意类型，分组条件包含STAT_GROUP_BY_CREATIVE_MATERIAL_MODE时返回，允许值：STATIC_ASSEMBLE（程序化创意）、INTERVENE（自定义创意）
	ExternalAction       string                   `json:"external_action,omitempty"`        // 转化类型。如果分组条件中包括STAT_GROUP_BY_EXTERNAL_ACTION,此字段有值
	LandingType          enum.LandingType         `json:"landing_type,omitempty"`           // 推广目的类型，分组条件包含STAT_GROUP_BY_LANDING_TYPE时返回
	PricingCategory      string                   `json:"pricing_category,omitempty"`       // 广告类型。如果分组条件中包括STAT_GROUP_BY_PRICING_CATEGORY,此字段有值
	ProvinceName         string                   `json:"province_name,omitempty"`          // 省份。如果分组条件中包括 STAT_GROUP_BY_PROVINCE_NAME 时返回
	CityName             string                   `json:"city_name,omitempty"`              // 城市。如果分组条件中包括 STAT_GROUP_BY_CITY_NAME 时返回。
	Gender               string                   `json:"gender,omitempty"`                 // 性别。如果分组条件中包括 STAT_GROUP_BY_GENDER 时返回。
	Age                  string                   `json:"age,omitempty"`                    // 年龄。如果分组条件中包括 STAT_GROUP_BY_AGE
	Platform             string                   `json:"platform,omitempty"`               // 平台。如果分组条件中包括 STAT_GROUP_BY_PLATFORM 时返回。
	Ac                   string                   `json:"ac,omitempty"`                     // 网络类型。如果分组条件中包括 STAT_GROUP_BY_AC 时返回。
	AdTag                string                   `json:"ad_tag,omitempty"`                 // 兴趣分类。如果分组条件中包括 STAT_GROUP_BY_AD_TAG 时返回。
	InterestTag          string                   `json:"interest_tag,omitempty"`           // 兴趣标签。如果分组条件中包括STAT_GROUP_BY_INTEREST_TAG,此字段有值
	MaterialID           uint64                   `json:"material_id,omitempty"`            // 素材ID。如果分组条件中包括STAT_GROUP_BY_MATERIAL_ID,此字段有值
	PlayableID           uint64                   `json:"playable_id,omitempty"`            // 试玩素材ID。如果分组条件中包括STAT_GROUP_BY_PLAYABLE_ID,此字段有值
	PlayableName         string                   `json:"playable_name,omitempty"`          // 试玩素材名称。如果分组条件中包括STAT_GROUP_BY_PLAYABLE_ID,此字段有值
	PlayableUrl          string                   `json:"playable_url,omitempty"`           // 试玩素材链接。如果分组条件中包括STAT_GROUP_BY_PLAYABLE_ID,此字段有值
	PlayableOrientation  enum.PlayableOrientation `json:"playable_orientation,omitempty"`   // 试玩素材展示方向。如果分组条件中包括STAT_GROUP_BY_PLAYABLE_ID,此字段有值
	PlayablePreviewUrl   string                   `json:"playable_preview_url,omitempty"`   // 试玩素材预览链接。如果分组条件中包括STAT_GROUP_BY_PLAYABLE_ID,此字段有值
}
