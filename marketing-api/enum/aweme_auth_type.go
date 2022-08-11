package enum

// AwemeAuthType 授权类型
type AwemeAuthType string

const (
	// AwemeAuthType_AWEME_ACCOUNT 抖音号授权
	AwemeAuthType_AWEME_ACCOUNT AwemeAuthType = "AWEME_ACCOUNT"
	// AwemeAuthType_VIDEO_ITEM 单视频授权
	AwemeAuthType_VIDEO_ITEM AwemeAuthType = "VIDEO_ITEM"
)
