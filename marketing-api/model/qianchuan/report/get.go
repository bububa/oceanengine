package report

import (
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取数据报表API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	EndDate time.Time `json:"end_date,omitempty"`
	// TimeGranularity 时间粒度 ，如果不传，返回查询日期内的聚合数据
	// 允许值:
	// TIME_GRANULARITY_DAILY (按天维度),会返回每天的数据
	// TIME_GRANULARITY_HOURLY (按小时维度)，会返回每小时维度的数据
	TimeGranularity enum.TimeGranularity `json:"time_granularity,omitempty"`
	// WordType 词类型，允许值：
	// SEARCH_WORD 搜索词
	// KEY_WORD 关键词
	WordType qianchuan.WordType `json:"word_type,omitempty"`
	// Fields 指定需要的指标名称
	Fields []string `json:"fields,omitempty"`
	// OrderField 排序字段，所有的统计指标均可参与排序
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序方式；默认值: DESC；允许值: ASC, DESC
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页码；默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，即每页展示的数据量；默认值: 20；取值范围: 1-1000
	PageSize int `json:"page_size,omitempty"`
	// Filtering 过滤字段，json格式，支持字段如下
	Filtering *StatFiltering `json:"filtering,omitempty"`
}

// StatFiltering 数据报表过滤条件
type StatFiltering struct {
	// CampaignIDs 广告组id列表：按照campaign_id过滤，最多支持100个
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// AdIDs 广告计划id列表：按照 ad_id 过滤，最多支持100个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// CreativeIDs 广告创意id列表：按照 creative_id 过滤，最多支持100个
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// AwemeIDs 按抖音id过滤，即关联的抖音号
	AwemeIDs []uint64 `json:"aweme_ids,omitempty"`
	// MarketingGoal 营销目标
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// OrderPlatform 下单平台，允许值：
	// ALL：全部
	// QIANCHUAN： 千川pc（默认）
	// ECP_AWEME：小店随心推
	OrderPlatform qianchuan.OrderPlatform `json:"order_platform,omitempty"`
	// MarketingScene 营销场景，允许值：
	// ALL：全部
	// FEED： 通投广告
	// SEARCH：搜索广告
	// 注意：当下单平台为“小店随心推”时，不支持
	MarketingScene qianchuan.MarketingScene `json:"marketing_scene,omitempty"`
	// CampaignScene 营销场景，可选值
	// DAILY_SALE：日常销售（默认）
	// 注意：当下单平台为“小店随心推”时，不支持
	CampaignScene qianchuan.CampaignScene `json:"campaign_scene,omitempty"`
	// ExternalAction 优化目标，如果不填，默认查的是当前营销目标下的所有优化目标
	// 当营销目标为直播带货时，允许值为：
	// AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION: 直播间商品点击
	// AD_CONVERT_TYPE_LIVE_COMMENT_ACTION: 直播间评论
	// AD_CONVERT_TYPE_LIVE_ENTER_ACTION: 进入直播间
	// AD_CONVERT_TYPE_LIVE_ROI: 直播间支付ROI
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION: 直播间下单
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY: 直播间成交
	// AD_CONVERT_TYPE_NEW_FOLLOW_ACTION: 直播间粉丝提升
	// 当营销目标为短视频带货时，允许值：
	// AD_CONVERT_TYPE_QC_FOLLOW_ACTION: 粉丝提升
	// AD_CONVERT_TYPE_QC_MUST_BUY: 点赞评论
	// AD_CONVERT_TYPE_SHOPPING: 商品购买
	// ROIAD_CONVERT_TYPE_VIDEO_ROI:商品支付
	ExternalAction qianchuan.ExternalAction `json:"external_action,omitempty"`
	// CreativeMaterialMode 按创意类型过滤，允许值：
	// CUSTOM_CREATIVE：自定义创意
	// PROGRAMMATIC_CREATIVE： 程序化创意
	// 注意：
	// 1.当下单平台为“小店随心推”时，不支持
	// 2.当创意类型为CUSTOM_CREATIVE时，包含高光快投创意
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// ImageMode 按素材样式过滤，支持视频和图片，枚举值见【附录-枚举值-素材类型】（暂不支持“小图”和“穿山甲开屏图片”）
	// 注意：
	// 1.当下单平台为“小店随心推”时，不支持
	// 2.素材样式为竖版视频时，包含高光快投创意
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// SmartBidType 投放场景（投放方式），允许值：
	// SMART_BID_CUSTOM：控成本投放
	// SMART_BID_CONSERVATIVE： 放量投放
	// 注意：当下单平台为“小店随心推”或营销场景为“搜索广告”时，不支持
	SmartBidType enum.SmartBidType `json:"smart_bid_type,omitempty"`
	// Status 按计划状态过滤，不传入即默认返回“全部（包含已删除）”，其他规则详见【附录-广告计划查询状态】（暂不支持“系统暂停”和“在投计划配额超限”）
	// 注意：当下单平台为“小店随心推”时，不支持
	Status qianchuan.AdStatusForSearch `json:"status,omitempty"`
	// IsHighlight 是否返回高光快投创意，允许值：
	// ALL：包含高光快投创意
	// ONLY_HIGHLIGHT：只返回高光快投创意
	// NO_HIGHLIGHT：返回非高光快投创意，默认值
	IsHighlight string `json:"is_highlight,omitempty"`
	// MaterialType 素材类型，允许值：
	// video ：视频素材
	// image ：图片素材
	// carousel 图文素材
	MaterialType qianchuan.MaterialType `json:"material_type,omitempty"`
	// VideoSource 视频来源筛选，以平台素材库接口的视频来源枚举值为准，允许值：
	// E_COMMERCE：本地上传
	// LIVE_HIGHLIGHT：直播剪辑素材
	// BP：巨量纵横共享素材
	// VIDEO_CAPTURE：易拍APP共享素材
	// ARTHUR：亚瑟共享素材
	// STAR：星图&即合共享素材
	// TADA：tada共享素材
	// CREATIVE_CENTER：巨量创意PC共享素材
	// 注意：仅素材类型为视频素材时，支持
	VideoSource enum.MaterialSource `json:"video_source,omitempty"`
	// MaterialID 素材id列表，一个素材唯一对应一个素材id，相同素材上传多次对应一个material_id
	MaterialID []uint64 `json:"material_id,omitempty"`
	// MaterialMode 素材样式筛选
	// 素材类型为视频素材时允许值：
	// VIDEO_LARGE：横版视频
	// VIDEO_VERTICAL：竖版视频
	// 素材类型为图片素材时允许值：
	// SMALL：横版小图
	// LARGE：横版大图
	// LARGE_VERTICAL：竖版图片
	// SQUARE ：商品卡方图
	MaterialMode []enum.MaterialMode `json:"material_mode,omitempty"`
	// ImageSource 图片来源筛选，允许值
	// E_COMMERCE：本地上传
	// CREATIVE_CENTER：巨量创意PC共享素材
	// SQUARE：商品图
	// JI_CHUANG：即创共享素材
	ImageSource []enum.MaterialSource `json:"image_source,omitempty"`
	// CarouselSource 图文来源
	// JI_CHUANG：即创共享素材
	// AWEME：抖音主页视频
	// E_COMMERCE：本地上传
	CarouselSource []enum.MaterialSource `json:"carousel_source,omitempty"`
	// AnalysisType 素材建议，允许值
	// first_publish：首发素材
	// high_quality：优质素材
	// low_efficiency：低效素材
	// poor_quality：低质素材
	// improvable：可提升素材
	AnalysisType []string `json:"analysis_type,omitempty"`
	// WordType 词类型，当word有传值时，必填，允许值：
	// SEARCH_WORD搜索词
	// KEY_WORD关键词
	WordType qianchuan.WordType `json:"word_type,omitempty"`
	// Word 具体需要过滤的词
	Word string `json:"word,omitempty"`
	// KeywordMerge 合并关键词，可选值
	// true、false（默认）
	// 注意：打开后，将以关键词维度聚合数据，汇总相同关键词在不同计划下的数据
	KeywordMerge bool `json:"keyword_merge,omitempty"`
	// Range 更多过滤条件，支持根据词的消耗、展示次数、转化数、转换成本进行过滤
	Range *FilterRange `json:"range,omitempty"`
}

// FilterRange 更多过滤条件，支持根据词的消耗、展示次数、转化数、转换成本进行过滤
type FilterRange struct {
	// Field 需要过来的指标，允许值：
	// 消耗：stat_cost
	// 展示次数：show_cnt
	// 转化数：convert_cnt
	// 转换成本：cpa_platform
	Field string `json:"field,omitempty"`
	// Calculation 计算方式，允许值：
	// UPPER 大于等于
	// LOWER 小于等于
	Calculation string `json:"calculation,omitempty"`
	// Value 数值范围
	Value string `json:"value,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_date", r.StartDate.Format("2006-01-02"))
	values.Set("end_date", r.EndDate.Format("2006-01-02"))
	if r.TimeGranularity != "" {
		values.Set("time_granularity", string(r.TimeGranularity))
	}
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if r.WordType != "" {
		values.Set("word_type", string(r.WordType))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 数据报表API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// Data json返回值
type GetResponseData struct {
	// List 数据列表
	List []Stat `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
