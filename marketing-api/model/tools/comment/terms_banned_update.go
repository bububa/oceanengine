package comment

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TermsBannedUpdateRequest 更新屏蔽词 API Request
type TermsBannedUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 为绑定的抖音号添加屏蔽词, 只允许传入1个，通过【获取绑定的抖音号】 接口获取，当is_apply_to_adv不传或为false时，aweme_id生效
	AwemeID string `json:"aweme_id,omitempty"`
	// IsApplyToAdv 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
	IsApplyToAdv bool `json:"is_apply_to_adv,omitempty"`
	// OriginTerms 待更新的屏蔽词; 屏蔽词长度范围为0-100
	OriginTerms []string `json:"origin_terms,omitempty"`
	// 更新后的屏蔽词; 如果new_terms已存在，则等同于删除origin_terms；如果origin_terms不存在，则等同于新增new_terms; 屏蔽词长度范围为0-100
	NewTerms []string `json:"new_terms,omitempty"`
}

// Encode implement PostRequest interface
func (r TermsBannedUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
