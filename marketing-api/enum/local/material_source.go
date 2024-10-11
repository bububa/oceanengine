package local

// MaterialSource 素材来源
type MaterialSource string

const (
	// MaterialSource_BP_PLATFORM 巨量引擎工作平台共享视频
	MaterialSource_BP_PLATFORM MaterialSource = "BP_PLATFORM"
	// MaterialSource_CREATIVE_AIGC 即创
	MaterialSource_CREATIVE_AIGC MaterialSource = "CREATIVE_AIGC"
	// MaterialSource_LOCAL_ADS_UPLOAD 本地上传
	MaterialSource_LOCAL_ADS_UPLOAD MaterialSource = "LOCAL_ADS_UPLOAD"
	// MaterialSource_STAR 星图平台
	MaterialSource_STAR MaterialSource = "STAR"
	// MaterialSource_MAPI MAPI上传
	MaterialSource_MAPI MaterialSource = "MAPI"
)
