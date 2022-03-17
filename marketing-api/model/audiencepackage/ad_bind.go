package audiencepackage

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdBindRequest 计划绑定定向包 API Request
type AdBindRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AudiencePackageID 删除的定向包id
	AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	// AdIDs 绑定的计划id列表，上限1000个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r AdBindRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdBindResponse 绑定定向包 API Response
type AdBindResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	} `json:"data,omitempty"`
}
