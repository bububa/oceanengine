package enum

// BidType 竞价策略
type BidType string

const (
	// BidType_CUSTOM 稳定成本
	BidType_CUSTOM BidType = "CUSTOM"
	// BidType_NO_BID 最大转化投放
	BidType_NO_BID BidType = "BidType_NO_BID"
	// BidType_UPPER_CONTROL 控制成本上限
	BidType_UPPER_CONTROL BidType = "UPPER_CONTROL"
	// BidType_CONSERVATIVE 放量投放
	BidType_CONSERVATIVE BidType = "CONSERVATIVE"
	// BidType_EXPLORE_UPGRADE 稳定成本-升级版
	BidType_EXPLORE_UPGRADE BidType = "EXPLORE_UPGRADE"
)
