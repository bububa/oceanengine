package wechat

// ShareMode 共享类型 可选值:
type ShareMode string

const (
	// BP_ALL_ACCOUNTS 组织内所有账户共享
	BP_ALL_ACCOUNTS ShareMode = "BP_ALL_ACCOUNTS"
	// COMPANY_ALL_ACCOUNTS 公司主体内所有账户共享
	COMPANY_ALL_ACCOUNTS ShareMode = "COMPANY_ALL_ACCOUNTS"
	// PART 指定账户共享
	PART ShareMode = "PART"
)

//	ShareType  可选值:
type ShareType string

const (
	// ACCOUNT 共享给账户
	ACCOUNT ShareType = "ACCOUNT"
	// GROUP 共享给组织
	GROUP ShareType = "GROUP"
)
