package ad

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// KeywordCheckRequest 关键词合规校验 API Request
type KeywordCheckRequest struct {
	// AdvertiserID 千川广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Keywords 关键词列表，一次最多上传200个词
	Keywords []string `json:"keywords,omitempty"`
}

// Encode implement PostRequest interface
func (r KeywordCheckRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// KeywordCheckResponse 关键词合规校验 API Response
type KeywordCheckResponse struct {
	model.BaseResponse
	// Data 校验结果
	Data *KeywordCheckData `json:"data,omitempty"`
}

type KeywordCheckData struct {
	// FailList 违规关键词列表
	FailList []KeywordCheckResult `json:"fail_list,omitempty"`
	// SuccessList 可用关键词列表
	SuccessList []KeywordCheckResult `json:"success_list,omitempty"`
}

type KeywordCheckResult struct {
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// FailReason 违规原因
	FailReason string `json:"fail_reason,omitempty"`
}
