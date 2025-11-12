package unipromotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrderReportGetRequest 获取随心推全域订单数据 API Request
type OrderReportGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 查询数据开始日期，例如2021-04-05，开始时间不可超过「当前日期-180天」
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询数据结束日期，例如2021-04-05，日期跨度不能超过180天
	EndDate string `json:"end_date,omitempty"`
	// TimeGranularity 时间粒度 ，如果不传，返回查询日期内的聚合数据 可选值:
	// TIME_GRANULARITY_DAILY：(按天维度),会返回每天的数据
	// TIME_GRANULARITY_HOURLY：(按小时维度)，会返回每小时维度的数据
	TimeGranularity enum.TimeGranularity `json:"time_granularity,omitempty"`
	// Filter 过滤条件
	Filter *OrderReportGetFilter `json:"filter,omitempty"`
	// OrderType 排序类型 可选值:
	// ASC 升序
	// DESC 降序（默认）
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// OrderField 排序字段 可选值:
	// total_pay_order_gmv_for_roi2 用户实际支付金额
	// stat_cost_for_roi2 整体消耗
	// total_prepay_and_pay_order_roi2 整体支付ROI
	// total_pay_order_count_for_roi2 整体成交订单数
	// total_pay_order_coupon_amount_for_roi2 整体成交智能优惠券金额
	// total_unfinished_estimate_order_gmv_for_roi2 整体未完结预售订单预估金额
	// total_ecom_platform_subsidy_amount_for_roi2 平台补贴金额
	// total_pay_order_gmv_include_coupon_for_roi2 整体成交金额
	OrderField string `json:"order_field,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：10, 20, 50, 100，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

type OrderReportGetFilter struct {
	// AdID 随心推商品全域计划id（非order_id），最多支持100个
	AdID []uint64 `json:"ad_id,omitempty"`
}

// Encode implements GetRequest
func (r OrderReportGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	if r.TimeGranularity != "" {
		values.Set("time_granularity", string(r.TimeGranularity))
	}
	if r.Filter != nil {
		values.Set("filter", string(util.JSONMarshal(r.Filter)))
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

// OrderReportGetResponse 获取随心推全域订单数据 API Response
type OrderReportGetResponse struct {
	model.BaseResponse
	Data *OrderReportGetResult `json:"data,omitempty"`
}

type OrderReportGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AdList 消耗指标
	AdList []OrderStat `json:"ad_list,omitempty"`
}

// OrderStat 消耗指标
type OrderStat struct {
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// StatDateTime 数据起始时间，当time_granularity有入參时返回
	StatDateTime string `json:"stat_date_time,omitempty"`
	// TotalPayOrderGmvForRoi2 用户实际支付金额，单位元
	TotalPayOrderGmvForRoi2 float64 `json:"total_pay_order_gmv_for_roi2,omitempty"`
	// StatCostForRoi2 整体消耗，单位元
	StatCostForRoi2 float64 `json:"stat_cost_for_roi2,omitempty"`
	// TotalPrepayAndPayOrderRoi2 整体支付ROI
	TotalPrepayAndPayOrderRoi2 float64 `json:"total_prepay_and_pay_order_roi2,omitempty"`
	// TotalPayOrderCountForRoi2 整体成交订单数
	TotalPayOrderCountForRoi2 int64 `json:"total_pay_order_count_for_roi2,omitempty"`
	// TotalPayOrderCouponAmountForRoi2 整体成交智能优惠券金额，单位元
	TotalPayOrderCouponAmountForRoi2 float64 `json:"total_pay_order_coupon_amount_for_roi2,omitempty"`
	// TotalUnfinishedEstimateOrderGmvForRoi2 整体未完结预售订单预估金额，单位元
	TotalUnfinishedEstimateOrderGmvForRoi2 float64 `json:"total_unfinished_estimate_order_gmv_for_roi2,omitempty"`
	// TotalEcomPlatformSubsidyAmountForRoi2 平台补贴金额，单位元
	TotalEcomPlatformSubsidyAmountForRoi2 float64 `json:"total_ecom_platform_subsidy_amount_for_roi2,omitempty"`
	// TotalPayOrderGmvIncludeCouponForRoi2 整体成交金额，单位元
	TotalPayOrderGmvIncludeCouponForRoi2 float64 `json:"total_pay_order_gmv_include_coupon_for_roi2,omitempty"`
}
