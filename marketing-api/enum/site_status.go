package enum

// SiteStatus 站点状态
type SiteStatus string

const (
	// SiteStatus_ENABLE 已发布, 已生效（对应后台 “已上线”）
	SiteStatus_ENABLE SiteStatus = "ENABLE"
	// SiteStatus_DISABLE 不可用
	SiteStatus_DISABLE SiteStatus = "DISABLE"
	// SiteStatus_DELETED 已删除
	SiteStatus_DELETED SiteStatus = "DELETED"
	// SiteStatus_EDIT 未发布, 未生效 （对应后台“未上线”）
	SiteStatus_EDIT SiteStatus = "EDIT"
	// SiteStatus_FORBIDDENT 惩罚广告主, 站点下线
	SiteStatus_FORBIDDENT SiteStatus = "FORBIDDEN"
	// SiteStatus_AUDIT_ACCEPTED 审核通过
	SiteStatus_AUDIT_ACCEPTED SiteStatus = "AUDIT_ACCEPTED"
	// SiteStatus_AUDIT_REJECTED 审核拒绝
	SiteStatus_AUDIT_REJECTED SiteStatus = "AUDIT_REJECTED"
	// SiteStatus_AUDIT_BANNED 审核封禁
	SiteStatus_AUDIT_BANNED SiteStatus = "AUDIT_BANNED"
	// SiteStatus_AUDIT_DOING 审核中
	SiteStatus_AUDIT_DOING SiteStatus = "AUDIT_DOING"
)
