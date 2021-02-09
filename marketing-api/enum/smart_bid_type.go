package enum

type SmartBidType string

const (
	SMART_BID_CUSTOM       SmartBidType = "SMART_BID_CUSTOM"       // 手动出价，即不使用自动出价
	SMART_BID_CONSERVATIVE SmartBidType = "SMART_BID_CONSERVATIVE" // 自动出价
)
