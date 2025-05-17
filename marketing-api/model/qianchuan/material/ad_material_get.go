package material

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdMaterialGetRequest 获取计划下素材列表 API Request
type AdMaterialGetRequest struct {
	// Filtering 过滤条件
	Filtering *AdMaterialGetFiltering `json:"filtering,omitempty"`
	// OrderType 排序方式 可选值:
	// ASC 升序
	// DESC 降序 默认
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// OrderField 排序字段，支持根据消耗等数据指标排序
	OrderField string `json:"order_field,omitempty"`
	// AdvertiserID 千川广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：10, 20, 50, 100，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// AdMaterialGetFiltering 过滤条件
type AdMaterialGetFiltering struct {
	// MaterialType 素材类型 可选值:
	// IMAGE 图片，图文
	// TITLE 标题
	// LIVE_ROOM 直播间画面
	// VIDEO 视频
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
	// HavingCost 消耗情况，仅material_type=VIDEO/IMAGE时支持
	// 全部 ALL
	// 有消耗 YES
	// 默认查询全部
	HavingCost string `json:"having_cost,omitempty"`
	// SearchKeyword 搜索关键词
	// 支持查询直播间/视频/标题/图片名称、直播间/视频/图片id
	SearchKeyword string `json:"search_keyword,omitempty"`
	// StartTime 数据开始时间
	// 注意：仅having_cost入参时，该筛选项生效
	StartTime string `json:"start_time,omitempty"`
	// EndTime 数据结束时间
	// 注意：仅having_cost入参时，该筛选项生效
	EndTime string `json:"end_time,omitempty"`
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
	VideoSource []enum.MaterialSource `json:"material_source,omitempty"`
	// AnalysisType 素材建议，仅material_type=VIDEO时支持
	//  CARRY_MATERIAL 搬运风险素材
	// LOW_EFFICIENCY_MATERIAL 低效素材
	// FIRST_PUBLISH_MATERIAL 首发素材
	// SIMILAR_RISK_MATERIAL 同质化素材
	// HIGH_QUALITY_MATERIAL 优质素材
	// POOR_QUALITY_MATERIAL 低质素材
	AnalysisType []enum.MaterialProperty `json:"analysis_type,omitempty"`
}

// Encode implements GetRequest interface
func (r AdMaterialGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// AdMaterialGetResponse 获取计划下素材列表 API Response
type AdMaterialGetResponse struct {
	Data *AdMaterialGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type AdMaterialGetResult struct {
	// PageInfo 分页结果
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AdMaterialInfos 返回的素材信息列表
	AdMaterialInfos []AdMaterialInfo `json:"ad_material_infos,omitempty"`
}
