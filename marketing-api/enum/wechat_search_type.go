package enum

// WechatSearchType  搜索类型
type WechatSearchType string

const (

	// WechatSearchType_CREATE_ONLY 只查询该账户创建的应用（默认值）
	WechatSearchType_CREATE_ONLY WechatSearchType = "CREATE_ONLY"
	// WechatSearchType_SHARE_ONLY 只查询被共享的应用
	WechatSearchType_SHARE_ONLY WechatSearchType = "SHARE_ONLY"
)
