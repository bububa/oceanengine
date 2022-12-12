package ad

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/creative"
)

// Ad 计划详情
type Ad struct {
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// CampaignID 广告组ID（若为托管计划，则返回null）
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// MarketingGoal 营销目标
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// PromotionWay 推广方式
	PromotionWay enum.PromotionWay `json:"promotion_way,omitempty"`
	// Name 计划名称
	Name string `json:"name,omitempty"`
	// Status 计划投放状态
	Status qianchuan.AdStatus `json:"status,omitempty"`
	// OptStatus 计划操作状态
	OptStatus qianchuan.AdOptStatus `json:"opt_status,omitempty"`
	// AdCreateTime 计划创建时间
	AdCreateTime string `json:"ad_create_time,omitempty"`
	// AdModifyTime 计划修改时间
	AdModityTime string `json:"ad_modity_time,omitempty"`
	// AwemeInfo 计划中关联的抖音号信息
	AwemeInfo []AwemeInfo `json:"aweme_info,omitempty"`
	// ProductInfo 商品列表
	ProductInfo []ProductInfo `json:"product_info,omitempty"`
	// RoomInfo 直播间列表
	RoomInfo []RoomInfo `json:"room_info,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// Audience 定向设置
	Audience *Audience `json:"audience,omitempty"`
	// CreativeMaterialMode 创意呈现方式
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// FirstIndustryID 创意一级行业ID
	FirstIndustryID uint64 `json:"first_industry_id,omitempty"`
	// SecondIndustryID 创意二级行业ID
	SecondIndustryID uint64 `json:"second_industry_id,omitempty"`
	// ThirdIndustryID 创意三级行业ID
	ThirdIndustryID uint64 `json:"third_industry_id,omitempty"`
	// AdKeywords 创意标签
	AdKeywords []string `json:"ad_keywords,omitempty"`
	// CreativeList 创意信息（若为托管计划，则返回空数组）
	CreativeList []creative.Creative `json:"creative_list,omitempty"`
	// ProgrammaticCreativeMediaList 程序化创意素材信息
	ProgrammaticCreativeMediaList []creative.ProgrammaticCreativeMedia `json:"programmatic_creative_media_list,omitempty"`
	// ProgrammaticCreativeTitleList 程序化创意标题信息
	ProgrammaticCreativeTitleList []creative.TitleMaterial `json:"programmatic_creative_title_list,omitempty"`
	// ProgrammaticCreativeCard 程序化创意推广卡片信息
	ProgrammaticCreativeCard []creative.ProgrammaticCreativeCard `json:"programmatic_creative_card,omitempty"`
	// CreativeAutoGenerate 是否开启「生成更多创意」
	CreativeAutoGenerate int `json:"creative_auto_generate,omitempty"`
	// IsHomepageHide 抖音主页是否隐藏视频
	IsHomepageHide int `json:"is_homepage_hide,omitempty"`
}
