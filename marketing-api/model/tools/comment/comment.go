package comment

import "github.com/bububa/oceanengine/marketing-api/enum"

// Comment 评论
type Comment struct {
	// ID 评论ID
	ID uint64 `json:"id,omitempty"`
	// AppName app名
	AppName string `json:"app_name,omitempty"`
	// UserInfo 评论用户信息
	UserInfo UserInfo `json:"user_info,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// AdName 计划名
	AdName string `json:"ad_name,omitempty"`
	// CreativeID 评论所属创意id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// Text 评论内容
	Text string `json:"text,omitempty"`
	// LikeCount 点赞数
	LikeCount int `json:"like_count,omitempty"`
	// ReplyCount 评论的回复数量
	ReplyCount int `json:"reply_count,omitempty"`
	// Stick 是否置顶，0：表示不置顶，1：表示置顶
	Stick int `json:"stick,omitempty"`
	// InBlackList 评论发布者是否在黑名单内，0表示不在，1表示在
	InBlackList int `json:"in_black_list,omitempty"`
	// ReplyStatus 评论回复状态，REPLIED表示已回复，NO_REPLY表示未回复
	ReplyStatus enum.CommentReplyStatus `json:"reply_status,omitempty"`
	// CreateTime 评论创建时间
	CreateTime string `json:"create_time,omitempty"`
}

// UserInfo 评论用户信息
type UserInfo struct {
	// UserScreenName 评论用户昵称
	UserScreenName string `json:"user_screen_name,omitempty"`
	// AwemeAccount 评论用户抖音号
	AwemeAccount string `json:"aweme_account,omitempty"`
}
