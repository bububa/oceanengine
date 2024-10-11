package enum

// LearningPhase 学习期状态
type LearningPhase string

const (
	// LearningPhase_DEFAULT 默认，不在学习中
	LearningPhase_DEFAULT LearningPhase = "DEFAULT"
	// LearningPhase_LEARNING 学习中
	LearningPhase_LEARNING LearningPhase = "LEARNING"
	// LearningPhase_LEARNED 学习成功
	LearningPhase_LEARNED LearningPhase = "LEARNED"
	// LearningPhase_LEARN_FAILED 学习失败
	LearningPhase_LEARN_FAILED LearningPhase = "LEARN_FAILED"
	// LearningPhase_ALL 不限
	LearningPhase_ALL LearningPhase = "ALL"
)
