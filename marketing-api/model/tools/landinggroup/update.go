package landinggroup

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 更新落地页组信息 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// GroupID 落地页组 ID
	GroupID uint64 `json:"group_id,omitempty"`
	// GroupTitle 落地页组名称
	GroupTitle string `json:"group_title,omitempty"`
	// AppendSites 新加站点列表
	AppendSites []uint64 `json:"append_sites,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新落地页组信息 API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *LandingGroup `json:"data,omitempty"`
}
