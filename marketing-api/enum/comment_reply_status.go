package enum

// CommentReplyStatus 评论回复状态
type CommentReplyStatus string

const (
	// CommentReplyStatus_REPLIED 已回复
	CommentReplyStatus_REPLIED CommentReplyStatus = "REPLIED"
	// CommentReplyStatus_NO_REPLY 未回复
	CommentReplyStatus_NO_REPLY CommentReplyStatus = "NO_REPLY"
	// CommentReplyStatus_REPLY_AUDIT_FAIL 回复审核拒接
	CommentReplyStatus_REPLY_AUDIT_FAIL CommentReplyStatus = "REPLY_AUDIT_FAIL"
	// CommentReplyStatus_REPLY_AUDIT_SUCCESS 回复审核通过
	CommentReplyStatus_REPLY_AUDIT_SUCCESS CommentReplyStatus = "REPLY_AUDIT_SUCCESS"
	// CommentReplyStatus_REPLY_TO_AUDIT 回复待审核
	CommentReplyStatus_REPLY_TO_AUDIT CommentReplyStatus = "REPLY_TO_AUDIT"
)
