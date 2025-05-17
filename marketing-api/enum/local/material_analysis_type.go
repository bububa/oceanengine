package local

// MaterialAnalysisType 评估类型
type MaterialAnalysisType string

const (
	// MaterialAnalysisType_FIRST_PUBLISH 首发
	MaterialAnalysisType_FIRST_PUBLISH MaterialAnalysisType = "FIRST_PUBLISH"
	// MaterialAnalysisType_FIRST_PUBLISH_AND_HIGH_QUALITY 首发&优质
	MaterialAnalysisType_FIRST_PUBLISH_AND_HIGH_QUALITY MaterialAnalysisType = "FIRST_PUBLISH_AND_HIGH_QUALITY"
	// MaterialAnalysisType_HIGH_QUALITY 优质
	MaterialAnalysisType_HIGH_QUALITY MaterialAnalysisType = "HIGH_QUALITY"
)
