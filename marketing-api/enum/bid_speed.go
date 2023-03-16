package enum

// BidSpeed 投放速度
type BidSpeed string

const (
	// BidSpeed_BALANCE 匀速
	BidSpeed_BALANCE BidSpeed = "BALANCE"
	// BidSpeed_FAST 加速
	BidSpeed_FAST BidSpeed = "FAST"
)
