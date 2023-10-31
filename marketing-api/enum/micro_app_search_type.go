package enum

// MicroAppSearchType 搜索类型，可选值:
type MicroAppSearchType string

const (
	// MicroAppSearchType_CREATE_ONLY 只查询该账户创建的应用（默认值）
	MicroAppSearchType_CREATE_ONLY MicroAppSearchType = "CREATE_ONLY"
	// MicroAppSearchType_SHARE_ONLY 只查询被共享的应用
	MicroAppSearchType_SHARE_ONLY MicroAppSearchType = "SHARE_ONLY"
)
