package enum

// Credential 密钥对label，枚举值：“primary”或“backup”
type Credential string

const (
	Credential_PRIMARY Credential = "primary"
	Credential_BACKUP  Credential = "backup"
)
