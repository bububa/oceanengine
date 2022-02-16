package enum

// SiteOptStatus 落地页组站点启用状态
type SiteOptStatus string

const (
	// SITE_OPT_STATUS_ENABLE 启用
	SITE_OPT_STATUS_ENABLE SiteOptStatus = "SITE_OPT_STATUS_ENABLE"
	// SITE_OPT_STATUS_DISABLE 禁用
	SITE_OPT_STATUS_DISABLE SiteOptStatus = "SITE_OPT_STATUS_DISABLE"
)
