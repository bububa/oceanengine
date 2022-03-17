package creative

import (
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateStatusRequest 更新创意状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreativeIDs 需要更新的广告创意id，一次最多更新10个创意
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// OptStatus 批量更新的广告计划状态
	OptStatus qianchuan.CreativeOptStatus `json:"opt_status,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
