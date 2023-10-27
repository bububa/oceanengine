package rta

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取可用的RTA策略 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 获取可用的RTA策略 API Response
type GetResponse struct {
	model.BaseResponse
	Data struct {
		// List 可用的RTA策略列表
		List []RtaInfo `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
