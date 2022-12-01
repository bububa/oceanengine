package project

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeleteRequest 批量删除项目 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectIDs 项目ID集合，list长度限制1～10
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
