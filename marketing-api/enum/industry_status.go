package enum

// IndustryStatus 行业状态
type IndustryStatus string

const (
	// IndystryStatus_NONVALID 禁投
	IndystryStatus_NONVALID IndustryStatus = "NONVALID"
	// IndustryStatus_VALID 生效
	IndustryStatus_VALID IndustryStatus = "VALID"
)
