package thirdsite

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// PreviewRequest 获取第三方落地页预览地址 API Request
type PreviewRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
}

// Encode implement GetRequest interface
func (r PreviewRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("site_id", strconv.FormatUint(r.SiteID, 10))
	return values.Encode()
}

// PreviewResponse 获取第三方落地页预览地址 API Response
type PreviewResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *PreviewResponseData `json:"data,omitempty"`
}

// PreviewResponseData json返回值
type PreviewResponseData struct {
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
	// URL 站点URL
	URL string `json:"url,omitempty"`
}
