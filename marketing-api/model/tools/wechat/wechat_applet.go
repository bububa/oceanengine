package wechat

import "github.com/bububa/oceanengine/marketing-api/enum"

// WechatApplet 微信小程序
type WechatApplet struct {
	// Name 小程序名称
	Name string `json:"name,omitempty"`
	// UserName 小程序原始ID
	UserName string `json:"user_name,omitempty"`
	// InstanceID 小程序资产ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Path 小程序路径
	Path string `json:"path,omitempty"`
	// RemarkMessage 资产备注信息，最大长度不超过15
	RemarkMessage string `json:"remark_message,omitempty"`
	// IconImageURL 小程序icon图片的url
	IconImageURL string `json:"icon_image_url,omitempty"`
	// HeaderImageURL 顶部头图的url
	HeaderImageURL string `json:"header_image_url,omitempty"`
	// Labels 小程序标签
	Labels []string `json:"labels,omitempty"`
	// GuideText 引导文案
	GuideText string `json:"guide_text,omitempty"`
	// ImagesVerticalURL 小程序竖图的url
	ImagesVerticalURL []string `json:"images_vertical_url,omitempty"`
	// ImagesHorizontalURL 小程序横图的url
	ImagesHorizontalURL []string `json:"images_horizontal_url,omitempty"`
	// Introduction 小程序简介
	Introduction string `json:"introduction,omitempty"`
	// AuditStatus 审核状态:
	// AUDIT_ACCEPTED 审核成功
	// AUDITING 审核中
	// AUDIT_REJECTED 审核失败
	AuditStatus enum.WechatAuditStatus `json:"audit_status,omitempty"`
	// Reason 审核拒绝原因
	Reason string `json:"reason,omitempty"`
	// AdvertiserID 所属广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTIme 修改时间
	ModifyTime string `json:"modify_time,omitempty"`
}
