package qianchuan

// AutoExtendTarget 可开放定向
type AutoExtendTarget string

const (
	// AutoExtendTarget_REGION 地域
	AutoExtendTarget_REGION AutoExtendTarget = "REGION"
	// AutoExtendTarget_GENDER 性别
	AutoExtendTarget_GENDER AutoExtendTarget = "GENDER"
	// AutoExtendTarget_AGE 年龄
	AutoExtendTarget_AGE AutoExtendTarget = "AGE"
	// AutoExtendTarget_INTEREST_TAG 兴趣关键词
	AutoExtendTarget_INTEREST_TAG AutoExtendTarget = "INTEREST_TAG"
	// AutoExtendTarget_INTEREST_ACTION 行为兴趣
	AutoExtendTarget_INTEREST_ACTION AutoExtendTarget = "INTEREST_ACTION"
)
