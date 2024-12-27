package enum

// SharedWalletType 共享钱包类型
type SharedWalletType string

const (
	// MAIN_WALLET 共享钱包 即大钱包
	MAIN_WALLET SharedWalletType = "MAIN_WALLET"
	// SUB_CONSUME_WALLET 投放子钱包，可以挂多个广告主，可以投放
	SUB_CONSUME_WALLET SharedWalletType = "SUB_CONSUME_WALLET"
	// SUB_MANAGE_WALLET 管理子钱包，只能挂一个广告主（如代理商），不能投放
	SUB_MANAGE_WALLET SharedWalletType = "SUB_MANAGE_WALLET"
)
