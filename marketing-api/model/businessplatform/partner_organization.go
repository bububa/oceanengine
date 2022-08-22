package businessplatform

import "github.com/bububa/oceanengine/marketing-api/enum"

// PartnerOrganization 发起合作组织请求的组织
type PartnerOrganization struct {
	// ID 发起合作组织请求的组织id
	ID uint64 `json:"partner_organization_id,omitempty"`
	// Status 合作状态
	// 枚举值：BOUND（已绑定）、BINDING（绑定中）、INVALID（失效）、DELETED（删除）
	Status enum.PartnerOrganizationStatus `json:"status,omitempty"`
	// Remark 备注，合作组织备注信息
	Remark string `json:"remark,omitempty"`
}
