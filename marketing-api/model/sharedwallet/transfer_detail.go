package sharedwallet

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TransferDetailRequest 资金共享-查询转账单信息 API Request
type TransferDetailRequest struct {
	// AccountID 鉴权账户
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 鉴权账户类型 可选值:
	// AD 巨量广告主
	// AGENT 代理
	// LOCAL 本地推广告主
	AccountType string `json:"account_type,omitempty"`
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// TransferBizRequestNo 发起转账的幂等id
	TransferBizRequestNo string `json:"transfer_biz_request_no,omitempty"`
	// TransferSerial 转账单号，与transfer_biz_request_no两者传其一即可
	TransferSerial string `json:"transfer_serial,omitempty"`
}

// Encode implements GetRequest interface
func (r TransferDetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", r.AccountType)
	values.Set("biz_request_no", r.BizRequestNo)
	if r.TransferBizRequestNo != "" {
		values.Set("transfer_biz_request_no", r.TransferBizRequestNo)
	}
	if r.TransferSerial != "" {
		values.Set("transfer_serial", r.TransferSerial)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// TransferDetailResponse 资金共享-查询转账单信息 API Response
type TransferDetailResponse struct {
	model.BaseResponse
	// Data 转账单信息
	Data *TransferDetail `json:"data,omitempty"`
}

// TransferDetail 转账单信息
type TransferDetail struct {
	// TransferSerial 转账单号
	TransferSerial string `json:"transfer_serial,omitempty"`
	// BizRequestNo 幂等id
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// TransferDirection 转账方向(以目标账户视角确定) 可选值:
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	TransferDirection enum.TransferDirection `json:"transfer_direction,omitempty"`
	// TransferAmount 转账总金额(单位：分)
	TransferAmount int64 `json:"transfer_amount,omitempty"`
	// TransferStatus 转账总状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败(终态)
	// TRANSFER_ING 转账中
	// TRANSFER_PART 转账部分成功(终态)
	// TRANSFER_SUCCESS 转账成功(终态)
	TransferStatus enum.TransferStatus `json:"transfer_status,omitempty"`
	// TransferFinishTime 转账完成时间 yyyy-MM-dd HH:mm:ss
	TransferFinishTime string `json:"transfer_finish_time,omitempty"`
	// TransferCreateTime 转账创建时间yyyy-MM-dd HH:mm:ss
	TransferCreateTime string `json:"transfer_create_time,omitempty"`
	// TrnsferWalletRecordList 账户信息列表                                                   a
}

// TrnsferWalletRecord 账户信息                                                   a
type TransferWalletRecord struct {
	// MainWalletID 大钱包id
	MainWalletID uint64 `json:"main_wallet_id,omitempty"`
	// SubWalletID 小钱包id
	SubWalletID uint64 `json:"sub_wallet_id,omitempty"`
	// TransferAmount 转账金额(单位：分)
	TransferAmount int64 `json:"transfer_amount,omitempty"`
	// TransferStatus 转账总状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败(终态)
	// TRANSFER_ING 转账中
	// TRANSFER_PART 转账部分成功(终态)
	// TRANSFER_SUCCESS 转账成功(终态)
	TransferStatus enum.TransferStatus `json:"transfer_status,omitempty"`
	// TransferCapitalRecordList 转账资金类型列表
	TransferCapitalRecordList []TransferCapitalRecord `json:"transfer_capital_record_list,omitempty"`
}
