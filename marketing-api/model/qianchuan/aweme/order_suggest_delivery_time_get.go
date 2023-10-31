package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrderSuggestDeliveryTimeGetRequest 获取建议延长时长 API Request
type OrderSuggestDeliveryTimeGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OrderID 需要追加预算的订单id
	OrderID uint64 `json:"order_id,omitempty"`
	// AddAmount 追加的预算
	AddAmount float64 `json:"add_amount,omitempty"`
}

// Encode implement GetRequest interface
func (r OrderSuggestDeliveryTimeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("order_id", strconv.FormatUint(r.OrderID, 10))
	values.Set("add_amount", strconv.FormatFloat(r.AddAmount, 'f', -1, 64))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OrderSuggestDeliveryTimeGetResponse 获取建议延长时长 API Response
type OrderSuggestDeliveryTimeGetResponse struct {
	model.BaseResponse
	Data struct {
		// SuggestDeliveryTime 建议追加投放时长
		// 短视频：xx天
		// 直播：xx小时
		SuggestDeliveryTime float64 `json:"suggest_delivery_time,omitempty"`
	} `json:"data,omitempty"`
}
