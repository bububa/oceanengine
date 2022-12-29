package enum

// CommentStickType 操作类型
type CommentStickType string

const (
	// CommentStickType_CANCEL_STICK 取消置顶
	CommentStickType_CANCEL_STICK CommentStickType = "CANCEL_STICK"
	// CommentStickType_STICK_ON_TOP 置顶
	CommentStickType_STICK_ON_TOP CommentStickType = "STICK_ON_TOP"
)
