package dmp

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AudienceDeleteRequest 删除人群 API Request
type AudienceDeleteRequest struct {
	// AdvertiserID 千川广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AudienceID 人群ID
	AudienceID uint64 `json:"audience_id,omitempty"`
}

// Encode implement PostRequest interface
func (r AudienceDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AudienceDeleteResponse 删除人群 API Response
type AudienceDeleteResponse struct {
	model.BaseResponse
	Data struct {
		// AudienceID 生成的人群ID
		AudienceID uint64 `json:"audience_id,omitempty"`
	} `json:"data,omitempty"`
}
