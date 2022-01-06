package comment

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// AwemeBandedCreateRequest 添加屏蔽用户 API Request
type AwemeBandedCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// BandedType 屏蔽类型，; 允许值：CUSTOM_TYPE：自定义规则，根据昵称关键词屏蔽；AWEME_TYPE：根据抖音id屏蔽。
	BandedType string `json:"banded_type,omitempty"`
	// AwemeIDs 抖音id列表，抖音id为抖音页面展示id；当banned_type为AWEME_TYPE时必填；每次最多传50个抖音id，抖音id限制长度20，纯数字id不能以0开头；最多屏蔽5000个抖音ID。
	AwemeIDs []string `json:"aweme_ids,omitempty"`
	// NicknameKeywords 抖音昵称关键词列表，当banned_type为CUSTOM_TYPE时必填；关键词长度不大于40个字符，中文算2个字符；每次最多传50个关键词；最多屏蔽5000个关键词。
	NicknameKeywords []string `json:"nickname_keywords,omitempty"`
}

// Encode implement PostRequest interface
func (r AwemeBandedCreateRequest) Encode() []byte {
	js, _ := json.Marshal(r)
	return js
}

// AwemeBandedCreateResponse 添加屏蔽用户 API Response
type AwemeBandedCreateResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *AwemeBandedCreateResponseData `json:"data,omitempty"`
}

type AwemeBandedCreateResponseData struct {
	// Success 屏蔽成功的昵称关键词列表或抖音id列表
	Success []string `json:"success,omitempty"`
	// Fail 屏蔽失败的昵称关键词列表或抖音id列表
	Fail []FailItem `json:"fail,omitempty"`
}

type FailItem struct {
	// Message 失败原因
	Message string `json:"message,omitempty"`
	// ResultKey 屏蔽失败的昵称关键词列表或抖音id
	ResultKey string `json:"result_key,omitempty"`
}
