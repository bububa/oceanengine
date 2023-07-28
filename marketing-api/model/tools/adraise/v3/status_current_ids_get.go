package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StatusCurrentIDsGetRequest 获取广告起量状态 API Request
type StatusCurrentIDsGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionIDs 广告ID列表，目前仅支持1个
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r StatusCurrentIDsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("promotion_ids", string(util.JSONMarshal(r.PromotionIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// StatusCurrentIDsGetResponse 获取广告起量状态 API Response
type StatusCurrentIDsGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *StatusCurrentIDsGetResult `json:"data,omitempty"`
}

type StatusCurrentIDsGetResult struct {
	Data      []RaiseStatus `json:"data,omitempty"`
	ErrorList []RaiseError  `json:"error_list,omitempty"`
}

type RaiseStatus struct {
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// Status 当前一键起量状态 可选值:
	// DISABLE_RAISE
	// ENABLE_PRERAISE
	// ENABLE_RAISE
	// FINISH_RAISE
	// HAS_PRERAISE
	// RAISING
	Status enum.AdRaiseStatus `json:"status,omitempty"`
	// RaiseInfo 当前的起量信息， 当status=RAISING时返回该字段
	RaiseInfo *RaiseInfo `json:"raise_info,omitempty"`
}

type RaiseError struct {
	// PromotionID 获取一键起量状态失败的广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// ErrorReason 获取一键起量状态失败的错误原因
	ErrorReason string `json:"error_reason,omitempty"`
}
