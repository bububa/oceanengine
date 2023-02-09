package enum

// SiteOptStatus 落地页组站点启用状态
type SiteOptStatus string

// 落地页组站点启用状态
const (
	// SITE_OPT_STATUS_ENABLE 启用
	SITE_OPT_STATUS_ENABLE SiteOptStatus = "SITE_OPT_STATUS_ENABLE"
	// SITE_OPT_STATUS_DISABLE 禁用
	SITE_OPT_STATUS_DISABLE SiteOptStatus = "SITE_OPT_STATUS_DISABLE"
)

// 站点操作状态
const (
	// SITE_OPT_STATUS_PUBLISHED 发布
	SITE_OPT_STATUS_PUBLISHED SiteOptStatus = "published"
	// SITE_OPT_STATUS_UNPUBLISHED 下线
	SITE_OPT_STATUS_UNPUBLISHED SiteOptStatus = "unpublished"
	// SITE_OPT_STATUS_DELETE 删除
	SITE_OPT_STATUS_DELETE SiteOptStatus = "delete"
	//SITE_OPT_STATUS_UNDELETE 恢复删除
	SITE_OPT_STATUS_UNDELETE SiteOptStatus = "undeleted"
)
