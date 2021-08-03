package enum

// LearningPhrase 学习期状态
type LearningPhrase string

const (
	// LearningPhrase_DEFAULT 默认，不在学习中
	LearningPhrase_DEFAULT LearningPhrase = "DEFAULT"
	// LearningPhrase_LEARNING 学习中
	LearningPhrase_LEARNING LearningPhrase = "LEARNING"
	// LearningPhrase_LEARNED 学习成功
	LearningPhrase_LEARNED LearningPhrase = "LEARNED"
	// LearningPhrase_LEARN_FAILED 学习失败
	LearningPhrase_LEARN_FAILED LearningPhrase = "LEARN_FAILED"
)
