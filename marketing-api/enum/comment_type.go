package enum

// CommentType 评论内容类型
type CommentType string

const (
	// IMAGE_COMMENT 图片评论
	IMAGE_COMMENT CommentType = "IMAGE_COMMENT"
	// IMAGE_TEXT_COMMENT 图文评论
	IMAGE_TEXT_COMMENT CommentType = "IMAGE_TEXT_COMMENT"
	// TEXT_COMMENT 文字评论
	TEXT_COMMENT CommentType = "TEXT_COMMENT"
)
