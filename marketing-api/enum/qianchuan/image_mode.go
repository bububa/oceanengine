package qianchuan

// ImageMode 素材类型
type ImageMode string

const (
	// ImageMode_SMALL 小图，宽高比1.52，大小1.5M以下，下限：456 & 300，上限：1368 & 900
	ImageMode_SMALL ImageMode = "SMALL"
	// ImageMode_LARGE 大图，横版大图宽高比1.78，大小1.5M以下，下限：1280 & 720，上限：2560 & 1440
	ImageMode_LARGE ImageMode = "LARGE"
	// ImageMode_VIDEO_LARGE 横版视频，封面图宽高比1.78（下限：1280 & 720，上限：2560 & 1440））,视频资源支持mp4、mpeg、3gp、avi文件格式，宽高比16:9，大小不超过1000M
	ImageMode_VIDEO_LARGE ImageMode = "VIDEO_LARGE"
	// ImageMode_VIDEO_VERTICAL 竖版视频，封面图宽高比0.56（9:16），下限：720 & 1280，上限：1440 & 2560，视频资源支持mp4、mpeg、3gp、avi文件格式，大小不超过100M
	ImageMode_VIDEO_VERTICAL ImageMode = "VIDEO_VERTICAL"
	// ImageMode_LARGE_VERTICAL 大图竖图，宽高比0.56，大小1.5M以下，下限：720 & 1280，上限：1440 & 2560
	ImageMode_LARGE_VERTICAL ImageMode = "LARGE_VERTICAL"
	// ImageMode_UNION_SPLASH 穿山甲开屏图片，仅限穿山甲开屏使用，下限：1080 & 1920，上限：2160 & 3840，比例0.56
	ImageMode_UNION_SPLASH ImageMode = "UNION_SPLASH"
	// ImageMode_AWEME_LIVE_ROOM 直播画面类型
	ImageMode_AWEME_LIVE_ROOM ImageMode = "AWEME_LIVE_ROOM"
)
