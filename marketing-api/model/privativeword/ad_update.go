package privativeword

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdUpdateRequest 设置计划否定词 API Request
type AdUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	AdWord
}

// Encode implement PostRequest interface
func (r AdUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdUpdateResponse 设置计划否定词 API Response
type AdUpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdUpdateResponseData `json:"data,omitempty"`
}

// AdUpdateResponseData json返回值
type AdUpdateResponseData struct {
	// AdID 广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// CampaignID 广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
}
