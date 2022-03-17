package comment

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TermsBandedDeleteRequest 删除屏蔽词 API Request
type TermsBandedDeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Terms 待删除的屏蔽词列表; 如果删除的屏蔽词不存在也会显示成功; 一次最多操作500个词，单个屏蔽词长度范围为0-100
	Terms []string `json:"terms,omitempty"`
}

// Encode implement PostRequest interface
func (r TermsBandedDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
