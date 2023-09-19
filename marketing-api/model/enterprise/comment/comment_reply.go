package comment

import "github.com/bububa/oceanengine/marketing-api/enum/enterprise"

// CommentReply 回复
type CommentReply struct {
	// ReplyID 回复id
	ReplyID uint64 `json:"reply_id,omitempty"`
	// CommentID 回复成功的评论id
	CommentID uint64 `json:"comment_id,omitempty"`
	// RepliedCommentID 回复成功的二级评论id
	RepliedCommentID uint64 `json:"replied_comment_id,omitempty"`
	// ReplyText 回复内容
	ReplyText string `json:"reply_text,omitempty"`
	// RepliedCommentText 回复成功的二级评论内容
	RepliedCommentText string `json:"replied_comment_text,omitempty"`
	// CreateTime 回复时间
	CreateTime string `json:"create_time,omitempty"`
	// ReplyStatus 回复状态
	// 枚举值： REPLY_TO_AUDIT回复待审核、REPLY_AUDIT_FAIL回复审核拒接、REPLY_AUDIT_SUCCESS回复审核通过
	ReplyStatus enterprise.CommentReplyStatus `json:"reply_status,omitempty"`
	// Status 回复状态
	// 枚举值： DELETE已删除、PUBLISH公开可见、SELF_VISIBLE仅自己可见、PART_VISIBLE部分可见
	Status enterprise.CommentReplyStatus `json:"status,omitempty"`
	// ReplyAwemeName 回复用户昵称
	ReplyAwemeName string `json:"reply_aweme_name,omitempty"`
	// ReplyOpenID 回复用户open_id
	ReplyOpenID string `json:"reply_open_id,omitempty"`
	// IsEDouyinID 是否是企业号自己
	IsEDouyinID bool `json:"is_e_douyin_id,omitempty"`
}
