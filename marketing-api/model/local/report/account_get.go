package report

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AccountGetRequest 查询账户数据 API Request
type AccountGetRequest struct {
	// LocalAccountID 本地推投放账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// TimeGranularity 时间粒度，如果不传，返回查询日期内的聚合数据，允许值：
	// TIME_GRANULARITY_DAILY（按天维度），会返回账户每天的数据
	// TIME_GRANULARITY_HOURLY（按小时维度），会返回账户每小时的数
	// TIME_GRANULARITY_TOTAL（汇总），会返回账户汇总的数据
	TimeGranularity enum.TimeGranularity `json:"time_granularity,omitempty"`
	// StartDate 投放开始时间，格式 2021-04-05，开始时间不得早于今日-365天
	StartDate string `json:"start_date,omitempty"`
	// EndDate 投放结束时间，格式 2021-04-05
	// 若不传time_granularity，则时间跨度不能超过365天；
	// 若传time_granularity为TIME_GRANULARITY_DAILY时，则时间跨度不能超过365天；
	// 若传time_granularity为TIME_GRANULARITY_TOTAL时，则时间跨度不能超过365天；
	// 若传time_granularity为TIME_GRANULARITY_HOURLY时，则时间跨度不能超过7天
	EndDate string `json:"end_date,omitempty"`
	// Metrics 指标集，允许值可参考应答返回数据指标
	Metrics []string `json:"metrics,omitempty"`
	// OrderType 排序方式，允许值：
	// ASC：升序
	// DESC：降序（默认）
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// OrderField 排序字段，允许值参考数据指标，默认根据时间排序
	// 注意：如果根据指标排序，order_field必须在metrics范围内
	OrderField string `json:"order_field,omitempty"`
	// Filtering 过滤条件
	Filtering *AccountGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：10（默认值）、20、50、100
	PageSize int `json:"page_size,omitempty"`
}

type AccountGetFilter struct {
	// MarketingGoal 营销场景，允许值：
	// LIVE 直播
	// VIDEO_IMAGE 短视频/图文
	MarketingGoal local.MarketingGoal `json:"marketing_goal,omitempty"`
	// LocalDeliveryScene 营销目的，不传，默认查询全部，可选值:
	// CLUE 获取线索
	// CONTENT_HEATING 线上互动
	// POI_CUSTOMER 线下到店
	// PURCHASE 团购成交
	LocalDeliveryScene local.LocalDeliveryScene `json:"local_delivery_scene,omitempty"`
	// CampaignType 单元类型，不传，默认查询全部，可选值:
	// GENERAL 通投
	// SEARCHING 搜索
	CampaignType local.AdType `json:"campaign_type,omitempty"`
	// ExternalAction 优化目标，不传，默认查询全部，可选值:
	// CLUE_ACQUISITION 获取线索
	// FOLLOW_ACTION 粉丝增长
	// CLUE_CONFIRM确认意向
	// LIVE_ENGAGEMENT 直播加热
	// LIVE_ENTER_ACTION 直播间观看
	// LIVE_OTO_CLICK 直播间团购点击
	// LIVE_OTO_GROUP_BUYING 直播间团购购买
	// LIVE_STAY_TIME 直播间停留
	// NATIVE_ACTION 用户互动
	// OTO_PAY 团购成交
	// POI_RECOMMEND 门店浏览
	// PRIVATE_MESSAGE 私信消息
	// SHOW 展示量
	ExternalAction local.ExternalAction `json:"external_action,omitempty"`
	// DeliveryMode 投放模式，不传，默认查询全部，可选值:
	// CDP_AUTO_MODE 自动投放
	// MANUAL_MODE 手动投放
	DeliveryMode local.DeliveryMode `json:"devlivery_mode,omitempty"`
}

// Encode implements GetRequest interface
func (r AccountGetRequest) Encode() string {
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

// AccountGetResponse 查询账户数据 API Response
type AccountGetResponse struct {
	model.BaseResponse
	Data *AccountGetResult `json:"data,omitempty"`
}

type AccountGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// DataList
	DataList []Report `json:"project_list,omitempty"`
}
