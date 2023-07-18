package qianchuan

// AudienceAccountType 推送账户
type AudienceAccountType string

const (
	// AudienceAccountType_LOCAL 本账户
	AudienceAccountType_LOCAL AudienceAccountType = "LOCAL"
	// AudienceAccountType_NO_LOCAL 其他账户
	AudienceAccountType_NO_LOCAL AudienceAccountType = "NO_LOCAL"
)
