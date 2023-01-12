package enum

// DeliveryMode 投放模式，允许值：MANUAL手动投放 (默认值），PROCEDURAL自动投放
type DeliveryMode string

const (
	DeliveryMode_MANUAL     DeliveryMode = "MANUAL"
	DeliveryMode_PROCEDURAL DeliveryMode = "PROCEDURAL"
)
