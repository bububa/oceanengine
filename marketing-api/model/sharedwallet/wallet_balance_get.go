package sharedwallet

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// WalletBalanceGetRequest 资金共享-批量查询钱包余额 API Request
type WalletBalanceGetRequest struct {
	// AccountID 账户ID，对于巨量AD、千川、本地推是广告主ID，对于代理商是代理商ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型 可选值:
	// AD 巨量广告主
	// AGENT 代理
	// LOCAL 本地推广告主
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// WalletIDList 资金共享钱包id列表，注意：传入的列表长度不可大于200
	WalletIDList []uint64 `json:"wallet_id_list,omitempty"`
	// WalletBalanceFilters 余额过滤条件
	WalletBalanceFilters *WalletBalanceFilter `json:"wallet_balance_filters,omitempty"`
}

// CapitalType 资金类型包括 预付、授信
type CapitalType string

const (
	// CapitalType_CREDIT 授信
	CapitalType_CREDIT CapitalType = "CREDIT"
	// CapitalType_NO_LIMIT 不限
	CapitalType_NO_LIMIT CapitalType = "NO_LIMIT"
	// CapitalType_PREPAY 预付
	CapitalType_PREPAY CapitalType = "PREPAY"
)

// DeliveryType 投放类型：竞价、品牌、通用
type DeliveryType string

const (
	// DeliveryType_BIDDING 竞价
	DeliveryType_BIDDING DeliveryType = "BIDDING"
	// DeliveryType_BRAND 品牌
	DeliveryType_BRAND DeliveryType = "BRAND"
	// DeliveryType_GENERAL 通用
	DeliveryType_GENERAL DeliveryType = "GENERAL"
	// DeliveryType_NO_LIMIT 不过滤
	DeliveryType_NO_LIMIT DeliveryType = "NO_LIMIT"
)

type WalletBalanceFilter struct {
	// AccountPlatformType 业务线：AD/千川/本地推;不填默认不过滤 可选值:
	// NO_LIMIT 不过滤
	// ONLY_AD 仅巨量AD专用
	// ONLY_AD_SHARED 仅巨量AD/千川/本地推专用
	// ONLY_ECP 仅巨量千川专用
	// ONLY_LOCAL 仅巨量本地推专用
	//  默认值: NO_LIMIT
	AccountPlatformType enum.AccountPlatformType `json:"account_platform_type,omitempty"`
	// CapitalType 资金类型包括 预付、授信; 不填默认不过滤资金类型 可选值:
	// CREDIT 授信
	// NO_LIMIT 不限
	// PREPAY 预付
	//  默认值: NO_LIMIT
	CapitalType CapitalType `json:"capital_type,omitempty"`
	// DeliveryType 投放类型：竞价、品牌、通用
	// BIDDING 竞价
	// BRAND 品牌
	// GENERAL 通用
	// NO_LIMIT 不过滤
	//  默认值: NO_LIMIT
	DeliveryType DeliveryType `json:"delivery_type,omitempty"`
}

