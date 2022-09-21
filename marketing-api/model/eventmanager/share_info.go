package eventmanager

import "github.com/bububa/oceanengine/marketing-api/enum"

type ShareMode string

const (
	// ShareMode_ALL 组织下某业务线账
	ShareMode_ALL ShareMode = "ALL"
	// ShareMode_PART 指定账户共享
	ShareMode_PART ShareMode = "PART"
)

// ShareInfo 共享账户ID（adv+bpid+枚举值）集合
type ShareInfo struct {
	// ShareMode 共享类型
	ShareMode ShareMode `json:"share_mode,omitempty"`
	// AccountInfo 共享账号信息
	AccountInfo *ShareAccount `json:"account_info,omitempty"`
	// AllAccountType 业务线
	AllAccountType enum.AccountType `json:"all_account_type,omitempty"`
}

// ShareAccount 共享账号信息
type ShareAccount struct {
	// AccountID 账号id
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 业务线
	AccountType enum.AccountType `json:"account_type,omitempty"`
}
