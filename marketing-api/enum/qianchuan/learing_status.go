package qianchuan

// LearningStatus 学习期状态
type LearningStatus string

const (
	// LearningStatus_LEARNING 学习期
	LearningStatus_LEARNING LearningStatus = "LEARNING"
	// LearningStatus_LEARNED 学习期结束
	LearningStatus_LEARNED LearningStatus = "LEARNING"
	// LearningStatus_LEARN_FAILED 学习期失败
	LearningStatus_LEARN_FAILED LearningStatus = "LEARN_FAILED"
	// LearningStatus_DEFAULT 无学习期状态
	LearningStatus_DEFAULT LearningStatus = "DEFAULT"
)
