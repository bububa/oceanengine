package material

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdGetRequest 获取素材关联计划 API Request
type AdGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// MaterialType 素材类型，可选值:
	// IMAGE 图片，图文
	// LIVE_ROOM 直播间画面
	// TITLE 标题
	// VIDEO 视频
	MaterialType MaterialType `json:"material_type,omitempty"`
	// MarketingScene 广告类型，可选值:
	// FEED 通投
	// SEARCH 搜索
	// SHOPPING_MALL 商城广告
	MarketingScene qianchuan.MarketingScene `json:"marketing_scene,omitempty"`
	// MarketingGoal 营销场景，可选值:
	// VIDEO_PROM_GOODS：推商品
	// LIVE_PROM_GOODS：推直播间
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// StartTime 计划数据开始时间，格式“YYYY-MM-DD”
	// 注意：最早开始时间不大于“当前时间-180天”
	StartTime string `json:"start_time,omitempty"`
	// EndTime 计划数据结束时间，格式“YYYY-MM-DD”
	EndTime string `json:"end_time,omitempty"`
	// Fields 需要查询的消耗指标，取值见返回值中metric相关指标
	Fields []string `json:"fields,omitempty"`
	// OrderField 排序字段
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序方式 可选值:
	// ASC 升序
	// DESC 降序
	OrderType enum.OrderType `json:"order_type,omitempty"`
}

// Encode implements GetRequest interface
func (r AdGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("material_id", strconv.FormatUint(r.MaterialID, 10))
	values.Set("material_type", string(r.MaterialType))
	values.Set("marketing_scene", string(r.MarketingScene))
	values.Set("marketing_goal", string(r.MarketingGoal))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AdGetResponse 获取素材关联计划 API Response
type AdGetResponse struct {
	Data *AdGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type AdGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AdList 计划列表
	AdList []MaterialAd `json:"ad_list,omitempty"`
}

// MaterialAd 素材关联计划
type MaterialAd struct {
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// AdName 计划名称
	AdName string `json:"ad_name,omitempty"`
	// Metrics 指标
	Metrics *report.Metrics `json:"metrics,omitempty"`
}
