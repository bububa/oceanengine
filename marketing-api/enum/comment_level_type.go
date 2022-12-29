package enum

// CommentLevelType 评论等级
type CommentLevelType string

const (
	// CommentLevelType_LEVEL_ALL 所有等级
	CommentLevelType_LEVEL_ALL CommentLevelType = "LEVEL_ALL"
	// CommentLevelType_LEVEL_ONE 一级评论
	CommentLevelType_LEVEL_ONE CommentLevelType = "LEVEL_ONE"
	// CommentLevelType_LEVEL_TWO 二级评论
	CommentLevelType_LEVEL_TWO CommentLevelType = "LEVEL_TWO"
)
