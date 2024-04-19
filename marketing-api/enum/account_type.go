package enum

// AccountType 账户类型
type AccountType string

const (
	// AccountTYpe_AD 广告账户
	AccountType_AD AccountType = "AD"
	// AccountType_QIANCHUAN 千川广告账户
	AccountType_QIANCHUAN AccountType = "QIANCHUAN"
	// AccountType_BP 巨量纵横组织
	AccountType_BP AccountType = "BP"
	// AccountType_STAR 星图
	AccountType_STAR AccountType = "STAR"
	// AccountType_AGENT 巨量方舟
	AccountType_AGENT AccountType = "AGENT"
	// AccountType_LOCAL 本地推
	AccountType_LOCAL AccountType = "LOCAL"
	// AccountType_DOUPLUS DOU+
	AccountType_DOUPLUS AccountType = "DOU+"
)
