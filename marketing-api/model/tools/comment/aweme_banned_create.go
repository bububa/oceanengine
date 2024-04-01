package comment

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeBannedCreateRequest 添加屏蔽用户 API Request
type AwemeBannedCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 为绑定的抖音号添加屏蔽词, 只允许传入1个，通过【获取绑定的抖音号】 接口获取，当is_apply_to_adv不传或为false时，aweme_id生效
	AwemeID string `json:"aweme_id,omitempty"`
	// IsApplyToAdv 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
	IsApplyToAdv bool `json:"is_apply_to_adv,omitempty"`
	// BannedType 屏蔽类型，; 允许值：CUSTOM_TYPE：自定义规则，根据昵称关键词屏蔽；AWEME_TYPE：根据抖音id屏蔽。
	BannedType string `json:"banned_type,omitempty"`
	// AwemeUserIDs 抖音用户id列表，传入屏蔽用户的抖音号；当banned_type为AWEME_TYPE时必填；每次最多传50个抖音id，纯数字id不能以0开头；一个抖音号最多屏蔽5000个抖音ID。
	AwemeUserIDs []string `json:"aweme_user_ids,omitempty"`
	// NicknameKeywords 抖音昵称关键词列表，当banned_type为CUSTOM_TYPE时必填；关键词长度不大于40个字符，中文算2个字符；每次最多传50个关键词；最多屏蔽5000个关键词。
	NicknameKeywords []string `json:"nickname_keywords,omitempty"`
}

// Encode implement PostRequest interface
func (r AwemeBannedCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AwemeBannedCreateResponse 添加屏蔽用户 API Response
type AwemeBannedCreateResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *AwemeBannedCreateResponseData `json:"data,omitempty"`
}

// AwemeBannedCreateResponseData json返回值
type AwemeBannedCreateResponseData struct {
	// Success 屏蔽成功的昵称关键词列表或抖音id列表
	Success []string `json:"success,omitempty"`
	// Fail 屏蔽失败的昵称关键词列表或抖音id列表
	Fail []FailItem `json:"fail,omitempty"`
}

// FailItem 失败信息
type FailItem struct {
	// Message 失败原因
	Message string `json:"message,omitempty"`
	// ResultKey 屏蔽失败的昵称关键词列表或抖音id
	ResultKey string `json:"result_key,omitempty"`
}
