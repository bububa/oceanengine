package enum

// StampType 是否电子签章
type StampType string

const (
	// StampType_ONLINE 电子签章
	StampType_ONLINE StampType = "ONLINE"
	// StampType_DEFAULT 线下签章
	StampType_DEFAULT StampType = "DEFAULT"
)
