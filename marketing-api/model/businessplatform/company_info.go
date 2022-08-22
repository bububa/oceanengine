package businessplatform

import "github.com/bububa/oceanengine/marketing-api/enum"

// CompanyInfo 主体信息
type CompanyInfo struct {
	// CompanyID 公司主体id
	CompanyID uint64 `json:"company_id,omitempty"`
	// CompanyName 公司主体名
	CompanyName string `json:"company_name,omitempty"`
	// Type 主体关系，可选值:
	// BP_OTHER:
	// BP_OWN:
	Type enum.CompanyType `json:"type,omitempty"`
	// Status 公司主体状态，可选值:
	// EXPIRED:已过期
	// FAILED: 审核拒绝
	// NOT_STARTED:待提交
	// PROCESSING: 审核中
	// SUCCESS: 审核通过
	// WAITING: 待审核
	Status enum.CompanyStatus `json:"status,omitempty"`
}
