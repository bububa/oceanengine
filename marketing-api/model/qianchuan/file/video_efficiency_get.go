package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoEfficiencyGetRequest 获取低效素材 API Request
type VideoEffeciencyGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialIDs 需要查询的素材id，不超过100个
	MaterialIDs []string `json:"material_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoEffeciencyGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("material_ids", string(util.JSONMarshal(r.MaterialIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoEffeciencyGetResponse 获取低效素材 API Response
type VideoEffeciencyGetResponse struct {
	model.BaseResponse
	Data struct {
		// InEffecientMaterialIDs 低效素材id列表
		InEffecientMaterialIDs []string `json:"in_effecient_material_ids,omitempty"`
	} `json:"data,omitempty"`
}
