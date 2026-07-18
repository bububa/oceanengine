package enum

// StdMaterialType 标准投放项目素材类型
type StdMaterialType string

const (
	// StdMaterialType_VIDEO 视频
	StdMaterialType_VIDEO StdMaterialType = "VIDEO"
	// StdMaterialType_TITLE 标题
	StdMaterialType_TITLE StdMaterialType = "TITLE"
	// StdMaterialType_IMAGE 图片
	StdMaterialType_IMAGE StdMaterialType = "IMAGE"
	// StdMaterialType_CAROUSEL 图文
	StdMaterialType_CAROUSEL StdMaterialType = "CAROUSEL"
	// StdMaterialType_TRIAL_PLAY_MATERIAL 试玩素材
	StdMaterialType_TRIAL_PLAY_MATERIAL StdMaterialType = "TRIAL_PLAY_MATERIAL"
	// StdMaterialType_INSTANT_PLAY_MATERIAL 直玩素材
	StdMaterialType_INSTANT_PLAY_MATERIAL StdMaterialType = "INSTANT_PLAY_MATERIAL"
	// StdMaterialType_MINI_PROGRAM_INFO 字节小程序&小游戏信息
	StdMaterialType_MINI_PROGRAM_INFO StdMaterialType = "MINI_PROGRAM_INFO"
	// StdMaterialType_OTHER 其他素材类型（包括产品信息、短剧专辑链接、落地页、行动号召等）
	StdMaterialType_OTHER StdMaterialType = "OTHER"
)
