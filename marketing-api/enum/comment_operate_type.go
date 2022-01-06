package enum

// CommentOperateType 评论操作类型
type CommentOperateType string

const (
	// CommentOperateType_REPLY 回复评论，可批量
	CommentOperateType_REPLY CommentOperateType = "REPLY"
	// CommentOperateType_REPLY_TO_REPLY 二级回复，回复该评论下的其他回复
	CommentOperateType_REPLY_TO_REPLY CommentOperateType = "REPLY_TO_REPLY"
	// CommentOperateType_STICK_ON_TOP 置顶
	CommentOperateType_STICK_ON_TOP CommentOperateType = "STICK_ON_TOP"
	// CommentOperateType_CANCEL_STICK 取消置顶
	CommentOperateType_CANCEL_STICK CommentOperateType = "CANCEL_STICK"
	// CommentOperateType_HIDE 隐藏评论，仅评论者自己可见
	CommentOperateType_HIDE CommentOperateType = "HIDE"
	// CommentOperateType_BLOCK_USERS 屏蔽用户，将用户加入黑名单，该用户之后的评论自动设置为仅评论者自见
	CommentOperateType_BLOCK_USERS CommentOperateType = "BLOCK_USERS"
)
