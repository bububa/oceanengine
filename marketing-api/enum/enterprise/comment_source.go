package enterprise

// CommentSource 流量来源
type CommentSource string

const (
	// FROM_NATURAL 自然流量
	FROM_NATURAL CommentSource = "FROM_NATURAL"
	// FROM_DOUPLUS Dou+
	FROM_DOUPLUS CommentSource = "FROM_DOUPLUS"
	// FROM_PERFORM 竞价广告
	FROM_PERFORM CommentSource = "FROM_PERFORM"
	// FROM_BRAND 品牌广告
	FROM_BRAND CommentSource = "FROM_BRAND"
	// FROM_OTHER 其他流量
	FROM_OTHER CommentSource = "FROM_OTHER"
)
