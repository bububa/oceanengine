package sharedwallet

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// WalletInfoGetRequest 资金共享-批量查询钱包信息 API Request
type WalletInfoGetRequest struct {
	// AccountID 账户ID，对于巨量AD、千川、本地推是广告主ID，对于代理商是代理商ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型 可选值:
	// AD 巨量广告主
	// AGENT 代理
	// LOCAL 本地推广告主
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// WalletIDList 资金共享钱包id列表，注意：传入的列表长度不可大于200
	WalletIDList []uint64 `json:"wallet_id_list,omitempty"`
}

// Encode implments GetRequest interface
func (r WalletInfoGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	values.Set("wallet_id_list", string(util.JSONMarshal(r.WalletIDList)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// WalletInfoGetResponse 资金共享-批量查询钱包信息 API Response
type WalletInfoGetResponse struct {
	model.BaseResponse
	Data struct {
		// WalletInfo 钱包信息列表
		WalletInfo []WalletInfo `json:"wallet_info,omitempty"`
	} `json:"data,omitempty"`
}

// WalletInfo 钱包信息
type WalletInfo struct {
	// WalletID 共享钱包ID
	WalletID uint64 `json:"wallet_id,omitempty"`
	// WalletType 共享钱包类型 可选值:
	// MAIN_WALLET 共享钱包 即大钱包
	// SUB_CONSUME_WALLET 投放子钱包，可以挂多个广告主，可以投放
	// SUB_MANAGE_WALLET 管理子钱包，只能挂一个广告主（如代理商），不能投放
	WalletType enum.SharedWalletType `json:"wallet_type,omitempty"`
	// CommonWalletInfo 通用钱包信息
	CommonWalletInfo *CommonWalletInfo `json:"common_wallet_info,omitempty"`
	// MainWalletInfo 共享钱包(大钱包)信息
	MainWalletInfo *SharedMainWalletInfo `json:"main_wallet_info,omitempty"`
	// SubWalletInfo 投放钱包(小钱包)信息
	SubWalletInfo *SubWalletInfo `json:"sub_wallet_info,omitempty"`
}

// SharedMainWalletInfo 共享钱包(大钱包)信息
type SharedMainWalletInfo struct {
	// CompanyID 公司ID
	CompanyID uint64 `json:"company_id,omitempty"`
	// MainSharedRange 大钱包共享范围
	MainSharedRange *MainSharedRange `json:"main_shared_range,omitempty"`
}

// MainSharedRange 大钱包共享范围
type MainSharedRange struct {
	// Items
	Items []MainSharedRangeItem `json:"items,omitempty"`
}

type MainSharedRangeItem struct {
	// AccountPlatform 账户业务线 可选值:
	// AD 巨量广告
	// ECP 巨量千川
	// LOCAL_ADS 巨量本地推
	AccountPlatform string `json:"account_platform,omitempty"`
	// CustomerCategory 客户类型 可选值:
	// AGENT 代理商
	// DIRECT 直客
	// PARTNER 合作方
	// SELF_SERVICE 自助客户
	// SVIP_CUSTOMER VIP大型客户
	// VIRTUAL_CUS 虚客
	CustomerCategory []enum.CustomerCategory `json:"customer_category,omitempty"`
}

// SubWalletInfo 投放钱包(小钱包)信息
type SubWalletInfo struct {
	// MainWalletID 所属大钱包ID
	MainWalletID uint64 `json:"main_wallet_id,omitempty"`
	// AdvCnt 钱包下的 adv 数量
	AdvCnt int64 `json:"adv_cnt,omitempty"`
	// SubSharedRange 小钱包共享范围
	SubSharedRange *SubSharedRange `json:"sub_shared_range,omitempty"`
}

// SubSharedRange 小钱包共享范围
type SubSharedRange struct {
	// AccountPlatformList 账户业务线范围
	AccountPlatformList []string `json:"account_platform_list,omitempty"`
}
