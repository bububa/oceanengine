package wechat

import "github.com/bububa/oceanengine/marketing-api/enum"

// WechatGame 微信小游戏
type WechatGame struct {
	// Name 小程序名称
	Name string `json:"name,omitempty"`
	// UserName 小程序原始ID
	UserName string `json:"user_name,omitempty"`
	// InstanceID 小程序资产ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Path 小程序路径
	Path string `json:"path,omitempty"`
	// AuditStatus 审核状态:
	// AUDIT_ACCEPTED 审核成功
	// AUDITING 审核中
	// AUDIT_REJECTED 审核失败
	AuditStatus enum.WechatAuditStatus `json:"audit_status,omitempty"`
	// Reason 审核拒绝原因
	Reason string `json:"reason,omitempty"`
	// AuthorizationStatus 授权状态 枚举值：AUTHORIZED已授权、UNAUTHORIZED未授权、AUTHORIZATION_FAILED授权失败
	AuthorizationStatus enum.WechatAuthorizationStatus `json:"authorization_status,omitempty"`
	// ReasonUnauthorize 授权失败原因
	ReasonUnauthorize string `json:"reason_unauthorize,omitempty"`
	// AccountType 所属账户类型
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// AccountID 所属账户ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AntiAddictionURL 防沉迷提示URL，该url3小时内查看有效
	AntiAddictionURL string `json:"anti_addiction_url,omitempty"`
	// ScreenRecordURL 游戏内容视频URL，该url3小时内查看有效
	ScreenRecordURL string `json:"screen_record_url,omitempty"`
	// RealNameURL 实名认证URL，该url3小时内查看有效
	RealNameURL string `json:"real_name_url,omitempty"`
	// AgeRemindURL 适龄提醒URL，该url3小时内查看有效
	AgeRemindURL string `json:"age_remind_url,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTIme 修改时间
	ModifyTime string `json:"modify_time,omitempty"`
}