// Encode implements GetRequest interface
func (r WalletBalanceGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	values.Set("wallet_id_list", string(util.JSONMarshal(r.WalletIDList)))
	if r.WalletBalanceFilters != nil {
		values.Set("wallet_balance_filters", string(util.JSONMarshal(r.WalletBalanceFilters)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// WalletBalanceGetResponse 资金共享-批量查询钱包余额 API Response
type WalletBalanceGetResponse struct {
	model.BaseResponse
	Data *WalletBalanceInfo `json:"data,omitempty"`
}

// WalletBalanceInfo 共享钱包余额信息
type WalletBalanceInfo struct {
	// WalletID 共享钱包ID
	WalletID uint64 `json:"wallet_id,omitempty"`
	// BasicBalanceInfo 常用余额信息
	BasicBalanceInfo *BasicBalanceInfo `json:"basic_balance_info,omitempty"`
	// GeneralBalanceInfo 通用余额信息
	GeneralBalanceInfo *GeneralBalanceInfo `json:"general_balance_info,omitempty"`
}

// BasicBalanceInfo 常用余额信息
type BasicBalanceInfo struct {
	// TotalBalance 总余额(单位元)
	TotalBalance float64 `json:"total_balance,omitempty"`
	// TotalValidBalance 总可用余额(单位元)
	TotalValidBalance float64 `json:"total_valid_balance,omitempty"`
	// TotalTransferableBalance 总可转余额(单位元)
	TotalTransferableBalance float64 `json:"total_transferable_balance,omitempty"`
	// NonGrantBalance 非赠款总余额(单位元)
	NonGrantBalance float64 `json:"non_grant_balance,omitempty"`
	// NonGrantValidBalance 非赠款可用余额(单位元)                         n
	NonGrantValidBalance float64 `json:"non_grant_valid_balance,omitempty"`
	// NonGrantTransferableBalance 非赠款可转余额(单位元)                                                        n
	NonGrantTransferableBalance float64 `json:"non_grant_transferable_balance,omitempty"`
	// PrepayBiddingBalance 预付竞价可用余额(单位元)
	PrepayBiddingBalance float64 `json:"prepay_bidding_balance,omitempty"`
	// PrepayBiddingValidBalance 预付竞价可用余额(单位元)
	PrepayBiddingValidBalance float64 `json:"prepay_bidding_valid_balance,omitempty"`
	// PrepayBrandBalance 预付品牌可用余额(单位元)
	PrepayBrandBalance float64 `json:"prepay_brand_balance,omitempty"`
	// PrepayBrandValidBalance 预付品牌可用余额(单位元)
	PreBrandValidBalance float64 `json:"prepay_brand_valid_balance,omitempty"`
	// PrepayGeneralBalance 预付通用可用余额(单位元)
	PrepayGeneralBalance float64 `json:"prepay_general_balance,omitempty"`
	// PrepayGeneralValidBalance 预付通用可用余额(单位元)
	PrepayGeneralValidBalance float64 `json:"prepay_general_valid_balance,omitempty"`
	// CreditBiddingBalance 授信竞价可用余额(单位元)
	CreditBiddingBalance float64 `json:"credit_bidding_balance,omitempty"`
	// CreditBiddingValidBalance 授信竞价可用余额(单位元)
	CreditBiddingValidBalance float64 `json:"credit_bidding_valid_balance,omitempty"`
	// CreditBrandBalance 授信品牌可用余额(单位元)
	CreditBrandBalance float64 `json:"credit_brand_balance,omitempty"`
	// CreditBrandValidBalance 授信品牌可用余额(单位元)
	CreditBrandValidBalance float64 `json:"credit_brand_valid_balance,omitempty"`
	// CreditGeneralBalance 授信通用可用余额(单位元)
	CreditGeneralBalance float64 `json:"credit_general_balance,omitempty"`
	// CreditGeneralValidBalance 授信通用可用余额(单位元)
	CreditGeneralValidBalance float64 `json:"credit_general_valid_balance,omitempty"`
}

// GeneralBalanceInfo 通用余额信息
type GeneralBalanceInfo struct {
	// AdOnlyBalanceInfo 巨量广告业务线余额信息
	AdOnlyBalanceInfo *BasicBalanceInfo `json:"ad_only_balance_info,omitempty"`
	// EcpOnlyBalanceInfo 巨量千川业务线余额信息
	EcpOnlyBalanceInfo *BasicBalanceInfo `json:"ecp_only_balance_info,omitempty"`
	// LocalOnlyBalanceInfo 巨量本地推业务线余额信息
	LocalOnlyBalanceInfo *BasicBalanceInfo `json:"local_only_balance_info,omitempty"`
	// AdSharedBalanceInfo 巨量AD/千川/本地推业务线余额信息
	AdSharedBalanceInfo *BasicBalanceInfo `json:"ad_shared_balance_info,omitempty"`
}
