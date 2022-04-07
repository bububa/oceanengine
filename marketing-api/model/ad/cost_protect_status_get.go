package ad

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// CostProtectStatusGetRequest 批量获取计划成本保障状态 API Request
type CostProtectStatusGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 广告计划id，每次最多传入50个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r CostProtectStatusGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.AdIDs)
	values.Set("ad_ids", string(ids))
	return values.Encode()
}

// CostProtectStatusGetResponse 批量获取计划成本保障状态 API Response
type CostProtectStatusGetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data struct {
		// List .
		List []CostProtectStatus `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// CostProtectStatus guan广告成本保护状态
type CostProtectStatus struct {
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Status 计划成本保障状态
	Status string `json:"status,omitempty"`
}
