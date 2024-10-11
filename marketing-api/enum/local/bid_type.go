package local

// BidType 出价方式
type BidType string

const (
	// BidType_ALL 不限
	BidType_ALL BidType = "ALL"
	// BidType_MANUAL 手动出价
	BidType_MANUAL BidType = "MANUAL"
	// BidType_SMART 智能出价
	BidType_SMART BidType = "SMART"
)
