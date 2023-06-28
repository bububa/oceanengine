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
	// AdvertiserID 千川广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering
	Filtering *OrderGetFilter `json:"filtering,omitempty"`
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
	// AUDIT: 审核中
	// BOOK: 未开播
	// DELIVERY_OK: 投放中
	// FINISHED: 投放完成
	// FROZEN: 投放终止
	// OFFLINE_AUDIT: 审核不通过
	// OVER: 投放结束
	// UNPAID: 未支付
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
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OrderGetResponse 获取随心推订单列表 API Response
type OrderGetResponse struct {
	model.BaseResponse
	Data *OrderGetResult `json:"data,omitempty"`
}

type OrderGetResult struct {
	// List
	List []Order `json:"list,omitempty"`
	// PageInfo
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
