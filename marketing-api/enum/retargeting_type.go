package enum

// RetargetingType 定向人群包类型
type RetargetingType string

const (
	RETARGETING_INCLUDE RetargetingType = "RETARGETING_INCLUDE"
	RETARGETING_EXCLUDE RetargetingType = "RETARGETING_EXCLUDE"
	RETARGETING_NONE    RetargetingType = "RETARGETING_NONE"
)
