package enum

// AppFileTag 文件标示
type AppFileTag string

const (
	// AppFileTag_DEFAULT 默认
	AppFileTag_DEFAULT AppFileTag = "DEFAULT"
	// AppFileTag_AGE_REMINDER 适龄提醒
	AppFileTag_AGE_REMINDER AppFileTag = "AGE_REMINDER"
	// AppFileTag_ANTI_ADDICTION_TIPS 防沉迷提示
	AppFileTag_ANTI_ADDICTION_TIPS AppFileTag = "ANTI_ADDICTION_TIPS"
	// AppFileTag_MATERIAL_SCREENSHOT 素材截图
	AppFileTag_MATERIAL_SCREENSHOT AppFileTag = "MATERIAL_SCREENSHOT"
	// AppFileTag_REAL_NAME_VERIFIED 实名注册
	AppFileTag_REAL_NAME_VERIFIED AppFileTag = "REAL_NAME_VERIFIED"
)
