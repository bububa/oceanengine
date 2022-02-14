package landinggroup

import "github.com/bububa/oceanengine/marketing-api/enum"

// Site 站点
type Site struct {
	// MemberID 成员 ID，即站点在落地页组中的唯一标识
	MemberID uint64 `json:"member_id,omitempty"`
	// SiteID 站点 ID
	SiteID uint64 `json:"site_id,omitempty"`
	// SiteURL 站点URL
	SiteURL string `json:"site_url,omitempty"`
	// SiteOptStatus 站点启用状态
	SiteOptStatus enum.SiteOptStatus `json:"site_opt_status,omitempty"`
	// SiteAuditStatus 站点审核状态
	SiteAuditStatus enum.SiteAuditStatus `json:"site_audit_status,omitempty"`
}
