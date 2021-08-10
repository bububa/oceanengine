package thirdsite

import "github.com/bububa/oceanengine/marketing-api/enum"

// Site 站点
type Site struct {
	// AuditStatus 站点审核状态
	AuditStatus enum.SiteAuditStatus `json:"audit_status,omitempty"`
	// CreateTime 站点创建时间
	CreateTime string `json:"create_time,omitempty"`
	// Name 站点名称
	Name string `json:"name,omitempty"`
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
	// Thumbnail 缩略图地址
	Thumbnail string `json:"thumbnail,omitempty"`
	// URL 站点预览地址
	URL string `json:"url,omitempty"`
}
