package enum

// CommentHideStatus 隐藏状态
type CommentHideStatus string

const (
	// CommentHideStatus_ALL 全部
	CommentHideStatus_ALL CommentHideStatus = "ALL"
	// CommentHideStatus_HIDE 已隐藏
	CommentHideStatus_HIDE CommentHideStatus = "HIDE"
	// CommentHideStatus_NOT_HIDE 未隐藏
	CommentHideStatus_NOT_HIDE CommentHideStatus = "NOT_HIDE"
)
