package report

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ProjectGetRequest 获取项目数据 API Request
type ProjectGetRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// TimeGranularity 时间粒度，允许值：
	// TIME_GRANULARITY_DAILY 天维度（默认值）
	// TIME_GRANULARITY_HOURLY 小时维度
	// TIME_GRANULARITY_TOTAL 汇总
	TimeGranularity enum.TimeGranularity `json:"time_granularity,omitempty"`
	// StartDate 查询起始日期，格式：yyyy-mm-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询结束日期，格式：yyyy-mm-dd
	// 当time_granularity = TIME_GRANULARITY_DAILY/TIME_GRANULARITY_TOTAL时，时间跨度不能超过365天
	// 当time_granularity = TIME_GRANULARITY_HOURLY时，时间跨度不能超过7天
	EndDate string `json:"end_date,omitempty"`
	// OrderType 排序方式，允许值：
	// ASC 升序（默认值）
	// DESC 降序
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// OrderField 排序字段，允许值可参考应答返回数据指标
	OrderField string `json:"order_field,omitempty"`
	// Metrics 指标集，允许值可参考应答返回数据指标
	Metrics []string `json:"metrics,omitempty"`
	// Filtering 过滤器
	Filtering *ProjectGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：10（默认值）、20、50、100
	PageSize int `json:"page_size,omitempty"`
}

type ProjectGetFilter struct {
	// CdpProjectIDs 项目ID
	CdpProjectIDs []uint64 `json:"cdp_project_ids,omitempty"`
	// LocalDeliveryScene 推广目的，允许值：
	// CLUE 线索
	// CONTENT_HEATING 内容加热
	// POI_CUSTOMER 门店引流
	// PURCHASE 团购成交
	LocalDeliveryScene local.LocalDeliveryScene `json:"local_delivery_scene,omitempty"`
	// MarketingGoal 营销场景，允许值：
	// LIVE 直播
	// VIDEO_IMAGE 短视频/图文
	MarketingGoal local.MarketingGoal `json:"marketing_goal,omitempty"`
	// CampaignType 广告类型，允许值：
	// GENERAL 通投广告
	// SEARCHING 线索广告
	CampaignType local.AdType `json:"campaign_type,omitempty"`
	// ExternalAction 优化目标
	ExternalAction local.ExternalAction `json:"external_action,omitempty"`
}

// Encode implements GetRequest interface
func (r ProjectGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	if r.TimeGranularity != "" {
		values.Set("time_granularity", string(r.TimeGranularity))
	}
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	values.Set("metrics", string(util.JSONMarshal(r.Metrics)))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// ProjectGetResponse 获取项目数据 API Response
type ProjectGetResponse struct {
	model.BaseResponse
	Data *ProjectGetResult `json:"data,omitempty"`
}

type ProjectGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ProjectList
	ProjectList []Report `json:"project_list,omitempty"`
}
