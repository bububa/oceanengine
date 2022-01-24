package ad

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// DetailGetRequest 获取计划详情（含创意信息） API Request
type DetailGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DetailGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	return values.Encode()
}

// DetailGetResponse 获取计划详情（含创意信息） API Response
type DetailGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *Ad `json:"data,omitempty"`
}
