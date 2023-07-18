package qianchuan

// InActiveRetargetingType 失效类型
type InActiveRetargetingType string

const (
	// InActiveRetargetingType_EXPIRE 人群包过期
	InActiveRetargetingType_EXPIRE InActiveRetargetingType = "EXPIRE"
	// InActiveRetargetingType_TAG_OFFLINE 人群包tag下线
	InActiveRetargetingType_TAG_OFFLINE InActiveRetargetingType = "TAG_OFFLINE"
	// InActiveRetargetingType_MANUAL_OFFLINE 精选人群包手动下线
	InActiveRetargetingType_MANUAL_OFFLINE InActiveRetargetingType = "MANUAL_OFFLINE"
)
