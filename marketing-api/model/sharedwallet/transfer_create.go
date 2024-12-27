package sharedwallet

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TransferCreateRequest 资金共享-发起转账 API Request
type TransferCreateRequest struct {
	// AccountID 鉴权账户
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 鉴权账户类型 可选值:
	// AD 巨量广告主
	// AGENT 代理
	// LOCAL 本地推广告主
	AccountType string `json:"account_type,omitempty"`
	// MainWalletID 大钱包id
	MainWalletID uint64 `json:"main_wallet_id,omitempty"`
	// TransferDirection 转账方向，以小钱包视角确定 可选值:
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	TransferDirection enum.TransferDirection `json:"transfer_direction,omitempty"`
	// TargetWalletDetailList 申请转账的小钱包列表
	TargetWalletDetailList []TargetWalletDetail `json:"target_wallet_detail_list,omitempty"`
}

// TargetWalletDetail 申请转账的小钱包
type TargetWalletDetail struct {
	// SubWallletID 申请转账的小钱包id
	SubWalletID uint64 `json:"sub_wallet_id,omitempty"`
	// TransferCapitalDetailList 申请转账资金列表
	TransferCapitalDetailList []CapitalDetail `json:"transfer_capital_detail_list,omitempty"`
}

// Encode implements PostRequest interface
func (r TransferCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TransferCreateResponse 资金共享-发起转账 API Response
type TransferCreateResponse struct {
	model.BaseResponse
	Data struct {
		// TransferSerial 转账单编号
		TransfrSerial string `json:"transfer_serial,omitempty"`
	} `json:"data,omitempty"`
}
