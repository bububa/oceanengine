package enum

// CommentReplyStatus 评论回复状态
type CommentReplyStatus string

const (
	// CommentReplyStatus_REPLIED 已回复
	CommentReplyStatus_REPLIED CommentReplyStatus = "REPLIED"
	// CommentReplyStatus_NO_REPLY 未回复
	CommentReplyStatus_NO_REPLY CommentReplyStatus = "NO_REPLY"
)
