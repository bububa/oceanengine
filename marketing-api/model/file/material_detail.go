package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialDetailRequest 查询素材标签信息 API Request
type MaterialDetailRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialIDs 素材 id 列表，最多50
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r MaterialDetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("material_ids", string(util.JSONMarshal(r.MaterialIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MaterialDetailResponse 查询素材标签信息 API Response
type MaterialDetailResponse struct {
	model.BaseResponse
	Data struct {
		// Materials 素材列表
		Materials []Material `json:"materials,omitempty"`
	} `json:"data,omitempty"`
}
