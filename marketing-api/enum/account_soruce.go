package enum

// 账户类型
type AccountSource string

const (
	NORMAL_ADVERTISER AccountSource = "NORMAL_ADVERTISER" // 普通广告主
	LUBAN_ACCOUNT     AccountSource = "LUBAN_ACCOUNT"     // 鲁班广告主
)
