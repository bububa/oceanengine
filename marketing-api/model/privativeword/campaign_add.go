package privativeword

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CampaignAddRequest 批量新增组否定词 API Request
type CampaignAddRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignList 广告组否定词列表，批量操作不能超过50个项目
	CampaignList []AdWord `json:"campaign_list,omitempty"`
}

// Encode implement PostRequest interface
func (r CampaignAddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CampaignAddResponse 批量新增组否定词 API Response
type CampaignAddResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CampaignAddResponseData `json:"data,omitempty"`
}

// CampaignAddResponseData json返回值
type CampaignAddResponseData struct {
	// CampaignErrorList 请求失败的广告组id
	CampaignErrorList []string `json:"campaign_error_list,omitempty"`
	// CampaignList 添加成功的广告组否定词列表
	CampaignList []CampaignAddWordList `json:"campaign_list,omitempty"`
}

// CampaignAddWordList 添加广告组否定词计划结果
type CampaignAddWordList struct {
	// CampaignID 广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
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
