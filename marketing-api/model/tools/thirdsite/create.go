package thirdsite

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建第三方落地页站点 API Request
type CreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 站点名称, 长度限制，1-50 字
	Name string `json:"name,omitempty"`
	// URL 站点URL
	URL string `json:"url,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建第三方落地页站点 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CreateResponseData `json:"data,omitempty"`
}

// CreateResponseData json返回值
type CreateResponseData struct {
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
}
