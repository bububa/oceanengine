package ad

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/creative"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建计划（含创意生成规则）
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MarketingGoal 营销目标
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// CampaignScene 营销场景 ，允许值：
	// DAILY_SALE日常销售
	CampaignScene qianchuan.CampaignScene `json:"campaign_scene,omitempty"`
	// MarketingScene 广告类型，允许值：
	// FEED 通投广告
	MarketingScene qianchuan.MarketingScene `json:"marketing_scene,omitempty"`
	// PromotionWary 推广方式
	PromotionWay enum.PromotionWay `json:"promotion_way,omitempty"`
	// Name 计划名称，长度为1-100个字符，其中1个汉字算2位字符。名称不可重复，否则会报错
	Name string `json:"name,omitempty"`
	// LabAdType 推广方式，允许值：
	// LAB_AD 托管
	LabAdType enum.AdLabType `json:"lad_ad_type,omitempty"`
	// CampaignID 千川广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// AwemeID 抖音id，即商品广告背后关联的抖音号，可通过【查询可推广抖音号列表】接口获取名下可推广抖音号
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// ProductIDs 商品id列表，即准备推广的商品列表，可通过【查询店铺商品列表】接口获取名下可推广商品; 目前仅支持推一个商品，但需以数组入参
	ProductIDs []uint64 `json:"product_ids,omitempty"`
	// ChannelProductInfos 渠道品信息
	// 注意：如果当前商品在该抖音号下存在渠道品，需要入參channel_product_info
	// 渠道品相关介绍见《【抖店】销售渠道品功能操作手册》
	ChannelProductInfos []ChannelProduct `json:"channel_product_infos,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting DeliverySetting `json:"delivery_setting,omitempty"`
	// Audience 人群定向
	Audience *Audience `json:"audience,omitempty"`
	// CreativeMaterialMode 创意呈现方式，和抖音号关系类型相关
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// FirstIndustryID 创意一级行业ID
	FirstIndustryID uint64 `json:"first_industry_id,omitempty"`
	// SecondIndustryID 创意二级行业ID
	SecondIndustryID uint64 `json:"second_industry_id,omitempty"`
	// ThirdIndustryID 创意三级行业ID
	ThirdIndustryID uint64 `json:"third_industry_id,omitempty"`
	// AdKeywords 创意标签。最多20个标签，且每个标签长度要求为1~20个字符，汉字算2个字符
	AdKeywords []string `json:"ad_keywords,omitempty"`
	// CreativeList 自定义素材信息
	CreativeList []creative.Creative `json:"creative_list,omitempty"`
	// CreativeAutoGenerate是否开启「生成更多创意」，允许值：0 关闭（默认值）、1 开启
	CreativeAutoGenerate int `json:"creative_auto_generate,omitempty"`
	// IsHomepageHide 抖音主页是否隐藏视频，和抖音号关系类型相关，返回值参考【附录-抖音号授权类型】;bind_type为OFFICIAL或SELF时，允许值：1 隐藏、0 不隐藏（默认值）;bind_type不为OFFICIAL或SELF时，需传入唯一允许值0 不隐藏
	IsHomePageHide int `json:"is_homepage_hide,omitempty"`
	// ProgrammaticCreativeMadiaList 程序化创意素材信息，最多支持 9 个创意。当 creative_material_mode 不为 PROGRAMMATIC_CREATIVE 时，该字段不填数据；当 creative_material_mode 为 PROGRAMMATIC_CREATIVE 时，该字段必填;请至少添加一个视频类型素材
	ProgrammaticCreativeMadiaList []creative.ProgrammaticCreativeMedia `json:"programmatic_creative_media_list,omitempty"`
	// ProgrammaticCreativeTitleList 程序化创意标题信息，最多支持 10 个标题。当 creative_material_mode 不为 PROGRAMMATIC_CREATIVE 时，该字段不填数据；当 creative_material_mode 为 PROGRAMMATIC_CREATIVE 时，该字段必填; 请至少添加一个标题
	ProgrammaticCreativeTitleList []creative.TitleMaterial `json:"programmatic_creative_title_list,omitempty"`
	// ProgrammaticCreativeCard 程序化创意推广卡片信息。当 creative_material_mode 不为 PROGRAMMATIC_CREATIVE 时，该字段不填数据；当 creative_material_mode 为 PROGRAMMATIC_CREATIVE 时，该字段必填
	ProgrammaticCreativeCard *creative.ProgrammaticCreativeCard `json:"programmatic_creative_card,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建计划（含创意生成规则）
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CreateResult `json:"data,omitempty"`
}

type CreateResult struct {
	// AdID 创建的计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// NoticeInfos 提示信息
	NoticeInfos []NoticeInfo `json:"notice_info,omitempty"`
}

// NoticeInfo 提示信息
type NoticeInfo struct {
	// Code 提示错误码
	Code int `json:"code,omitempty"`
	// Message 提示错误信息
	Message string `json:"message,omitempty"`
	// SearchKeywordError 搜索关键词错误列表，仅marketing_scene=SEARCH 情况下会返回该信息
	SearchKeywordError []SearchKeywordError `json:"search_keyword_error,omitempty"`
}

// SearchKeywordError 搜索关键词错误
type SearchKeywordError struct {
	// SearchKeyword 错误的搜索关键词
	SearchKeyword string `json:"search_keyword,omitempty"`
	// ErrorMessage 错误的原因
	ErrorMessage string `json:"error_message,omitempty"`
}
