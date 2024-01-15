package enum

// VideoHpVisibility 原生广告视频素材主页可见性设置
type VideoHpVisibility string

const (
	// VideoHpVisibility_ALWAYS_VISIBLE 主页始终可见
	VideoHpVisibility_ALWAYS_VISIBLE VideoHpVisibility = "ALWAYS_VISIBLE"
	// VideoHpVisibility_HIDE_AFTER_END_DATE 指定日期后隐藏
	VideoHpVisibility_HIDE_AFTER_END_DATE VideoHpVisibility = "HIDE_AFTER_END_DATE"
	// VideoHpVisibility_HIDE_AFTER_NO_PLAYBACK 无播放后隐藏
	VideoHpVisibility_HIDE_AFTER_NO_PLAYBACK VideoHpVisibility = "HIDE_AFTER_NO_PLAYBACK"
	// VideoHpVisibility_HIDE_VIDEO_ON_HP 主页隐藏（默认值）
	VideoHpVisibility_HIDE_VIDEO_ON_HP VideoHpVisibility = "HIDE_VIDEO_ON_HP"
)
