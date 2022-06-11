package file

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// VideoEffeciencyGetRequest 获取低效素材 API Request
type VideoEffeciencyGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialIDs 素材列表，单次最多可查询100个
	MaterialIDs []string `json:"material_ids,omitempty"`
}

// Encode implement GetRequest interfacd
func (r VideoEffeciencyGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.MaterialIDs)
	values.Set("material_ids", string(ids))
	return values.Encode()
}

// VideoEffeciencyGetResponse 获取低效素材 API Response
type VideoEffeciencyGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// IDs <BS>低效素材id列表
		IDs []string `json:"inefficient_material_ids,omitempty"`
	} `json:"data,omitempty"`
}
