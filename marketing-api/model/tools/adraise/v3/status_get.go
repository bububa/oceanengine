package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StatusGetRequest 获取一键起量方案列表 API Request
type StatusGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionIDs 广告ID列表，目前仅支持1个
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r StatusGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("promotion_ids", string(util.JSONMarshal(r.PromotionIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// StatusGetResponse 获取广告建议起量预算 API Response
type StatusGetResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data struct {
		List []PromotionRaiseStatus `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

type PromotionRaiseStatus struct {
	// RaiseInfo 起量信息
	RaiseInfo []RaiseInfo `json:"raise_info,omitempty"`
}
