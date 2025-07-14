package enum

// PlayMaterialImageMode 素材类型
type PlayMaterialImageMode string

const (
	// IMAGE_MODE_TRIAL_PLAY 横屏试玩素材
	IMAGE_MODE_TRIAL_PLAY PlayMaterialImageMode = "IMAGE_MODE_TRIAL_PLAY"
	// IMAGE_MODE_TRIAL_PLAY_VERTICAL 竖屏试玩素材
	IMAGE_MODE_TRIAL_PLAY_VERTICAL PlayMaterialImageMode = "IMAGE_MODE_TRIAL_PLAY_VERTICAL"
	// IMAGE_MODE_INSTANT_PLAY 直玩素材
	IMAGE_MODE_INSTANT_PLAY PlayMaterialImageMode = "IMAGE_MODE_INSTANT_PLAY"
)
