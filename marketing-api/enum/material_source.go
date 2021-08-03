package enum

// MaterialSource 素材来源
type MaterialSource string

const (
	// MaterialSource_AD_SITE ad后台本地上传
	MaterialSource_AD_SITE MaterialSource = "AD_SITE"
	// MaterialSource_CREATIVE_CENTER 创意中心
	MaterialSource_CREATIVE_CENTER MaterialSource = "CREATIVE_CENTER"
	// MaterialSource_OPEN_API 开放平台
	MaterialSource_OPEN_API MaterialSource = "OPEN_API"
	// MaterialSource_SUPPLIER 即合视频
	MaterialSource_SUPPLIER MaterialSource = "SUPPLIER"
	// MaterialSource_VIDEO_CAPTURE 易拍视频
	MaterialSource_VIDEO_CAPTURE MaterialSource = "VIDEO_CAPTURE"
	// MaterialSource_ACCOUNT_PUSH 推送视频
	MaterialSource_ACCOUNT_PUSH MaterialSource = "ACCOUNT_PUSH"
	// MaterialSource_STAR 星图视频
	MaterialSource_STAR MaterialSource = "STAR"
	// MaterialSource_CEWEBRITY_VIDEO 达人视频
	MaterialSource_CEWEBRITY_VIDEO MaterialSource = "CEWEBRITY_VIDEO"
	// MaterialSource_BP 巨量纵横
	MaterialSource_BP MaterialSource = "BP"
	// MaterialSource_E_COMMERCE 巨量千川
	MaterialSource_E_COMMERCE MaterialSource = "E_COMMERCE"
	// MaterialSource_OTHERS 其他来源
	MaterialSource_OTHERS MaterialSource = "OTHERS"
)
