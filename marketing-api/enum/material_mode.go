package enum

// MaterialMode 素材类型
type MaterialMode string

const (
	// MaterialMode_SMALL 小图，宽高比1.52，大小1.5M以下，下限：456 & 300，上限：1368 & 900
	MaterialMode_SMALL MaterialMode = "SMALL"
	// MaterialMode_LARGE 大图，横版大图宽高比1.78，大小1.5M以下，下限：1280 & 720，上限：2560 & 1440
	MaterialMode_LARGE MaterialMode = "LARGE"
	// MareitalMode_LARGE_VERTICAL 竖版图片
	MareitalMode_LARGE_VERTICAL MaterialMode = "LARGE_VERTICAL"
	// MaterialMode_SQUARE 商品卡方图
	MaterialMode_SQUARE MaterialMode = "SQUARE"
	// MaterialMode_VIDEO_LARGE 横版视频，封面图宽高比1.78（下限：1280 & 720，上限：2560 & 1440））,视频资源支持mp4、mpeg、3gp、avi文件格式，宽高比16:9，大小不超过1000M
	MaterialMode_VIDEO_LARGE MaterialMode = "VIDEO_LARGE"
	// MaterialMode_VIDEO_VERTICAL 竖版视频，封面图宽高比0.56（9:16），下限：720 & 1280，上限：1440 & 2560，视频资源支持mp4、mpeg、3gp、avi文件格式，大小不超过100M
	MaterialMode_VIDEO_VERTICAL MaterialMode = "VIDEO_VERTICAL"
	// MaterialMode_LARGE_VERTICAL 大图竖图，宽高比0.56，大小1.5M以下，下限：720 & 1280，上限：1440 & 2560
	MaterialMode_LARGE_VERTICAL MaterialMode = "LARGE_VERTICAL"
	// MaterialMode_UNION_SPLASH 穿山甲开屏图片，仅限穿山甲开屏使用，下限：1080 & 1920，上限：2160 & 3840，比例0.56
	MaterialMode_UNION_SPLASH MaterialMode = "UNION_SPLASH"
	// MaterialMode_AWEME_LIVE_ROOM 直播画面类型
	MaterialMode_AWEME_LIVE_ROOM MaterialMode = "AWEME_LIVE_ROOM"
	// MaterialMode_CAROUSEL 图文
	MaterialMode_CAROUSEL MaterialMode = "CAROUSEL"
	// MaterialMode_VIDEO 视频
	MaterialMode_VIDEO MaterialMode = "VIDEO"
)
