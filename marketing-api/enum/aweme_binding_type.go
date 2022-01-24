package enum

// AwemeBindingType 抖音号授权类型
type AwemeBindingType string

const (
	// AwemeBindingType_OFFICIAL 官方
	AwemeBindingType_OFFICIAL AwemeBindingType = "OFFICIAL"
	// AwemeBindingType_SELF 自运营
	AwemeBindingType_SELF AwemeBindingType = "SELF"
	// AwemeBindingType_AWEME_COOPERATOR 合作达人
	AwemeBindingType_AWEME_COOPERATOR AwemeBindingType = "AWEME_COOPERATOR"
)
