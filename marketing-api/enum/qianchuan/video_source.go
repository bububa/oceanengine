package qianchuan

// VideoSource 视频来源
type VideoSource = string

const (
	// VideoSource_E_COMMERCE 本地上传
	VideoSource_E_COMMERCE VideoSource = "E_COMMERCE"
	// VideoSource_LIVE_HIGHLIGHT 直播剪辑素材
	VideoSource_LIVE_HIGHLIGHT VideoSource = "LIVE_HIGHLIGHT"
	// VideoSource_BP 巨量纵横共享素材
	VideoSource_BP VideoSource = "BP"
	// VideoSource_VIDEO_CAPTURE 易拍APP共享素材
	VideoSource_VIDEO_CAPTURE VideoSource = "VIDEO_CAPTURE"
	// VideoSource_ARTHUR 亚瑟共享素材
	VideoSource_ARTHUR VideoSource = "ARTHUR"
	// VideoSource_STAR 星图&即合共享素材
	VideoSource_STAR VideoSource = "STAR"
	// VideoSource_TADA tada共享素材
	VideoSource_TADA VideoSource = "TADA"
	// VideoSource_CREATIVE_CENTER 巨量创意PC共享素材
	VideoSource_CREATIVE_CENTER VideoSource = "CREATIVE_CENTER"
)
