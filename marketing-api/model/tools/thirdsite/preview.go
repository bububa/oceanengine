package thirdsite

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
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
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("site_id", strconv.FormatUint(r.SiteID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
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
