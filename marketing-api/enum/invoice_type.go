package enum

// InvoiceType 开票类型
type InvoiceType string

const (
	// InvoiceType_GENERAL 增值税普通发票
	InvoiceType_GENERAL InvoiceType = "GENERAL"
	// InvoiceType_SPECIAL 增值税专用发票
	InvoiceType_SPECIAL InvoiceType = "SPECIAL"
	// InvoiceType_E_GENERAL 增值税电子普通发票
	InvoiceType_E_GENERAL InvoiceType = "E_GENERAL"
	// InvoiceType_E_SPECIAL 增值税电子专用发票
	InvoiceType_E_SPECIAL InvoiceType = "E_SPECIAL"
	// InvoiceType_ALL_E_GENERAL 电子发票（普通发票）
	InvoiceType_ALL_E_GENERAL InvoiceType = "ALL_E_GENERAL"
	// InvoiceType_ALL_E_SPECIAL 电子发票（增值税专用发票）
	InvoiceType_ALL_E_SPECIAL InvoiceType = "ALL_E_SPECIAL"
)
