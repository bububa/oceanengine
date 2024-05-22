package enum

// AwemeAuthType 授权类型
type AwemeAuthType string

const (
	// AwemeAuthType_AWEME_ACCOUNT 抖音号授权
	AwemeAuthType_AWEME_ACCOUNT AwemeAuthType = "AWEME_ACCOUNT"
	// AwemeAuthType_VIDEO_ITEM 单视频授权
	AwemeAuthType_VIDEO_ITEM AwemeAuthType = "VIDEO_ITEM"
	// AwemeAuthType_LIVE_ACCOUNT 直播授权
	AwemeAuthType_LIVE_ACCOUNT AwemeAuthType = "LIVE_ACCOUNT"
	// AwemeAuthType_AWEME_HOMEPAGE 主页作品授权
	AwemeAuthType_AWEME_HOMEPAGE AwemeAuthType = "AWEME_HOMEPAGE"
)
