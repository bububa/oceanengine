package spi

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
)

// Message 正文数据
type Message struct {
	// MessageID 唯一标识一条推送消息/数据; 64字符
	MessageID string `json:"message_id,omitempty"`
	// AccountRelation 推送广告主账号的订阅来源，是订阅的哪个账号下的广告主数据; min=1，max=1k
	AccountRelation string `json:"account_relation,omitempty"`
	// ServiceLabel 订阅服务类型
	ServiceLabel string `json:"service_label,omitempty"`
	// Data 推送数据信息，具体结构参考子文档
	Data string `json:"data,omitempty"`
	// AdvertiserIDs 消息对应的广告主账号，report.advertiser.beforeday 时返回所有已产出的广告主ID组，其余服务类型一次仅返回一个广告主ID
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// SubscribeTaskID 订阅任务id，订阅任务的主键; 64字符
	SubscribeTaskID uint64 `json:"subscribe_task_id,omitempty"`
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
	AgentIDs map[string][]uint64 `json:"agent_ids,omitempty"`
	// CoreUserIDs 用户授权账号
	CoreUserIDs map[string][]uint64 `json:"core_user_ids,omitempty"`
	// CcIDs 巨量纵横账号
	Ccids map[string][]uint64 `json:"cc_ids,omitempty"`
	// AdvertiserIDs 广告主id直接授权
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
}

// MessageDataBase .
type MessageDataBase struct {
	// Event 事件名称
	Event string `json:"event,omitempty"`
	// FromUserID 触发事件发送者open_id
	FromUserID string `json:"from_user_id,omitempty"`
	// ToUserID 触发事件接收者open_id
	ToUserID string `json:"to_user_id,omitempty"`
	// CoreUserIDs 登陆用户ids
	CoreUserIDs []uint64 `json:"core_user_ids,omitempty"`
	// AdIDs 聚合周期内发生指标变更的计划列表，service_label 为report.ad.activeprogram时有值
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// CreativeIDs 聚合周期内发生指标变更的创意列表，service_label 为report.creative.activeprogram时有值
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// ProjectIDs 状态发生变更的项目ids
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
	// PromotionIDs 状态发生变更的广告ids
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
	// MaterialIDs 聚合周期内发生指标变更的素材列表，service_label 为report.material.activeprogram时有值
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// AppID 应用APPID
	AppID uint64 `json:"app_id,omitempty"`
	// UserID 创意归属广告主id
	UserID uint64 `json:"user_id,omitempty"`
	// CoreUserID 登陆用户id
	CoreUserID uint64 `json:"core_user_id,omitempty"`
	// AdvertiserID 广告主I D
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 创意归属计划id
	AdID uint64 `json:"ad_id,omitempty"`
}

// MessageData .
type MessageData struct {
	// Content 事件内容
	Content *EventContent `json:"content,omitempty"`
	MessageDataBase
}

// MessageDataV2 .
type MessageDataV2 struct {
	// Content 事件内容
	Content string `json:"content,omitempty"`
	MessageDataBase
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
	// Message 通知内容，示例："尊敬的客户，您好！因您当前账户下的【所有广告】低效在投计划数过多，已经严重影响您的账户健康度，根据平台相关规则，【所有广告】已被【限制投放】。限制不会自动解除。建议您及时清理账户中的低效计划，完成清理后，可以联系和您对接的销售同学进行解除限制申诉（如您没有对接销售可通过平台右下方智能客服提工单反馈），感谢您对巨量引擎的支持。"
	Message string `json:"message,omitempty"`
	// AdIDs 被暂停的计划列表
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// GroupIDs 被暂停的管家项目列表
	GroupIDs []uint64 `json:"group_ids,omitempty"`
	// ProjectIDs 状态发生变更的项目ids
	ProjectIDs []uint64 `json:"project_ids,omitempty"`
	// PromotionIDs 成功占用配额的巨量引擎体验版的广告列表
	PromotionIDs []uint64 `json:"promotion_ids,omitempty"`
	// SumQuota 变化后在投计划配额
	SumQuota int64 `json:"sum_quota,omitempty"`
	// SumQuotaBefore 变化前在投计划配额
	SumQuotaBefore int64 `json:"sum_quota_before,omitempty"`
	// CampaignType 广告组类型，FEED 信息流、 SEARCH 搜索广告
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// DeliveryRange 允许值: DEFAULT默认（默认值）, UNION穿山甲,
	// 当campaign_type = feed && delivery_range= default可获取所有广告配额;
	// 当campaign_type = feed && delivery_range= union时可获取穿山甲广告配额;
	// 当campaign_type = search && delivery_range= default时可获取搜索广告配额;
	DeliveryRange enum.AdDeliveryRange `json:"delivery_range,omitempty"`
	// QualificationType 资质类型
	// AGENT_ARRANGEMENT 广告代投协议
	// AUTHORIZATION_CONTRACT 委托加工/生产合同
	// DISTRIBUTION_AUTHORIZATION 经销授权书
	// ICP_RECORD_AUTHORIZATION icp域名备案/授权
	// OTHER_CERTIFICATION 其他资质
	// PATENT_CERTIFICATE 专利证书
	// PORTRAIT_AUTHORIZATION 肖像授权书
	// QUALITY_REPORT 鉴定/质检报告
	// SOFTWARE_COPYRIGHT_REGISTRATION_CERTIFICATE 软件著作权登记证书
	// TRADEMARK_REGISTRATION_CERTIFICATE 商标注册证
	// VIDEO_MATERIAL_PRODUCTION_IP_AUTHORIZATION 广告视频/素材/作品/IP授权
	QualificationType enum.DeliveryQualificationType `json:"qualification_type,omitempty"`
	// QualificationID 资质ID
	QualificationID []uint64 `json:"qualification_id,omitempty"`
	// Status 状态
	Status string `json:"status,omitempty"`
	// MaterialIDs 聚合周期内发生指标变更的素材列表
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// MaterialType 素材类型：
	// image：图片
	// video：视频
	MaterialType string `json:"material_type,omitempty"`
	// IllegalMaterialIDs 违规素材id list
	IllegalMaterialIDs []uint64 `json:"illegal_material_ids,omitempty"`
	// EventID 违规单id
	EventID uint64 `json:"event_id,omitempty"`
	// BusinessLine 业务线AD
	BusinessLine string `json:"business_line,omitempty"`
	// AdID 计划id（AD2.0为广告id）
	AdID uint64 `json:"ad_id,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// ViolationEvidenceImg 违规证据截图
	ViolationEvidenceImg string `json:"violation_evidence_img,omitempty"`
	// Score 扣罚分值
	Score int64 `json:"score,omitempty"`
	// RejectReason 拒绝理由
	RejectReason string `json:"reject_reason,omitempty"`
	// IllegalType 违规类型 可选值:
	// GENERAL 一般违规（AD）
	// SERIOUS 严重违规（AD）
	IllegalType enum.SecurityScoreIllegalType `json:"illegal_type,omitempty"`
	// DisposalID 处置单id
	DisposalID uint64 `json:"disposal_id,omitempty"`
	// DisposalTerm 执行时长
	DisposalTerm string `json:"disposal_term,omitempty"`
	// DisposalStartTime 处置开始时间
	DisposalStartTime string `json:"disposal_start_time,omitempty"`
	// DisposalEndTime 处置结束时间
	DisposalEndTime string `json:"disposal_end_time,omitempty"`
	// DisposalReason 处置原因
	DisposalReason string `json:"disposal_reason,omitempty"`
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
