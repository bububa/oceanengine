package enum

// BidType 竞价策略
type BidType string

const (
	// BidType_CUSTOM 稳定成本
	BidType_CUSTOM BidType = "CUSTOM"
	// BidType_NO_BID 最大转化投放
	BidType_NO_BID BidType = "BidType_NO_BID"
)
