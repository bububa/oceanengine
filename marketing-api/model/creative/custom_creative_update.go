package creative

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CustomCreativeUpdateRequest 修改自定义创意 API Request
type CustomCreativeUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划ID，计划ID要属于广告主ID，且非删除计划，否则会报错
	AdID uint64 `json:"ad_id,omitempty"`
	// CreativeList 自定义素材信息, 最多支持10个创意。投放位置和创意类型决定素材规格。
	CreativeList []CreativeInfo `json:"creative_list,omitempty"`
	// AdData 广告计划数据
	AdData *AdData `json:"ad_data,omitempty"`
}

// Encode implement PostRequest interface
func (r CustomCreativeUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
