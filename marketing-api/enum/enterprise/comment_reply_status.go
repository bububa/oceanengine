package enterprise

// CommentReplyStatus 回复状态
type CommentReplyStatus string

const (
	// CommentReplyStatus_NO_REPLY 未回复
	CommentReplyStatus_NO_REPLY CommentReplyStatus = "NO_REPLY"
	// CommentReplyStatus_REPLY 已回复
	CommentReplyStatus_REPLY CommentReplyStatus = "REPLY"
	// CommentReplyStatus_REPLY_TO_AUDIT 回复待审核
	CommentReplyStatus_REPLY_TO_AUDIT CommentReplyStatus = "REPLY_TO_AUDIT"
	// CommentReplyStatus_REPLY_AUDIT_FAIL 回复审核拒接
	CommentReplyStatus_REPLY_AUDIT_FAIL CommentReplyStatus = "REPLY_AUDIT_FAIL"
	// CommentReplyStatus_REPLY_AUDIT_SUCCESS 回复审核通过
	CommentReplyStatus_REPLY_AUDIT_SUCCESS CommentReplyStatus = "REPLY_AUDIT_SUCCESS"
	// CommentReplyStatus_DELETE已删除
	CommentReplyStatus_DELETE CommentReplyStatus = "DELETE"
	// CommentReplyStatus_PUBLISH 公开可见
	CommentReplyStatus_PUBLISH CommentReplyStatus = "PUBLISH"
	// CommentReplyStatus_SELF_VISIBLE 仅自己可见
	CommentReplyStatus_SELF_VISIBLE CommentReplyStatus = "SELF_VISIBLE"
	// CommentReplyStatus_PART_VISIBLE 部分可见
	CommentReplyStatus_PART_VISIBLE CommentReplyStatus = "PART_VISIBLE"
)
