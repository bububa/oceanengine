package ad

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/brand"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/creative"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/live"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/product"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/shop"
)

// Ad 计划详情
type Ad struct {
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// CampaignID 广告组ID（若为托管计划，则返回null）
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// MarketingGoal 营销目标
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// CampaignScene 营销场景
	CampaignScene qianchuan.CampaignScene `json:"campaign_scene,omitempty"`
	// MarketingScene 广告类型，FEED 通投广告，SEARCH 搜索广告
	MarketingScene qianchuan.MarketingScene `json:"marketing_scene,omitempty"`
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
	AdModifyTime string `json:"ad_modify_time,omitempty"`
	// AwemeInfo 计划中关联的抖音号信息
	AwemeInfo []aweme.Aweme `json:"aweme_info,omitempty"`
	// ProductInfo 商品列表
	ProductInfo []product.Product `json:"product_info,omitempty"`
	// RoomInfo 直播间列表
	RoomInfo []live.Room `json:"room_info,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// Audience 定向设置
	Audience *Audience `json:"audience,omitempty"`
	// ShopInfo 店铺信息
	ShopInfo *shop.Shop `json:"shop_info,omitempty"`
	// BrandInfo 品牌信息
	BrandInfo *brand.Brand `json:"brand_info,omitempty"`
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
	// IsIntelligent 是否启用智选流量，0 关闭、1 开启
	IsIntelligent int `json:"is_intelligent,omitempty"`
	// DynamicCreative 是否启用动态创意，0 关闭、1 开启
	DynamicCreative int `json:"dynamic_creative,omitempty"`
	// Keywords 搜索关键词列表
	Keywords []Keyword `json:"keyword,omitempty"`
	// PrivateWords 搜索否定词
	PrivateWords *PrivateWords `json:"private_words,omitempty"`
	// LabAdType 推广方式，NOT_LAB_AD：非托管计划，LAB_AD：托管计划
	LabAdType enum.AdLabType `json:"lab_ad_type,omitempty"`
}

// Keyword 搜索关键词
type Keyword struct {
	// ID 关键词id
	ID uint64 `json:"id,omitempty"`
	// WordID 词id，不同计划下如果关键词字面相同，词id会相同
	WordID uint64 `json:"word_id,omitempty"`
	// Word 关键词字面
	Word string `json:"word,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// Status 关键词状态
	// CONFIRM 审核通过且可代入
	// REJECT 审核拒绝
	// AUDIT 新建审核中
	// DELETE 已删除
	// PAUSED 词暂停
	Status qianchuan.KeywordStatus `json:"status,omitempty"`
	// MatchType 匹配类型，PHRASE 短语匹配、EXTENSIVE 广泛匹配、PRECISION 精准匹配
	MatchType enum.KeywordMatchType `json:"match_type,omitempty"`
	// WordPackageID 词包ID
	WordPackageID uint64 `json:"word_package_id,omitempty"`
	// WordPackageName 词包名称
	WordPackageName string `json:"word_package_name,omitempty"`
}

// PrivateWords 搜索否定词
type PrivateWords struct {
	// PhraseWords 短语否定词列表
	PhraseWords []string `json:"phrase_words,omitempty"`
	// PreciseWords 精确否定词列表
	PreciseWords []string `json:"precise_words,omitempty"`
}
