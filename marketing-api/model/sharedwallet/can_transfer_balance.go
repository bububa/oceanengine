package sharedwallet

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CanTransferBalanceRequest 资金共享-最大可转余额查询 API Request
type CanTransferBalanceRequest struct {
	// AccountID 鉴权账户
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 鉴权账户类型 可选值:
	// AD 巨量广告主
	// AGENT 代理
	// LOCAL 本地推广告主
	AccountType string `json:"account_type,omitempty"`
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// MainWalletID 大钱包id
	MainWalletID uint64 `json:"main_wallet_id,omitempty"`
	// SubWalletList 小钱包id列表
	SubWalletList []uint64 `json:"sub_wallet_list,omitempty"`
	// TransferDirection 转账方向，以小钱包视角确定 可选值:
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	TransferDirection enum.TransferDirection `json:"transfer_direction,omitempty"`
}

// Encode implements GetRequest interface
func (r CanTransferBalanceRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", r.AccountType)
	values.Set("biz_request_no", r.BizRequestNo)
	values.Set("main_wallet_id", strconv.FormatUint(r.MainWalletID, 10))
	values.Set("sub_wallet_list", string(util.JSONMarshal(r.SubWalletList)))
	values.Set("transfer_direction", string(r.TransferDirection))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CanTransferBalanceResponse 资金共享-最大可转余额查询 API Response
type CanTransferBalanceResponse struct {
	model.BaseResponse
	Data struct {
		// CanTransferDetailList 可转余额信息列表
		CanTransferDetailList []CanTransferDetail `json:"can_transfer_detail_list,omitempty"`
	} `json:"data,omitempty"`
}

// CanTransferDetail 可转余额信息
type CanTransferDetail struct {
	// RemitterWalletID 减款钱包id
	RemitterWalletID uint64 `json:"remitter_wallet_id,omitempty"`
	// NonBrandMaxTransferBalance 减款钱包非品牌资金最大可转出金额（单位：分）
	NonBrandMaxTransferBalance int64 `json:"non_brand_max_transfer_balance,omitempty"`
	// RemitterCapitalDetailList 减款钱包可转资金列表
	RemitterCapitalDetailList []RemitterCapitalDetail `json:"remitter_capital_detail_list,omitempty"`
	// PayeeTransferAmountDetailList 加款钱包可转余额信息列表
	PayeeTransferAmountDetailList []PayeeTransferAmountDetail `json:"payee_transfer_amount_detail_list,omitempty"`
}

// RemitterCapitalDetail 减款钱包可转资金
type RemitterCapitalDetail struct {
	// Platform 业务线 可选值:
	// AD 广告
	// AD_ALL 广告通用
	// BENDITUI 本地推
	// QIANCHUAN 千川
	// STAR 星图
	Platform string `json:"platform,omitempty"`
	// CapitalType 资金类型 可选值:
	// CREDIT_BIDDING 授信竞价
	// CREDIT_BRAND 授信品牌
	// CREDIT_GENERAL 授信通用
	// PREPAY_BIDDING 预付竞价
	// PREPAY_BRAND 预付品牌
	// PREPAY_GENERAL 预付通用
	CapitalType enum.CapitalType `json:"capital_type,omitempty"`
}

// PayeeTransferAmountDetail 加款钱包可转余额信息
type PayeeTransferAmountDetail struct {
	// PayeeWalletID 加款钱包id
	PayeeWalletID uint64 `json:"payee_wallet_id,omitempty"`
	// NonBrandMinTransferBalance 加款钱包非品牌资金最小转入金额（单位：分）
	NonBrandMinTransferBalance int64 `json:"non_brand_min_transfer_balance,omitempty"`
	// CapitalDetailList 加款钱包可转资金列表
	CapitalDetailList []CapitalDetail `json:"capital_detail_list,omitempty"`
}

// CapitalDetail 加款钱包可转资金
type CapitalDetail struct {
	// Platform 业务线 可选值:
	// AD 广告
	// AD_ALL 广告通用
	// BENDITUI 本地推
	// QIANCHUAN 千川
	// STAR 星图
	Platform string `json:"platform,omitempty"`
	// CapitalType 可转资金类型 可选值:
	// CREDIT_BIDDING 授信竞价
	// CREDIT_BRAND 授信品牌
	// CREDIT_GENERAL 授信通用
	// PREPAY_BIDDING 预付竞价
	// PREPAY_BRAND 预付品牌
	// PREPAY_GENERAL 预付通用
	CapitalType enum.CapitalType `json:"capital_type,omitempty"`
	// TransferBalance 可转资金金额(单位：分)
	TransferBalance int64 `json:"transfer_balance,omitempty"`
}
