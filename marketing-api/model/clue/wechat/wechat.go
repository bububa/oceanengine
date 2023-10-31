package wechat

// WechatType 微信号类型
type WechatType string

const (
	// WechatType_PERSONAL 个人号
	WechatType_PERSONAL WechatType = "PERSONAL"
	// WechatType_PUBLIC 公众号
	WechatType_PUBLIC WechatType = "PUBLIC"
	// WechatType_ENTERPRISE 企业微信
	WechatType_ENTERPRISE WechatType = "ENTERPRISE"
)

// AuditStatus 审核状态
type AuditStatus string

const (
	// AuditStatus_AUDITING 审核中
	AuditStatus_AUDITING AuditStatus = "AUDITING"
	// AuditStatus_AUDIT_REJECTED 审核拒绝
	AuditStatus_AUDIT_REJECTED AuditStatus = "AUDIT_REJECTED"
	// AuditStatus_AUDIT_ACCEPTED 审核通过
	AuditStatus_AUDIT_ACCEPTED AuditStatus = "AUDIT_ACCEPTED"
)

// Wechat 微信号
type Wechat struct {
	// WechatNumber 微信号
	WechatNumber string `json:"wechat_number,omitempty"`
	// WechatType 微信号类型。 允许值:
	// PERSONAL: 个人号
	// PUBLIC: 公众号
	// ENTERPRISE: 企业微信
	WechatType WechatType `json:"wechat_type,omitempty"`
	// NickName 微信昵称
	NickName string `json:"nick_name,omitempty"`
	// Suffix 尾缀
	Suffix string `json:"suffix,omitempty"`
	// AuditStatus 审核状态 允许值:
	// AUDITING: 审核中
	// AUDIT_REJECTED: 审核拒绝
	// AUDIT_ACCEPTED: 审核通过
	AuditStatus AuditStatus `json:"audit_status,omitempty"`
	// RejectReason 审核拒绝理由
	RejectReason string `json:"reject_reason,omitempty"`
	// HasQrCode 是否配置二维码 允许值:
	// TRUE: 是
	// FALSE: 否
	HasQrCode string `json:"has_qr_code,omitempty"`
	// AppID 公众号的开发者ID
	AppID string `json:"app_id,omitempty"`
	// InstanceCount 生效号码包数量
	InstanceCount int64 `json:"instance_count,omitempty"`
	// ClueCount 复制次数
	ClueCount int64 `json:"clue_count,omitempty"`
	// CreateTime 创建时间，格式：2022-07-21 15:36:09
	CreateTime string `json:"create_time,omitempty"`
	// ModTime 修改时间，格式：2022-07-21 15:36:09
	ModTime string `json:"mod_time,omitempty"`
	// Status 启用状态 0启用 1禁用
	Status int `json:"status,omitempty"`
}
