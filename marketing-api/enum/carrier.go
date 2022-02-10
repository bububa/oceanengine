package enum

// Carrier 电信运营商
type Carrier string

const (
	// Carrier_MOBILE 移动
	Carrier_MOBILE Carrier = "MOBILE"
	// Carrier_UNICOM 联通
	Carrier_UNICOM Carrier = "UNICOM"
	// Carrier_TELCOM 电信
	Carrier_TELCOM Carrier = "TELCOM"
)
