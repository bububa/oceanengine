package businessplatform

import "github.com/bububa/oceanengine/marketing-api/enum"

// Account 账户
type Account struct {
	// ID 账户id
	ID uint64 `json:"account_id,omitempty"`
	// Type 主体类型:
	// AD: 广告账户
	// QIANCHUAN:千川广告账户
	Type enum.AccountType `json:"account_type,omitempty"`
}
