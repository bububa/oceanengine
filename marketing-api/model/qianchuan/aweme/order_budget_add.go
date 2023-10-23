package aweme

import "github.com/bububa/oceanengine/marketing-api/util"

// OrderBudgetAddRequest 追加随心推订单预算 API Request
type OrderBudgetAddRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OrderID 需要追加预算的订单id
	OrderID uint64 `json:"order_id,omitempty"`
	// RenewalBudget 追加的预算
	RenewalBudget float64 `json:"renewal_budget,omitempty"`
	// RenewvalDeliverySeconds 延长的投放时间
	// 短视频订单，0-7天（步进单位为1天）
	// 直播订单，0-24小时（步进单位为0.5小时）
	RenewvalDeliverySeconds float64 `json:"renewval_delivery_seconds,omitempty"`
}

// Encode implement PostRequest interface
func (r OrderBudgetAddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
