package enum

// RtaTargetType 生效维度
type RtaTargetType string

const (
	// RtaTargetType_ADV 广告账户
	RtaTargetType_ADV RtaTargetType = "ADV"
	// RtaTargetType_CAMPAIGN 广告组
	RtaTargetType_CAMPAIGN RtaTargetType = "CAMPAIGN"
	// RtaTargetType_PROJECT 项目（体验版）
	RtaTargetType_PROJECT RtaTargetType = "PROJECT"
)
