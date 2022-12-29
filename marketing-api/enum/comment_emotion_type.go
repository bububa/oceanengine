package enum

// CommentEmotionType 评论情感
type CommentEmotionType string

const (
	// CommentEmotionType_NEGATIVE 负向评论
	CommentEmotionType_NEGATIVE CommentEmotionType = "NEGATIVE"
	// CommentEmotionType_NEUTRAL 中性评论
	CommentEmotionType_NEUTRAL CommentEmotionType = "NEUTRAL"
	// CommentEmotionType_POSITIVE: 正向评论
	CommentEmotionType_POSITIVE CommentEmotionType = "POSITIVE"
)
