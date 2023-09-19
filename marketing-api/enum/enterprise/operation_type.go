package enterprise

// OperationType 企业号推广操作类型
type OperationType string

const (
	// OperationType_LIVE 直播开始
	OperationType_LIVE OperationType = "LIVE"
	// OperationType_VIDEO 发布视频
	OperationType_VIDEO OperationType = "VIDEO"
	// OperationType_DOU 开启DOU+推广
	OperationType_DOU OperationType = "DOU"
	// OperationType_AD 开启广告投放
	OperationType_AD OperationType = "AD"
)
