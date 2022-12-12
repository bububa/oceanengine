package auth

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetAuthStatusRequest 查询鉴权开启状态 API Request
type GetAuthStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r GetAuthStatusRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetAuthStatusResponse 查询鉴权开启状态
type GetAuthStatusResponse struct {
	model.BaseResponse
	Data struct {
		// Enabled 鉴权是否开启
		Enabled bool `json:"enabled,omitempty"`
	} `json:"data,omitempty"`
}
