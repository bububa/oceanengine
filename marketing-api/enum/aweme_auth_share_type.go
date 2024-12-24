package enum

// AwemeAuthShareType 授权共享类型
type AwemeAuthShareType = string

const (
	// AwemeAuthShareType_SHARE_BY_ONESELF 广告账户自主授权
	AwemeAuthShareType_SHARE_BY_ONESELF AwemeAuthShareType = "SHARE_BY_ONESELF"
	// AwemeAuthShareType_SHARE_BY_SAME_ENTITY 客户共享授权
	AwemeAuthShareType_SHARE_BY_SAME_ENTITY AwemeAuthShareType = "SHARE_BY_SAME_ENTITY"
	// AwemeAuthShareType_SHARE_FROM_BP 组织共享授权
	AwemeAuthShareType_SHARE_FROM_BP AwemeAuthShareType = "SHARE_FROM_BP"
)
