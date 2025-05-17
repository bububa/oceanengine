package unipromotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailRequest 获取全域推广计划详情 API Request
type DetailRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 全域推广计划id
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DetailResponse 获取全域推广计划详情 API Response
type DetailResponse struct {
	Data *AdDetail `json:"data,omitempty"`
	model.BaseResponse
}
