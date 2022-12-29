package comment

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TermsBannedAddRequest 添加屏蔽词 API Request
type TermsBannedAddRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 为绑定的抖音号添加屏蔽词, 只允许传入1个，通过【获取绑定的抖音号】 接口获取，当is_apply_to_adv不传或为false时，aweme_id生效
	AwemeID string `json:"aweme_id,omitempty"`
	// IsApplyToAdv 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
	IsApplyToAdv bool `json:"is_apply_to_adv,omitempty"`
	// Terms 屏蔽词列表，若添加的屏蔽词已存在，不会再次新增，同一个屏蔽词只会在屏蔽词中记录一次，一次最多添加100个，单个屏蔽词长度范围为0-20字
	Terms []string `json:"terms,omitempty"`
}

// Encode implement PostRequest interface
func (r TermsBannedAddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
