package unipromotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialGetRequest 获取全域推广计划下素材 API Request
type MaterialGetRequest struct {
	// Filtering 过滤条件
	Filtering *MaterialGetFilter `json:"filtering,omitempty"`
	// Fields 需要查询的消耗指标，见返回参数，默认stat_cost_for_roi2
	Fields []string `json:"fields,omitempty"`
	// OrderType 排序方式，允许值：
	// ASC 升序（默认）
	// DESC 降序
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// OrderField 排序字段，默认「stat_cost_for_roi2」倒序，同时支持根据消耗指标排序
	// 注意：排序字段必须在fields中
	OrderField string `json:"order_field,omitempty"`
	// AdvertiserID 千川广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：10, 20, 50, 100，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// MaterialGetFilter 过滤条件
type MaterialGetFilter struct {
	// MaterialType 素材类型，可选值:
	// IMAGE 图片，图文
	// LIVE_ROOM 直播间画面
	// TITLE 标题
	// VIDEO 视频
	MaterialType string `json:"material_type,omitempty"`
	// MaterialSelectType 素材类型，可选值:
	// ALL 全部素材
	// CUSTOM 自选投放素材
	// AUTO 智能优选素材
	// 注意：对于直播全域计划，仅material_type=VIDEO支持；对于商品全域计划，material_type=VIDEO/TITLE/IMAGE/CAROUSEL时均支持
	MaterialSelectType string `json:"material_select_type,omitempty"`
	// DELIVERY_OK 投放中，默认
	// DELETED 已删除
	// EXCLUDE 已排除
	// DELIVERY_NOT 不可投
	// ALL 全部
	// 注意：仅material_type=VIDEO/TITLE/CAROUSEL时支持
	MaterialStatus string `json:"material_status,omitempty"`
	// AnalysisType 素材评估，仅material_type=VIDEO时支持
	// 首发素材 FIRST_PUBLISH_MATERIAL
	// 优质素材 HIGH_QUALITY_MATERIAL
	// 低质素材 LOW_QUALITY_MATERIAL
	// 低效素材 INEFFICIENT_MATERIAL
	// 搬运素材 CARRY_MATERIAL
	// 同质化素材 SIMILAR_MATERIAL
	AnalysisType enum.MaterialProperty `json:"analysis_type,omitempty"`
	// SearchKeyword 搜索关键词，支持根据视频mid进行搜索
	// 注意：仅material_type=VIDEO时支持
	SearchKeyword string `json:"search_keyword,omitempty"`
	// ProductKeyword 商品搜索关键词，支持根据商品名称/id进行搜索
	// 注意：仅当marketing_goal=VIDEO_PROM_GOODS支持
	ProductKeyword string `json:"product_keyword,omitempty"`
}

// Encode implements GetRequest interface
func (r MaterialGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Fields != nil {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MaterialGetResponse 获取全域推广计划下素材 API Response
type MaterialGetResponse struct {
	Data *MaterialGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type MaterialGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AdMaterialInfos 素材列表
	AdMaterialInfos []Material `json:"ad_material_infos,omitempty"`
}
