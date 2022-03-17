package privativeword

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdAddRequest 批量新增计划否定词 API Request
type AdAddRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdList 广告计划否定词列表，批量操作不能超过50个项目
	AdList []AdWord `json:"ad_list,omitempty"`
}

// Encode implement PostRequest interface
func (r AdAddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdAddResponse 批量新增计划否定词 API Response
type AdAddResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdAddResponseData `json:"data,omitempty"`
}

// AdAddResponseData json返回值
type AdAddResponseData struct {
	// AdErrorList 请求失败的广告计划id
	AdErrorList []string `json:"ad_error_list,omitempty"`
	// AdList 添加成功的广告计划否定词列表
	AdList []AdAddWordList `json:"ad_list,omitempty"`
}

// AdAddWordList 添加广告计划否定词计划结果
type AdAddWordList struct {
	// AdID 广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Status 操作执行状态，只要有操作失败的词，就返回 failure，全部成功返回 success
	Status string `json:"status,omitempty"`
	// DuplicateWords 有重复的否定词词导致操作失败的列表
	DuplicateWords []AdWord `json:"duplicate_words,omitempty"`
	// ExceedLengthWords 否定词词超过200个导致操作失败的列表
	ExceedLengthWords []AdWord `json:"exceed_length_words,omitempty"`
	// ExceedLimitWords 单个否定词的长度超过限制导致操作失败的列表
	ExceedLimitWords []AdWord `json:"exceed_limit_words,omitempty"`
	// HasEmojiWords 有 emoji 导致操作失败的列表
	HasEmojiWords []AdWord `json:"has_emoji_words,omitempty"`
	// SuccessWords 操作成功的否定词列表
	SuccessWords []AdWord `json:"success_words,omitempty"`
}
