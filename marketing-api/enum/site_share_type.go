package enum

// SiteShareType 站点来源
type SiteShareType string

const (
	// SiteShareType_SHARE 共享站点
	SiteShareType_SHARE SiteShareType = "SHARE"
	// SiteShareType_MY_CREATIONS 账户创建站点 (默认）
	SiteShareType_MY_CREATIONS SiteShareType = "MY_CREATIONS"
)
