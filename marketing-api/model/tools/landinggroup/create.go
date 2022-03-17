package landinggroup

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建落地页组 API Request
type CreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// 落地页组名称，范围：1 <= 长度 <= 20
	GroupTitle string `json:"group_title,omitempty"`
	// GroupFlowType 流量分配方式
	GroupFlowType enum.GroupFlowType `json:"group_flow_type,omitempty"`
	// SiteIDs 橙子建站站点id列表 ：2 <= 长度 <= 10
	SiteIDs []uint64 `json:"site_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建落地页组 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *LandingGroup `json:"data,omitempty"`
}
