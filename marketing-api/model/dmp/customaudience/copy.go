package customaudience

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CopyRequest 推送dmp人群包到云图账户 API Request
type CopyRequest struct {
	// AdvertiserID 人群包所属广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CustomAudienceID 人群包ID
	CustomAudienceID uint64 `json:"custom_audience_id,omitempty"`
	// ToAdvertiserID 推送广告主ID（云图虚拟广告主ID，即virtual_adv_id），可以通过【获取广告账户关联云图账户信息】接口获取
	ToAdvertiserID []uint64 `json:"to_advertiser_id,omitempty"`
}

// Encode implement PostRequest interface
func (r CopyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
