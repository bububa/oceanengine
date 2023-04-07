package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ExtendPackageCreateV2Request 创建应用分包（支持所有账户体系） API Request
type ExtendPackageCreateV2Request struct {
	// AccountID 账户id，accout_type类型对应账户ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型
	// 允许值：BP 巨量纵横组织、 AD 广告主账号、STAR 星图
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// PackageID 应用包ID，获取方法见接口文档【查询应用信息】
	PackageID string `json:"package_id,omitempty"`
	// Mode 分包模式，Auto自动生成渠道号Manual自定义渠道号信息
	Mode string `json:"mode,omitempty"`
	// ChannelCount 创建数量，（mode=Auto时需指定）单次分包取值范围1~100
	ChannelCount int `json:"channel_count,omitempty"`
	// ChannelList 自定义渠道号信息，（mode=Manual时需指定），一次调用，list的size取值范围1~100
	ChannelList []ExtendPackageChannel `json:"channel_list,omitempty"`
}

// Encode implement GetRequest interface
func (r ExtendPackageCreateV2Request) Encode() []byte {
	return util.JSONMarshal(r)
}
