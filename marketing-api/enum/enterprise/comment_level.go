package enterprise

// CommentLevel 评论层级
type CommentLevel string

const (
	// LEVEL_ONE 一级评论
	LEVEL_ONE CommentLevel = "LEVEL_ONE"
	// LEVEL_TWO 二级评论，即回复
	LEVEL_TWO CommentLevel = "LEVEL_TWO"
)
