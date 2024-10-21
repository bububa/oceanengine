package report

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UniPromotionGetRequest 全域推广数据 API Request
type UniPromotionGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	EndDate string `json:"end_date,omitempty"`
	// MarketingGoal 按营销目标过滤，允许值：
	// LIVE_PROM_GOODS：直播带货
	// VIDEO_PROM_GOODS：商品全域
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// LabAdType 推广方式，允许值：LAB_AD 托管
	LabAdType enum.AdLabType `json:"lab_ad_type,omitempty"`
	// Fields 需要查询的消耗指标，具体可参考返回值
	Fields []string `json:"fields,omitempty"`
}

// Encode implement GetRequest interface
func (r UniPromotionGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	if r.MarketingGoal != "" {
		values.Set("marketing_goal", string(r.MarketingGoal))
	}
	if r.LabAdType != "" {
		values.Set("lab_ad_type", string(r.LabAdType))
	}
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// UniPromotionGetResponse 全域推广数据 API Response
type UniPromotionGetResponse struct {
	model.BaseResponse
	Data *UniPromotionStats `json:"data,omitempty"`
}

// UniPromotionStats 全域推广消耗指标
type UniPromotionStats struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音号id
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// StatDatetime 数据起始时间
	// 注意：如果dimension为空，此处返回aweme_id
	StatDatetime string `json:"stat_datetime,omitempty"`
	// StatCost 整体消耗，千分之一分
	StatCost float64 `json:"stat_cost,omitempty"`
	// TotalPrepayAndPayOrderRoi2 整体支付ROI
	TotalPrepayAndPayOrderRoi2 float64 `json:"total_prepay_and_pay_order_roi2,omitempty"`
	// TotalPayOrderGmvForRoi2 整体成交金额
	TotalPayOrderGmvForRoi2 float64 `json:"total_pay_order_gmv_for_roi2,omitempty"`
	// TotalPayOrderCountForRoi2 整体成交订单数
	TotalPayOrderCountForRoi2 int64 `json:"total_pay_order_count_for_roi2,omitempty"`
	// TotalCostPerPayOrderForRoi2 整体成交订单成本
	TotalCostPerPayOrderForRoi2 float64 `json:"total_cost_per_pay_order_for_roi2,omitempty"`
	// TotalPayOrderCouponAmountForRoi2 整体成交智能优惠券金额，单位元，小数点2位
	TotalPayOrderCouponAmountForRoi2 float64 `json:"total_pay_order_coupon_amount_for_roi2,omitempty"`
	// TotalPrepayOrderCountForRoi2 整体预售订单数
	TotalPrepayOrderCountForRoi2 int64 `json:"total_prepay_order_count_for_roi2,omitempty"`
	// TotalUnfinishedEstimateOrderGmvForRoi2 整体未完结预售订单预估金额，单位元，小数点2位
	TotalUnfinishedEstimateOrderGmvForRoi2 float64 `json:"total_unfinished_estimate_order_gmv_for_roi2,omitempty"`
	// TotalPrepayOrderGmvForRoi2 整体预售订单成本
	TotalPrepayOrderGmvForRoi2 float64 `json:"total_prepay_order_gmv_for_roi2,omitempty"`
}
