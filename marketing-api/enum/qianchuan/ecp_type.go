package qianchuan

// EcpType 账户类型
type EcpType string

const (
	// EcpType_SHOP 商家
	EcpType_SHOP EcpType = "SHOP"
	// EcpType_SHOP_STAR 商家达人
	EcpType_SHOP_STAR EcpType = "SHOP_STAR"
	// EcpType_COMMON_STAR 普通达人
	EcpType_COMMON_STAR EcpType = "COMMON_STAR"
	// EcpType_AGENT 百应机构
	EcpType_AGENT EcpType = "AGENT"
)
