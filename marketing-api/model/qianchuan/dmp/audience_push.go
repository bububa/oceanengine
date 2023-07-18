package dmp

import (
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AudiencePushRequest 推送人群 API Request
type AudiencePushRequest struct {
	// AdvertiserID 千川广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AudienceID 人群ID
	AudienceID uint64 `json:"audience_id,omitempty"`
	// AccountType 推送账户，允许值:
	// LOCAL：本账户
	// NO_LOCAL：其他账户
	AccountType qianchuan.AudienceAccountType `json:"account_type,omitempty"`
	// IsForBrand 人群应用，允许值：
	// 0：仅效果广告（默认）
	// 1：效果广告+品牌广告
	IsForBrand int `json:"is_for_brand,omitempty"`
	// AdvIDs 需要推送到的广告账户id列表
	// 注：仅在“推送账户”为“其他账户”时，条件必填
	AdvIDs []uint64 `json:"adv_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r AudiencePushRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AudiencePushResponse 推送人群 API Response
type AudiencePushResponse struct {
	model.BaseResponse
	Data struct {
		// AudienceID 生成的人群ID
		AudienceID uint64 `json:"audience_id,omitempty"`
	} `json:"data,omitempty"`
}
