package adraise

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// StatusRequest 获取当前起量状态 API Request
type StatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 广告计划id 列表，最多1000个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r StatusRequest) Encode() string {
	ret := &url.Values{}
	ret.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.AdIDs)
	ret.Set("ad_ids", string(ids))
	return ret.Encode()
}

// StatusResponse 获取当前起量状态 API Response
type StatusResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data struct {
		// Status 一键起量状态
		Status string `json:"status,omitempty"`
	} `json:"data,omitempty"`
}
