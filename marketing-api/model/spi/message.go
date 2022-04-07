package spi

import "encoding/json"

// Message 正文数据
type Message struct {
	// MessageID 唯一标识一条推送消息/数据; 64字符
	MessageID string `json:"message_id,omitempty"`
	// SubscribeTaskID 订阅任务id，订阅任务的主键; 64字符
	SubscribeTaskID string `json:"subscribe_task_id,omitempty"`
	// AdvertiserIDs 消息对应的广告主账号（全量，包含account_relation中所有map的value中的adv_id值）; min=1，max=1k
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// AccountRelation 推送广告主账号的订阅来源，是订阅的哪个账号下的广告主数据; min=1，max=1k
	AccountRelation string `json:"account_relation,omitempty"`
	// ServiceLabel 订阅服务类型
	ServiceLabel string `json:"service_label,omitempty"`
	// Data 推送数据信息，具体结构参考子文档
	Data string `json:"data,omitempty"`
	// PublishTime 消息实际产生时间
	PublishTime int64 `json:"publish_time,omitempty"`
	// Timestamp 毫秒时间戳，本条消息实际推送时间
	Timestamp int64 `json:"timestamp,omitempty"`
	// Nonce 随机数，和timestamp组合防重放
	Nonce int64 `json:"nonce,omitempty"`
}

// AccountRelation 账号对应关系，包含所有advertiser_ids的授权关系
type AccountRelation struct {
	// AgentIDs 代理商授权
	AgentIDs []uint64 `json:"agent_ids,omitempty"`
	// CoreUserIDs 用户授权账号
	CoreUserIDs []uint64 `json:"core_user_ids,omitempty"`
	// AdvertiserIDs 广告主id直接授权
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// CcIDs 巨量纵横账号
	Ccids []uint64 `json:"cc_ids,omitempty"`
}

// MessageData .
type MessageData struct {
	// AppID 应用APPID
	AppID uint64 `json:"app_id,omitempty"`
	// UserID 创意归属广告主id
	UserID uint64 `json:"user_id,omitempty"`
	// FromUserID 触发事件发送者open_id
	FromUserID string `json:"from_user_id,omitempty"`
	// ToUserID 触发事件接收者open_id
	ToUserID string `json:"to_user_id,omitempty"`
	// CoreUserID 登陆用户id
	CoreUserID uint64 `json:"core_user_id,omitempty"`
	// CoreUserIDs 登陆用户ids
	CoreUserIDs []uint64 `json:"core_user_ids,omitempty"`
	// AdID 创意归属计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// AdIDs 聚合周期内发生指标变更的计划列表，service_label 为report.ad.activeprogram时有值
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// CreativeIDs 聚合周期内发生指标变更的创意列表，service_label 为report.creative.activeprogram时有值
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// Event 事件名称
	Event string `json:"event,omitempty"`
	// Content 事件内容
	Content *EventContent `json:"content,omitempty"`
}

// EventContent 事件内容
type EventContent struct {
	// SiteID 站点ID
	SiteID uint64 `json:"site_id,omitempty"`
	// EventType 事件变更类型，枚举值：
	// 1 站点状态变更 、2 站点审核状态变更 、3 站点影子审核状态变更 、4 站点线上内容数据变更 、5 站点送审内容数据变更
	EventType int `json:"event_type,omitempty"`
	// AdvID 广告主ID
	AdvID uint64 `json:"adv_id,omitempty"`
	// From 变更前数据
	From json.RawMessage `json:"from,omitempty"`
	// To 变更后数据
	To json.RawMessage `json:"to,omitempty"`
	// Extra 额外数据
	Extra string `json:"extra,omitempty"`
	// MessageType 消息类型，枚举值：text 文本消息、emoji 表情消息、h5/poi H5/POI类卡片消息
	MessageType string `json:"message_type,omitempty"`
	// Text 文本消息内容
	Text string `json:"text,omitempty"`
	// ResourceType 资源类型
	ResourceType string `json:"resource_type,omitempty"`
	// ResourceHeight 资源高度
	ResourceHeight int `json:"resource_height,omitempty"`
	// ResourceWidth 资源宽度
	ResourceWidth int `json:"resource_width,omitempty"`
	// ResourceURL 资源链接
	ResourceURL string `json:"resource_url,omitempty"`
	// Title 卡片标题
	Title string `json:"title,omitempty"`
	// IconURL 卡片图标
	IconURL string `json:"icon_url,omitempty"`
	// Description 卡片描述
	Description string `json:"description,omitempty"`
	// LinkURL 跳转链接
	LinkURL string `json:"link_url,omitempty"`
	// Object 来源场景对应id（video对应视频id）
	Object string `json:"object,omitempty"`
	// CommentUser 用户信息
	CommentUser *CommentUser `json:"comment_user,omitempty"`
	// CommentContent 评论内容
	CommentContent string `json:"comment_content,omitempty"`
	// CommentItemOwnerID 内容所属者open_id
	CommentItemOwnerID string `json:"comment_item_owner_id,omitempty"`
	// CreateTime 评论时间
	CreateTime int64 `json:"create_time,omitempty"`
	// IsReply 是否回复
	IsReply bool `json:"is_reply,omitempty"`
	// ReplyToCommentID 回复的评论id
	ReplyToCommentID string `json:"reply_to_comment_id,omitempty"`
	// ReplyToCommentUser 回复评论的用户信息
	ReplyToCommentUser *CommentUser `json:"reply_to_comment_user,omitempty"`
	// ReplyToCommentContent 回复的评论内容
	ReplyToCommentContent string `json:"reply_to_comment_content,omitempty"`
}

// CommentUser 用户信息
type CommentUser struct {
	// OpenID 抖音ID
	OpenID string `json:"open_id,omitempty"`
	// AvatarURL 头像
	AvatarURL string `json:"avatar_url,omitempty"`
	// NickName 昵称
	NickName string `json:"nick_name,omitempty"`
}
