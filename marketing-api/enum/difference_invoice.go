package enum

// DifferenceInvoice 是否差额开
type DifferenceInvoice string

const (
	// DifferenceInvoice_FULL 全额开票
	DifferenceInvoice_FULL DifferenceInvoice = "FULL"
	// DifferenceInvoice_DIFF 差额开票
	DifferenceInvoice_DIFF DifferenceInvoice = "DIFF"
)
