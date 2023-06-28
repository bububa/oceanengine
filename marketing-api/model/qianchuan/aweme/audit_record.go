package aweme

// AuditRecord 审核详细内容
type AuditRecord struct {
	// DetailDescList 建议内容列表
	DetailDescList []string `json:"detail_desc_list,omitempty"`
	// StatusAttachDesp 订单状态附加描述
	StatusAttachDesp string `json:"status_attach_desp,omitempty"`
}
