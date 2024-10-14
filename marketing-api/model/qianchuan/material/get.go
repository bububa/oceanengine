package material

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取账户下素材列表 API Request
type GetRequest struct {
	// AdvertiserID 千川广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MarketingGoal
	// LIVE_PROM_GOODS 推直播间
	// VIDEO_PROM_GOODS 推商品
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// CampaignScene 按营销场景过滤，不传，默认查询全部，允许值：
	// 直播场景
	//   DAILY_SALE: 日常销售
	// LIVE_HEAT: 直播间加热
	// 推商品场景
	//   DAILY_SALE: 日常销售
	//   PLANT_GRASS: 人群种草
	CampaignScene qianchuan.CampaignScene `json:"campaign_scene,omitempty"`
	// MarketingScene 广告类型过滤，可选值:
	// FEED：通投
	// SEARCH：搜索
	// SHOPPING_MALL ：商城广告
	MarketingScene qianchuan.MarketingScene `json:"marketing_scene,omitempty"`
	// Filtering 过滤器
	Filtering *GetFilter `json:"filtering,omitempty"`
	// Fields 需要查询的消耗指标
	// 注意：不同素材类型支持的指标有所差异，具体见返回metrics指标
	Fields []string `json:"fields,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize  页面大小，允许值：10, 20, 50, 100，默认值：10
	PageSize int `json:"page_size,omitempty"`
	// OrderType 排序方式 可选值:
	// ASC升序
	// DESC降序（默认）
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// OrderField 排序字段
	// 注意：仅支持根据metrics中字段进行排序，默认stat_cost
	OrderField string `json:"order_field,omitempty"`
}

type GetFilter struct {
	// MaterialType 素材类型，可选值:
	// IMAGE 图片，图文
	// LIVE_ROOM 直播间画面
	// TITLE 标题
	// VIDEO 视频
	// 注意：直播间画面 仅支持推直播间计划，图片仅支持推商品计划
	MaterialType MaterialType `json:"material_type,omitempty"`
	// ImageMode 素材样式，仅material_type=VIDEO/IMAGE时支持
	// 当material_type=VIDEO时，支持如下
	// 横版视频 VIDEO_LARGE
	// 竖版视频  VIDEO_VERTICAL
	// 当material_type=IMAGE时，支持如下
	// 横版小图 SMALL
	// 横版大图 LARGE
	// 竖版图片LARGE_VERTICAL
	// 图文CAROUSEL
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// HavingCost  消耗情况，仅material_type=VIDEO/IMAGE时支持
	HavingCost string `json:"having_cost,omitempty"`
	// VideoSource 视频来源，仅material_type=VIDEO时支持
	// AWEME 抖音主页视频
	// E_COMMERCE 本地上传
	// LIVE_HIGHLIGHT 直播剪辑素材
	// BP 巨量纵横共享素材
	// VIDEO_CAPTURE 易拍APP共享素材
	// ARTHUR 亚瑟共享素材
	// STAR 星图&即合共享素材
	// TADA tada共享素材
	// CREATIVE_CENTER 巨量创意PC共享素材
	// JIANYING 剪映共享素材
	// JI_CHUANG 即创共享素材
	// QUNFENG 群峰共享素材
	VideoSource enum.MaterialSource `json:"video_source,omitempty"`
	// AnalysisType 素材建议，仅material_type=VIDEO时支持
	// CARRY_MATERIAL 搬运风险素材
	// LOW_EFFICIENCY_MATERIAL 低效素材
	// FIRST_PUBLISH_MATERIAL 首发素材
	// SIMILAR_RISK_MATERIAL 同质化素材
	// HIGH_QUALITY_MATERIAL 优质素材
	// POOR_QUALITY_MATERIAL 低质素材
	AnalysisType enum.MaterialProperty `json:"analysis_type,omitempty"`
	// SearchKeyword 搜索关键词
	// 支持查询直播间/视频/标题/图片名称、直播间/视频/图片id
	SearchKeyword string `json:"search_keyword,omitempty"`
	// StartTime 数据查询开始时间，精确到秒，yyyy-MM-dd HH:mm:ss
	StartTime string `json:"start_time,omitempty"`
	// EndTime 数据查询结束时间，精确到秒，yyyy-MM-dd HH:mm:s
	EndTime string `json:"end_time,omitempty"`
}

// Encode implements GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("marketing_goal", string(r.MarketingGoal))
	if r.MarketingScene != "" {
		values.Set("marketing_scene", string(r.MarketingScene))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Fields != nil {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 获取素材列表 API Response
type GetResponse struct {
	Data *GetResult `json:"data,omitempty"`
	model.BaseResponse
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AdMateriaInfos 返回的素材信息列表             l
	AdMaterialInfos []AdMaterialInfo `json:"ad_material_infos,omitempty"`
}
