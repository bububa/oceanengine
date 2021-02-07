package enum

// 试玩素材方向
type PlayableOrientation string

const (
	ORIENTATION_BOTH             PlayableOrientation = "BOTH"             // 横竖屏均适配
	ORIENTATION_PORTRAIT         PlayableOrientation = "PORTRAIT"         // 只适配竖屏
	ORIENTATION_VALIDATE_SUCCESS PlayableOrientation = "VALIDATE_SUCCESS" // 校验通过
	ORIENTATION_LANDSCAPE        PlayableOrientation = "LANDSCAPE"        // 只适配横屏
)
