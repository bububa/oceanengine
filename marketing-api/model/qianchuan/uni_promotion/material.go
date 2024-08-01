package unipromotion

// Material 素材
type Material struct {
	// MaterialInfo 素材信息
	MaterialInfo *MaterialInfo `json:"material_info,omitempty"`
}

// MaterialInfo 素材信息
type MaterialInfo struct {
	// MaterialStatus 素材状态，可选值:
	// DELIVERY_OK 投放中
	// DELETED 已删除
	MaterialStatus string `json:"material_status,omitempty"`
	// AuditStatus 审核状态，可选值:
	// PASS 审核通过
	// REJECT 审核拒绝
	// IN_PROGRESS 审核中
	AuditStatus string `json:"audit_status,omitempty"`
	// IsDeleted 是否已删除
	IsDeleted bool `json:"is_deleted,omitempty"`
}
