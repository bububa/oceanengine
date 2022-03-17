package thirdsite

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 修改第三方落地页站点 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
	// Name 站点名称, 长度限制，1-50 字
	Name string `json:"name,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 修改第三方落地页站点 API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json返回值
type UpdateResponseData struct {
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
}
