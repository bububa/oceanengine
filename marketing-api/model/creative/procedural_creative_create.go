package creative

import "github.com/bububa/oceanengine/marketing-api/util"

// ProceduralCreativeCreateRequest 创建程序化创意（营销链路） API Request
type ProceduralCreativeCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划ID，计划ID要属于广告主ID，且非删除计划，否则会报错
	AdID uint64 `json:"ad_id,omitempty"`
	// Creative 程序化素材信息，投放位置和创意类型决定素材规格。
	Creative *CreativeInfo `json:"creative,omitempty"`
	// AdData 广告计划数据
	AdData *AdData `json:"ad_data,omitempty"`
}

// Encode implement PostRequest interface
func (r ProceduralCreativeCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
