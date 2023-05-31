package qianchuan

// ExplainStatus 商品状态
type ExplainStatus string

const (
	// ExplainStatus_ALL 全部
	ExplainStatus_ALL ExplainStatus = "ALL"
	// ExplainStatus_BEINGEXPLAIN 讲解中
	ExplainStatus_BEINGEXPLAIN ExplainStatus = "BEINGEXPLAIN"
	// ExplainStatus_HASEXPLAIN 已讲解
	ExplainStatus_HASEXPLAIN ExplainStatus = "HASEXPLAIN"
	// ExplainStatus_UNEXPLAIN 未讲解
	ExplainStatus_UNEXPLAIN ExplainStatus = "UNEXPLAIN"
)
