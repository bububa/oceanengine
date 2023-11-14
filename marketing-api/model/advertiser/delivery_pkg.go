package advertiser

import "github.com/bububa/oceanengine/marketing-api/enum"

// DeliveryPkg 推广产品资质信息
type DeliveryPkg struct {
	// PkgID 推广产品组id
	PkgID uint64 `json:"pkg_id,omitempty"`
	// ConfigID 来自【推广产品资质规则配置查询接口】，行业的资质规则中的config_id
	ConfigID uint64 `json:"config_id,omitempty"`
	// ProductName 用户提交的推广产品名称
	ProductName string `json:"product_name,omitempty"`
	// Status 推广产品整组的审核状态，标识该推广产品是否审核通过 可选值:
	// STATUS_CONFIRM 审核通过
	// STATUS_CONFIRM_FAIL 审核不通过
	// STATUS_NOT_SUBMIT 未提交
	// STATUS_PENDING_CONFIRM 审核中
	// STATUS_WAIT_CONFIRM 待审核
	Status enum.DeliveryPkgStatus `json:"status,omitempty"`
	// IndustryID 一级到三级行业id
	IndustryID []uint64 `json:"industry_id,omitempty"`
	// IndustryName 一级到三级行业名称
	IndustryName []string `json:"industry_name,omitempty"`
	// NecessaryCombine 必填资质模块
	NecessaryCombine []DeliveryPkgCombine `json:"necessary_combine,omitempty"`
	// UnnecessaryCombine 选填资质模块
	UnnecessaryCombine []DeliveryPkgCombine `json:"unnecessary_combine,omitempty"`
	// Permission 权限信息
	Permission *DeliveryPkgPermission `json:"permission,omitempty"`
}

// DeliveryPkgCombine 资质模块
type DeliveryPkgCombine struct {
	// CombineID 推广类型id，来自【推广产品资质规则配置查询接口】，行业的资质规则中的promotion_type_id
	CombineID uint64 `json:"combine_id,omitempty"`
	// Description 推广类型描述
	Description string `json:"description,omitempty"`
	// DeliveryRules 资质规则
	DeliveryRules []DeliveryRule `json:"delivery_rules,omitempty"`
}

// DeliveryRules 资质规则
type DeliveryRule struct {
	// RuleID 原子规则id
	RuleID uint64 `json:"rule_id,omitempty"`
	// Deliveries 资质的具体信息
	Deliveries []Delivery `json:"deliveries,omitempty"`
}

// Delivery 资质的具体信息
type Delivery struct {
	// QualificationID 资质id
	QualificationID uint64 `json:"qualification_id,omitempty"`
	// QualType 资质类型id
	QualType uint64 `json:"qual_type,omitempty"`
	// QualTypeName 资质类型名称
	QualTypeName string `json:"qual_type_name,omitempty"`
	// Status 资质审核状态 可选值:
	// STATUS_CONFIRM 审核通过
	// STATUS_CONFIRM_FAIL 审核不通过
	// STATUS_NOT_SUBMIT 未提交
	// STATUS_PENDING_CONFIRM 审核中
	// STATUS_WAIT_CONFIRM 待审核
	Status enum.DeliveryPkgStatus `json:"status,omitempty"`
	// Attachments 资质图片附件
	Attachments []DeliveryAttachment `json:"attachments,omitempty"`
	// RejectReason 拒绝理由，若资质被拒绝，则会有拒绝理由
	RejectReason string `json:"reject_reason,omitempty"`
}

// DeliveryAttachment 资质图片附件
type DeliveryAttachment struct {
	// AttachmentID 附件id
	AttachmentID uint64 `json:"attachment_id,omitempty"`
	// PictureURL 图片链接
	PictureURL string `json:"picture_url,omitempty"`
}

// DeliveryPkgPermission 权限信息
type DeliveryPkgPermission struct {
	// CanEdit 是否支持编辑
	CanEdit bool `json:"can_edit,omitempty"`
	// CantEditReason 不支持编辑的原因
	CantEditReason string `json:"cant_edit_reason,omitempty"`
	// CanDelete 是否支持删除
	CanDelete bool `json:"can_delete,omitempty"`
	// CantDeleteReason 不支持删除的原因
	CantDeleteReason string `json:"cant_delete_reason,omitempty"`
}
