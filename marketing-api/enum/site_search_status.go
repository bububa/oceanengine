package enum

// SiteSearchStatus 建站粗粒度状态
type SiteSearchStatus string

const (
	// SiteSearchStatus_SITE_ALL 查询所有
	SiteSearchStatus_SITE_ALL SiteSearchStatus = "SITE_ALL"
	// SiteSearchStatus_SITE_ONLINE 已上线（包括建站状态中：enable，auditAccepted，auditDoing）
	SiteSearchStatus_SITE_ONLINE SiteSearchStatus = "SITE_ONLINE"
	// SiteSearchStatus_SITE_OFFLINE 已下线（包括建站状态中：disable，edit）
	SiteSearchStatus_SITE_OFFLINE SiteSearchStatus = "SITE_OFFLINE"
	// SiteSearchStatus_SITE_REJECTED 审核不通过（包括建站状态中：auditRejected）
	SiteSearchStatus_SITE_REJECTED SiteSearchStatus = "SITE_REJECTED"
	// SiteSearchStatus_SITE_TRASH 删除或者禁止（包括建站状态中：deleted，auditBanned）
	SiteSearchStatus_SITE_TRASH SiteSearchStatus = "SITE_TRASH"
)
