package enum

// SmartBidType 智能出价类型
type SmartBidType string

const (
	// SMART_BID_CUSTOM 手动出价，即不使用自动出价
	SMART_BID_CUSTOM SmartBidType = "SMART_BID_CUSTOM"
	// SMART_BID_CONSERVATIVE 自动出价
	SMART_BID_CONSERVATIVE SmartBidType = "SMART_BID_CONSERVATIVE"
	// SMART_BID_NO_BID
	SMART_BID_NO_BID SmartBidType = "SMART_BID_NO_BID"
)
