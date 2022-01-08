package comment

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// Reply 评论回复
type Reply struct {
	// ReplyToCommentID 回复的一级评论ID
	ReplyToCommentID json.Number `json:"reply_to_comment_id,omitempty"`
	// ReplyID 回复ID
	ReplyID json.Number `json:"reply_id,omitempty"`
	// ReplyToReplyID 二级回复所属的上级回复ID，只有二级回复才有。二级回复：表示回复“评论中其他回复内容”
	ReplyToReplyID model.FlexUint64 `json:"reply_to_reply_id,omitempty"`
	// UserInfo 回复用户信息
	UserInfo ReplyUserInfo `json:"user_info,omitempty"`
	// ReplyText 回复内容
	ReplyText string `json:"reply_text,omitempty"`
	// CreateTime 回复创建时间
	CreateTime string `json:"create_time,omitempty"`
}

// ReplyUserInfo 回复用户信息
type ReplyUserInfo struct {
	// ScreenName 回复用户昵称
	ScreenName string `json:"screen_name,omitempty"`
}
