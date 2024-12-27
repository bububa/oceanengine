package enum

// InvoiceStatus 开票状态
type InvoiceStatus string

const (
	// InvoiceStatus_NO 未开票
	InvoiceStatus_NO InvoiceStatus = "NO"
	// InvoiceStatus_PART 部分开票
	InvoiceStatus_PART InvoiceStatus = "PART"
	// InvoiceStatus_COMPLETE 已开票
	InvoiceStatus_COMPLETE InvoiceStatus = "COMPLETE"
)
