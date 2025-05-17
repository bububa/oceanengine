package material

import "github.com/bububa/oceanengine/marketing-api/util"

// AdMaterialDeleteRequest 删除广告计划下素材 API Request
type AdMaterialDeleteRequest struct {
	// MaterialIDs 待删除素材ID
	// 注意：最大支持100个素材
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implement PostRequest interface
func (r AdMaterialDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
