package material

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SuggestionRequest 计划下素材审核建议 API Request
type SuggestionRequest struct {
	// MaterialIDs 素材id列表
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implements GetRequest interface
func (r SuggestionRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	values.Set("material_ids", string(util.JSONMarshal(r.MaterialIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// SuggestionResponse 计划下素材审核建议 API Response
type SuggestionResponse struct {
	Data *SuggestionResult `json:"data,omitempty"`
	model.BaseResponse
}

type SuggestionResult struct {
	List []Suggestion `json:"list,omitempty"`
}

// Suggestion 素材审核建议
type Suggestion struct {
	// AuditRecords 审核列表
	AuditRecords []AuditRecord `json:"audit_records,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
}

// AuditPlatform 审核平台
type AuditPlatform string

const (
	// AuditPlatform_UNKNOWN 未知
	AuditPlatform_UNKNOWN AuditPlatform = "UNKNOWN"
	// AuditPlatform_AD 广告审核
	AuditPlatform_AD AuditPlatform = "AD"
	// AuditPlatform_CONTENT 内容审核
	AuditPlatform_CONTENT AuditPlatform = "CONTENT"
)

// AuditRecord 审核列表
type AuditRecord struct {
	// AuditPlatform 审核平台：可选值:
	// UNKNOWN：未知
	// AD：广告审核
	// CONTENT：内容审核
	AuditPlatform AuditPlatform `json:"audit_platform,omitempty"`
	// RejectReason 拒绝原因
	RejectReason []string `json:"reject_reason,omitempty"`
	// Suggestions 建议
	Suggestions []string `json:"suggestions,omitempty"`
}
