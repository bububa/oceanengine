package customercenter

// AdvertiserType 广告主类型
type AdvertiserType string

const (
	// AdvertiserType_DOU DOU+类广告主账号
	AdvertiserType_DOU AdvertiserType = "DOU+"
	// AdvertiserType_NORMAL 普通广告主帐号
	AdvertiserType_NORMAL AdvertiserType = "NORMAL"
	// AdvertiserType_NORMAL_AD 普通广告主帐号
	AdvertiserType_NORMAL_AD AdvertiserType = "NORMAL_AD"
	// AdvertiserType_LOCAL 本地推
	AdvertiserType_LOCAL AdvertiserType = "LOCAL"
)

// Advertiser 广告主
type Advertiser struct {
	// ID 广告主id
	ID uint64 `json:"advertiser_id,omitempty"`
	// Name 广告主名称（一个对象只会返回企业号和广告主其中一种）
	Name string `json:"advertiser_name,omitempty"`
	// AdvertiserType 广告主类型
	// 枚举值：DOU+ DOU+类广告主账号、NORMAL普通广告主帐号
	// DOU+类广告主账号不支持任何调用接口操作
	Type AdvertiserType `json:"advertiser_type,omitempty"`
	// EDouyinID 企业号id（一个对象只会返回企业号和广告主其中一种）
	EDouyinID string `json:"e_douyin_id,omitempty"`
	// EDouyinName 企业号名称（一个对象只会返回企业号和广告主其中一种）
	EDouyinName string `json:"e_douyin_name,omitempty"`
	// AccountId 账户id
	AccountId uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型
	// NORMAL_AD：普通广告主
	// LOCAL：本地推
	// DOU+：dou+
	AccountType AdvertiserType `json:"account_type,omitempty"`
	// AccountName 账户名称
	AccountName string `json:"account_name,omitempty"`
	// IsEnabled 账户状态，true为审核通过，false为审核不通过，注：账户审核通过方可进行转账
	IsEnabled bool `json:"is_enabled,omitempty"`
}
