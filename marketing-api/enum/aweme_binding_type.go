package enum

// AwemeBindType 抖音号授权类型
type AwemeBindType string

const (
	// AwemeBindType_OFFICIAL 官方
	AwemeBindType_OFFICIAL AwemeBindType = "OFFICIAL"
	// AwemeBindType_SELF 自运营
	AwemeBindType_SELF AwemeBindType = "SELF"
	// AwemeBindType_AWEME_COOPERATOR 合作达人
	AwemeBindType_AWEME_COOPERATOR AwemeBindType = "AWEME_COOPERATOR"
)
