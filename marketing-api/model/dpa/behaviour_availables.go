package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BehaviourAvailablesRequest 获取DPA可用行为 API Request
type BehaviourAvailablesRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库id
	PlatformID uint64 `json:"platform_id,omitempty"`
}

// Encode implement GetRequest interface
func (r BehaviourAvailablesRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("platform_id", strconv.FormatUint(r.PlatformID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// BehaviourAvailablesResponse 获取DPA可用行为 API Response
type BehaviourAvailablesResponse struct {
	model.BaseResponse
	Data struct {
		// List DPA行为列表
		List []Behaviour `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
