package comment

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Reply 评论回复
type Reply struct {
	// ReplyToCommentID 回复的一级评论ID
	ReplyToCommentID model.Uint64 `json:"reply_to_comment_id,omitempty"`
	// CommentID 评论id
	CommentID uint64 `json:"comment_id,omitempty"`
	// Text 评论内容
	Text string `json:"text,omitempty"`
	// CreateTime 评论创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ReplyStatus 评论回复状态，可选值:
	// NO_REPLY: 未回复
	// REPLY_AUDIT_FAIL: 回复审核拒接
	// REPLY_AUDIT_SUCCESS: 回复审核通过
	// REPLY_TO_AUDIT: 回复待审核
	ReplyStatus enum.CommentReplyStatus `json:"reply_status,omitempty"`
	// LikeCount 点赞数
	LikeCount int `json:"like_count,omitempty"`
	// HideStatus 隐藏状态, 可选值:
	// ALL: 全部
	// HIDE: 已隐藏
	// NOT_HIDE: 未隐藏
	HideStatus enum.CommentHideStatus `json:"hide_status,omitempty"`
	// IsStick 该评论是否置顶，0：表示不置顶，1：表示置顶
	IsStick int `json:"is_stick,omitempty"`
	// AwemeID 评论用户抖音号
	AwemeID string `json:"aweme_id,omitempty"`
	// AwemeName 评论用户昵称
	AwemeName string `json:"aweme_name,omitempty"`
	// ItemID 广告视频ID
	ItemID uint64 `json:"item_id,omitempty"`
}
