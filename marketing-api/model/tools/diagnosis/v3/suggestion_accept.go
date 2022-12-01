package v3

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/tools/diagnosis"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SuggestionAcceptRequest 采纳计划诊断建议 API Request
type SuggestionAcceptRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionID 广告计划ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// DiagnosisID 诊断id
	DiagnosisID string `json:"diagnosis_id,omitempty"`
	// Tools 工具列表
	Tools []diagnosis.Tool `json:"tools,omitempty"`
}

// Encode implement PostRequest interface
func (r SuggestionAcceptRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SuggestionAcceptResponse 采纳计划诊断建议 API Response
type SuggestionAcceptResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *SuggestionAcceptResponseData `json:"data,omitempty"`
}

// SuggestionAcceptResponseData json返回值
type SuggestionAcceptResponseData struct {
	// SuggestionAccept 采纳成功的工具列表
	SuggestionAccept []string `json:"suggestion_accept,omitempty"`
	// SuggestionAcceptFailed 采纳失败的工具列表
	SuggestionAcceptFailed []SuggestionAcceptFailed `json:"suggestion_accept_failed,omitempty"`
}

// SuggestionAcceptFailed 采纳失败
type SuggestionAcceptFailed struct {
	// ToolType 采纳失败的工具
	ToolType string `json:"tool_type,omitempty"`
	// FailReason 采纳失败的原因
	FailReason string `json:"fail_reason,omitempty"`
}
