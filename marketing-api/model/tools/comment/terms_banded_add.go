package comment

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TermsBandedAddRequest 添加屏蔽词 API Request
type TermsBandedAddRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Terms 待添加的屏蔽词列表; 若添加的屏蔽词已存在，不会再次新增，同一个屏蔽词只会在屏蔽词中记录一次;一次最多添加500个，单个屏蔽词长度范围为0-100
	Terms []string `json:"terms,omitempty"`
}

// Encode implement PostRequest interface
func (r TermsBandedAddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
