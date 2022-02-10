package enum

// PlayableOrientation 试玩素材方向
type PlayableOrientation string

const (
	// ORIENTATION_BOTH 横竖屏均适配
	ORIENTATION_BOTH PlayableOrientation = "BOTH"
	// ORIENTATION_PORTRAIT 只适配竖屏
	ORIENTATION_PORTRAIT PlayableOrientation = "PORTRAIT"
	// ORIENTATION_VALIDATE_SUCCESS 校验通过
	ORIENTATION_VALIDATE_SUCCESS PlayableOrientation = "VALIDATE_SUCCESS"
	// ORIENTATION_LANDSCAPE 只适配横屏
	ORIENTATION_LANDSCAPE PlayableOrientation = "LANDSCAPE"
)
