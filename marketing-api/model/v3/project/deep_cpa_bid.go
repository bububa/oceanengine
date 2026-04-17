package project

import "github.com/bububa/oceanengine/marketing-api/util"

// DeepCPABidUpdateRequest 批量更新深层项目出价 API Request
type DeepCPABidUpdateRequest struct {
	// AdvertiserID 巨量广告平台广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改出价，list长度限制1～10
	Data []DeepCPABidUpdateItem `json:"data,omitempty"`
}

type DeepCPABidUpdateItem struct {
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// DeepBid 深度出价，单位“元”，精度：两位小数。
	DeepBid float64 `json:"deepbid,omitempty"`
}

// Encode implements PostRequest
func (r DeepCPABidUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
