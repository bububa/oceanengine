package sharedwallet

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MainWalletGetRequest 资金共享-共享钱包信息查询 API Request
type MainWalletGetRequest struct {
	// AccountID 账户ID，对于巨量AD、千川、本地推是广告主ID，对于代理商是代理商ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型 可选值:
	// AD 巨量广告主
	// AGENT 代理
	// LOCAL 本地推广告主
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// MainWalletID 共享钱包ID即大钱包ID
	MainWalletID uint64 `json:"main_wallet_id,omitempty"`
}

// Encode implments GetRequest interface
func (r MainWalletGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	values.Set("main_wallet_id", strconv.FormatUint(r.MainWalletID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MainWalletGetResponse 资金共享-共享钱包信息查询 API Response
type MainWalletGetResponse struct {
	model.BaseResponse
	// Data 资金共享大钱包信息
	Data *MainWalletInfo `json:"data,omitempty"`
}

// CommonWalletInfo 通用钱包信息
type CommonWalletInfo struct {
	// WalletName 钱包名称
	WalletName string `json:"wallet_name,omitempty"`
	// WalletDescription 钱包描述
	WalletDescription string `json:"wallet_description,omitempty"`
	// WalletLabel 钱包标签
	WalletLabel string `json:"wallet_label,omitempty"`
	// SharedWalletStatus 钱包状态 可选值:
	// DEFUALT 默认测试
	SharedWalletStatus string `json:"shared_wallet_status,omitempty"`
	// CreateTime 钱包创建时间
	CreateTime string `json:"create_time,omitempty"`
}

// MainWalletInfo 资金共享大钱包信息
type MainWalletInfo struct {
	// MainWalletID 大钱包ID
	MainWalletID uint64 `json:"main_wallet_id,omitempty"`
	CommonWalletInfo
	// TotalBalance 钱包总余额(单位元)
	TotalBalance float64 `json:"total_balance,omitempty"`
	// UnallocatedBalance 钱包待分配余额(单位元)
	UnallocatedBalance *UnallocatedBalance `json:"unallocated_balance,omitempty"`
	// AllocatedBalance 钱包已分配余额(单位元)
	AllocatedBalance *AllocatedBalance `json:"allocated_balance,omitempty"`
}

// UnallocatedBalance 钱包待分配余额(单位元)
type UnallocatedBalance struct {
	// AdOnlyUnallocatedBalance 巨量广告业务线待分配余额(单位元)
	AdOnlyUnallocatedBalance *WalletBasicBalanceInfo `json:"ad_only_unallocated_balance,omitempty"`
	// EcpOnlyUnallocatedBalance 巨量千川业务线待分配余额(单位元)
	EcpOnlyUnallocatedBalance *WalletBasicBalanceInfo `json:"ecp_only_unallocated_balance,omitempty"`
	// LocalOnlyUnallocatedBalance 巨量本地推业务线待分配余额(单位元)
	LocalOnlyUnallocatedBalance *WalletBasicBalanceInfo `json:"local_only_unallocated_balance,omitempty"`
	// AdSharedUnallocatedBalance 巨量广告/千川/本地推业务线待分配余额(单位元)
	AdSharedUnallocatedBalance *WalletBasicBalanceInfo `json:"ad_shared_unallocated_balance,omitempty"`
}

// AllocatedBalance 钱包已分配余额(单位元)
type AllocatedBalance struct {
	// AdOnlyAllocatedBalance 巨量广告业务线待分配余额(单位元)
	AdOnlyAllocatedBalance *WalletBasicBalanceInfo `json:"ad_only_allocated_balance,omitempty"`
	// EcpOnlyAllocatedBalance 巨量千川业务线待分配余额(单位元)
	EcpOnlyAllocatedBalance *WalletBasicBalanceInfo `json:"ecp_only_allocated_balance,omitempty"`
	// LocalOnlyAllocatedBalance 巨量本地推业务线待分配余额(单位元)
	LocalOnlyAllocatedBalance *WalletBasicBalanceInfo `json:"local_only_allocated_balance,omitempty"`
	// AdSharedAllocatedBalance 巨量广告/千川/本地推业务线待分配余额(单位元)
	AdSharedAllocatedBalance *WalletBasicBalanceInfo `json:"ad_shared_allocated_balance,omitempty"`
}

type WalletBasicBalanceInfo struct {
	// AvailableBalance 可用余额(单位元)
	AvailableBalance *BasicBalanceInfo
	// UnavailableBalance 不可用余额(单位元)
	UnavailableBalance *BasicBalanceInfo
}
