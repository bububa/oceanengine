package site

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// ReadRequest 获取橙子建站站点详细信息 API Request
type ReadRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteID 橙子建站站点id列表 ：1 <= 长度 <= 20
	SiteID uint64 `json:"site_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ReadRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("site_id", strconv.FormatUint(r.SiteID, 10))
	return values.Encode()
}

// ReadResponse 获取橙子建站站点详细信息 API Response
type ReadResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *SiteDetail `json:"data,omitempty"`
}
