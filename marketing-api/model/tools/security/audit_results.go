package security

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AuditResultsRequest 广告素材预审结果 API Request
type AuditResultsRequest struct {
	// AccountID 业务广告账户ID
	AccountID uint64 `json:"account_id,omitempty"`
	// ObjectID 送审对象ID
	ObjectID uint64 `json:"object_id,omitempty"`
}

// Encode implements GetRequest interface
func (r AuditResultsRequest) Encode() string {
	values := util.GetUrlValues()
	defer util.PutUrlValues(values)
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("object_id", strconv.FormatUint(r.ObjectID, 10))
	return values.Encode()
}

type AuditResultsResponse struct {
	model.BaseResponse
	Data *AuditResult `json:"data,omitempty"`
}

type AuditResult struct {
	// 审核状态，枚举值：
	// APPROVE 审核通过
	// REJECT 审核拒绝
	// AUDITING 审核中
	Status string `json:"status,omitempty"`
	// ReasonText 审核建议文案
	ReasonText string `json:"reason_text,omitempty"`
}
