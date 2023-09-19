package comment

import "github.com/bububa/oceanengine/marketing-api/enum/enterprise"

// Comment 评论
type Comment struct {
	// CommentID 评论id
	CommentID uint64 `json:"comment_id,omitempty"`
	// CommentText 评论内容
	CommentText string `json:"comment_text,omitempty"`
	// CommentImage 评论图片url
	CommentImage string `json:"comment_image,omitempty"`
	// CreateTime 评论创建时间，时间格式 yyyy-mm-dd hh:mm:ss
	CreateTime string `json:"create_time,omitempty"`
	// ReplyCount 评论回复数
	ReplyCount int `json:"reply_count,omitempty"`
	// DiggCount 评论点赞数
	DiggCount int `json:"digg_count,omitempty"`
	// CommentAwemeName 评论用户昵称
	CommentAwemeName string `json:"comment_aweme_name,omitempty"`
	// CommentOpenID 评论用户open_id
	CommentOpenID string `json:"comment_open_id,omitempty"`
	// IsBannedUser 评论用户是否被屏蔽
	IsBannedUser int `json:"is_banned_user,omitempty"`
	// IsStick 评论是否被置顶
	IsStick int `json:"is_stick,omitempty"`
	// Level 评论层级
	// 枚举值：LEVEL_ONE一级评论、LEVEL_TWO二级评论
	Level enterprise.CommentLevel `json:"level,omitempty"`
	// Source 流量来源
	// 枚举值：FROM_NATURAL自然流量、FROM_DOUPLUSDou+、FROM_PERFORM竞价广告、FROM_BRAND品牌广告、FROM_OTHER其他流量
	Source enterprise.CommentSource `json:"source,omitempty"`
	// AdvertiserID 流量来源的广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 流量来源的广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// AdID 流量来源的广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// CreativeID 流量来源的广告创意id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// MaterialID 评论所属的广告视频素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// ItemID 评论所属的抖音视频id
	ItemID uint64 `json:"item_id,omitempty"`
	// ItemAwemeName 评论所属的抖音视频作者昵称
	ItemAwemeName string `json:"item_aweme_name,omitempty"`
	// ItemOpenID 评论所属的抖音视频作者open_id
	ItemOpenID string `json:"item_open_id,omitempty"`
	// RepliedCommentID 评论所属的广告视频素材id
	RepliedCommentID uint64 `json:"replied_comment_id,omitempty"`
}
