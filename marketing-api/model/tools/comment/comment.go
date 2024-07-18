package comment

import "github.com/bububa/oceanengine/marketing-api/enum"

// Comment 评论
type Comment struct {
	// ID 评论ID
	ID uint64 `json:"comment_id,omitempty"`
	// Text 评论内容
	Text string `json:"text,omitempty"`
	// CreateTime 评论创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ReplyCount 评论的回复数量
	ReplyCount int `json:"reply_count,omitempty"`
	// LikeCount 点赞数
	LikeCount int `json:"like_count,omitempty"`
	// ReplyStatus 评论回复状态，REPLIED表示已回复，NO_REPLY表示未回复
	ReplyStatus enum.CommentReplyStatus `json:"reply_status,omitempty"`
	// HideStatus 隐藏状态, 可选值:
	// ALL: 全部
	// HIDE: 已隐藏
	// NOT_HIDE: 未隐藏
	HideStatus enum.CommentHideStatus `json:"hide_status,omitempty"`
	// LevelType 查询的评论等级，可选值:
	// LEVEL_ALL: 所有评论（默认值）
	// LEVEL_ONE: 一级评论
	// LEVEL_TWO: 二级评论
	LevelType enum.CommentLevelType `json:"level_type,omitempty"`
	//     评论情感, 可选值:
	// NEGATIVE: 负向评论
	// NEUTRAL: 中性评论
	// POSITIVE: 正向评论
	EmotionType enum.CommentEmotionType `json:"emotion_type,omitempty"`
	// IsStick 该评论是否置顶，0：表示不置顶，1：表示置顶
	IsStick int `json:"is_stick,omitempty"`
	// AwemeID 评论用户抖音号
	AwemeID string `json:"aweme_id,omitempty"`
	// AwemeName 评论用户昵称
	AwemeName string `json:"aweme_name,omitempty"`
	// AdID 计划ID ，仅1.0平台适用
	AdID uint64 `json:"ad_id,omitempty"`
	// AdName 计划名
	AdName string `json:"ad_name,omitempty"`
	// CreativeID 评论所属创意id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// PromotionID 广告id，仅2.0平台适用
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// ItemID 广告视频ID
	ItemID uint64 `json:"item_id,omitempty"`
	// ItemTitle 视频标题
	ItemTitle string `json:"item_title,omitempty"`
	// CommentPermission 评论权限，允许值：
	// READ 只读
	// WRITE 可写
	CommentPermission enum.CommentPermission `json:"comment_permission,omitempty"`
	// MaterialID 评论关联的视频素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// CommentImages 评论图片内容，返回图片url列表
	CommentImages []string `json:"comment_images,omitempty"`
	// CommentType 评论内容类型
	// IMAGE_COMMENT 图片评论
	// IMAGE_TEXT_COMMENT 图文评论
	// TEXT_COMMENT 文字评论
	CommentType enum.CommentType `json:"comment_type,omitempty"`
}

// UserInfo 评论用户信息
type UserInfo struct {
	// UserScreenName 评论用户昵称
	UserScreenName string `json:"user_screen_name,omitempty"`
	// AwemeAccount 评论用户抖音号
	AwemeAccount string `json:"aweme_account,omitempty"`
}
