package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoOriginalGetRequest 获取首发素材 API Request
type VideoOriginalGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialIDs 需要查询的素材id，不超过100个
	MaterialIDs []string `json:"material_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoOriginalGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("material_ids", string(util.JSONMarshal(r.MaterialIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoOriginalGetResponse 获取首发素材 API Response
type VideoOriginalGetResponse struct {
	model.BaseResponse
	Data struct {
		// OriginalMaterialIDs 首发素材id列表
		OriginalMaterialIDs []string `json:"original_material_ids,omitempty"`
	} `json:"data,omitempty"`
}
