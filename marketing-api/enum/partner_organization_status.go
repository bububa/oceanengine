package enum

// PartnerOrganizationStatus 合作状态
type PartnerOrganizationStatus string

const (
	// PartnerOrganizationStatus_BOUND 已绑定
	PartnerOrganizationStatus_BOUND PartnerOrganizationStatus = "BOUND"
	// PartnerOrganizationStatus_BINDING 绑定中
	PartnerOrganizationStatus_BINDING PartnerOrganizationStatus = "BINDING"
	// PartnerOrganizationStatus_INVALID 失效
	PartnerOrganizationStatus_INVALID PartnerOrganizationStatus = "INVALID"
	// PartnerOrganizationStatus_DELETED 删除
	PartnerOrganizationStatus_DELETED PartnerOrganizationStatus = "DELETED"
)
