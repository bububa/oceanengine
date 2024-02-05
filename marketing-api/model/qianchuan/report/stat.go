package report

// Stat 统计数据
type Stat struct {
	Dimension
	Metrics
	// Fields 指标数据
	Fields *Metrics `json:"metrics,omitempty"`
	// AnalysisType 素材建议
	AnalysisType []string `json:"analysis_type,omitempty"`
}
