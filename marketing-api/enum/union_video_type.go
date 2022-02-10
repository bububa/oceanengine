package enum

// UnionVideoType 穿山甲视频创意类型
type UnionVideoType string

const (
	// UnionVideoType_ORIGINAL_VIDEO 原生视频
	UnionVideoType_ORIGINAL_VIDEO UnionVideoType = "ORIGINAL_VIDEO"
	// UnionVideoType_REWARDED_VIDEO 激励视频
	UnionVideoType_REWARDED_VIDEO UnionVideoType = "REWARDED_VIDEO"
	// UnionVideoType_SPLASH_VIDEO 穿山甲开屏
	UnionVideoType_SPLASH_VIDEO UnionVideoType = "SPLASH_VIDEO"
)
