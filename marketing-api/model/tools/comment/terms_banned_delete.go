package comment

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TermsBannedDeleteRequest 删除屏蔽词 API Request
type TermsBannedDeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 为绑定的抖音号添加屏蔽词, 只允许传入1个，通过【获取绑定的抖音号】 接口获取，当is_apply_to_adv不传或为false时，aweme_id生效
	AwemeID string `json:"aweme_id,omitempty"`
	// IsApplyToAdv 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
	IsApplyToAdv bool `json:"is_apply_to_adv,omitempty"`
	// Terms 屏蔽词列表，一次最多操作100个词，屏蔽词最大20字
	Terms []string `json:"terms,omitempty"`
}

// Encode implement PostRequest interface
func (r TermsBannedDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
