package aweme

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrderTerminateRequest 终止随心推订单 API Request
type OrderTerminateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OrderIDs 需要终止的订单id
	// 注意：目前仅支持一个
	OrderIDs []uint64 `json:"order_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r OrderTerminateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// OrderTerminateResponse 终止随心推订单 API Response
type OrderTerminateResponse struct {
	model.BaseResponse
	Data *OrderTerminateResult `json:"data,omitempty"`
}

type OrderTerminateResult struct {
	// OrderIDs 终止成功的订单ID
	OrderIDs []uint64 `json:"order_ids,omitempty"`
	// TerminateQuotaInfo 终止订单配额信息
	TerminateQuotaInfo *TerminateQuotaInfo `json:"terminate_quota_info,omitempty"`
}
