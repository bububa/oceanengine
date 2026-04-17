package project

import "github.com/bububa/oceanengine/marketing-api/util"

// CPABidUpdateRequest 批量更新项目出价 API Request
type CPABidUpdateRequest struct {
	// AdvertiserID 巨量广告平台广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改出价，list长度限制1～10
	Data []CPABidUpdateItem `json:"data,omitempty"`
}

type CPABidUpdateItem struct {
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// CPABid 出价，单位“元”，精度：两位小数
	CPABid float64 `json:"cpa_bid,omitempty"`
}

// Encode implements PostRequest
func (r CPABidUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
