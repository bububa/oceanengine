package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StatusUpdateRequest 更新项目状态 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64             `json:"advertiser_id,omitempty"`
	Data         []StatusUpdateData `json:"data,omitempty"`
}

type StatusUpdateData struct {
	// ProjectID 广告项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// OptStatus  操作ENABLE启 用广告、DISABLE 暂停广告
	OptStatus enum.OptStatus `json:"opt_status,omitempty"`
}

// Encode implement PostRequest interface
func (r StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
