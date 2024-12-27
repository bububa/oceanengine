package enum

// AccountPlatformType 业务线：AD/千川/本地推
type AccountPlatformType string

const (
	// AccountPlatformType_NO_LIMIT 不过滤
	AccountPlatformType_NO_LIMIT AccountPlatformType = "NO_LIMIT"
	// AccountPlatformType_ONLY_AD 仅巨量AD专用
	AccountPlatformType_ONLY_AD AccountPlatformType = "ONLY_AD"
	// AccountPlatformType_ONLY_AD_SHARED 仅巨量AD/千川/本地推专用
	AccountPlatformType_ONLY_AD_SHARED AccountPlatformType = "ONLY_AD_SHARED"
	// AccountPlatformType_ONLY_ECP 仅巨量千川专用
	AccountPlatformType_ONLY_ECP AccountPlatformType = "ONLY_ECP"
	// AccountPlatformType_ONLY_LOCAL 仅巨量本地推专用
	AccountPlatformType_ONLY_LOCAL AccountPlatformType = "ONLY_LOCAL"
)
