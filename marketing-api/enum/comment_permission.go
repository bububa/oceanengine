package enum

// CommentPermission 评论权限
type CommentPermission string

const (
	// CommentPermission_READ 只读
	CommentPermission_READ CommentPermission = "READ"
	// CommentPermission_WRITE 可写
	CommentPermission_WRITE CommentPermission = "WRITE"
)
