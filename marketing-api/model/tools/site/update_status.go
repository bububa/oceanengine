package site

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateStatusRequest 修改橙子建站站点状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SiteIDs 橙子建站站点id列表 ：1 <= 长度 <= 20; 可通过【获取橙子建站站点列表】获取应答字段list
	SiteID []string `json:"site_ids,omitempty"`
	// Status 站点操作状态
	Status enum.SiteOptStatus `json:"status,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateStatusResponse 更改橙子建站站点状态 API Response
type UpdateStatusResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *UpdateStatusResponseData `json:"data,omitempty"`
}

// UpdateStatusResponseData json 返回值
type UpdateStatusResponseData struct {
	// Success 更新成功的site_id 列表。如果全部更新失败，len(list) = 0
	Success []string `json:"success,omitempty"`
	// Fail 更新失败的site_id的list及失败的原因。如果全部更新成功，len(list) = 0
	Fail []FailMessage `json:"fail,omitempty"`
}
