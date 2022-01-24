package enum

// PromotionWay 计划投放速度类型
type PromotionWay string

const (
	// PromotionWay_FLOW_CONTROL_MODE_FAST 尽快投放
	PromotionWay_FLOW_CONTROL_MODE_FAST PromotionWay = "FLOW_CONTROL_MODE_FAST"
	// PromotionWay_FLOW_CONTROL_MODE_SMOOTH 均匀投放
	PromotionWay_FLOW_CONTROL_MODE_SMOOTH PromotionWay = "FLOW_CONTROL_MODE_SMOOTH"
	// PromotionWay_FLOW_CONTROL_MODE_BALANCE 优先低成本，对应千川后台「严格控制成本上限」勾选项
	PromotionWay_FLOW_CONTROL_MODE_BALANCE PromotionWay = "FLOW_CONTROL_MODE_BALANCE"
)
