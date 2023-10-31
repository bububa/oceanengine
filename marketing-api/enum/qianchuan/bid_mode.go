package qianchuan

// BidMode 出价模式
type BidMode string

const (
	// BidMode_PRICING_ACTION 按优化目标出价
	BidMode_PRICING_ACTION BidMode = "PRICING_ACTION"
	// BidMode_PRICING_PLAY 按播放量出价
	BidMode_PRICING_PLAY BidMode = "PRICING_PLAY"
)
