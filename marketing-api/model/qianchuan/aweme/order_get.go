package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrderGetRequest 获取随心推订单列表 API Request
type OrderGetRequest struct {
	// Filtering
	Filtering *OrderGetFilter `json:"filtering,omitempty"`
	// OrderField 排序字段。默认升序 可选值:
	// order_create_time 按照创建时间排序
	// pay_order_amount 按照成交金额排序
	// prepay_and_pay_order_roi 按照支付roi排序
	// stat_cost 按照消耗排序
	OrderField string `json:"order_field,omitempty"`
	// OrderCreateStartTime 计划创建开始时间，格式："yyyy-mm-dd"
	OrderCreateStartTime string `json:"order_create_start_time,omitempty"`
	// OrderCreateEndTime 计划创建结束时间，与order_create_start_date搭配使用，格式："yyyy-mm-dd"，时间跨度不能超过180天
	OrderCreateEndTime string `json:"order_create_end_time,omitempty"`
	// AdvertiserID 千川广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Cursor 页码，默认值：0
	Cursor int `json:"cursor,omitempty"`
	// Count 页面大小，允许值：10, 20, 50，默认值：10
	Count int `json:"count,omitempty"`
}

type OrderGetFilter struct {
	// MarketingGoal 按营销目标过滤，允许值：
	// VIDEO_PROM_GOODS：短视频带货
	// LIVE_PROM_GOODS：直播带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// Status 小店订单状态 可选值:
	// AUDIT 审核中
	// BOOK 未开播
	// CREATING 已支付，计划创建中
	// DELIVERY_OK 投放中
	// FAILED 计划创建失败
	// FINISHED 投放完成
	// FROZEN 投放终止
	// OFFLINE_AUDIT 审核不通过
	// OVER 投放结束
	// TIMEOUT 未支付超过15分钟，订单关闭
	// UNPAID 未支付
	// UNPAIDCANCEL 未支付取消订单
	Status qianchuan.OrderStatus `json:"status,omitempty"`
}

// Encode implement GetRequest interface
func (r OrderGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Cursor > 0 {
		values.Set("cursor", strconv.Itoa(r.Cursor))
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderCreateStartTime != "" {
		values.Set("order_create_start_time", r.OrderCreateStartTime)
	}
	if r.OrderCreateEndTime != "" {
		values.Set("order_create_end_time", r.OrderCreateEndTime)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OrderGetResponse 获取随心推订单列表 API Response
type OrderGetResponse struct {
	Data *OrderGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type OrderGetResult struct {
	// PageInfo
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List
	List []Order `json:"list,omitempty"`
	// FailList 获取失败的订单id列表
	FailList []uint64 `json:"fail_list,omitempty"`
}
