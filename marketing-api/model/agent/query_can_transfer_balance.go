package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryCanTransferBalanceRequest 转账-获取最大可转余额（代理） API Request
type QueryCanTransferBalanceRequest struct {
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// AgentID 代理商账户id
	AgentID uint64 `json:"agent_id,omitempty"`
	// AccountID 锚定账户id，1:N的1
	AccountID uint64 `json:"account_id,omitempty"`
	// TargetAccountIDList 目标账户id列表，1:N的N，需要列表内账户类型相同，最多支持100个
	TargetAccountIDList []uint64 `json:"target_account_id_list,omitempty"`
	// TransferDirection 转账方向，以目标账户视角确定 可选值:
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	TransferDirection enum.TransferDirection `json:"transfer_direction,omitempty"`
}

// Encode implements GetRequest interface
func (r QueryCanTransferBalanceRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("biz_request_no", r.BizRequestNo)
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("target_account_id_list", string(util.JSONMarshal(r.TargetAccountIDList)))
	values.Set("transfer_direction", string(r.TransferDirection))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QueryCanTransferBalanceResponse 转账-获取最大可转余额（代理） API Response
type QueryCanTransferBalanceResponse struct {
	model.BaseResponse
	Data struct {
		// CanTransferDetailList 可转余额信息列表
		CanTransferDetailList []CanTransferDetail `json:"can_transfer_detail_list,omitempty"`
	} `json:"data,omitempty"`
}

// CanTransferDetail 可转余额信息
type CanTransferDetail struct {
	// RemitterAccountID 减款方账户id
	RemitterAccountID uint64 `json:"remitter_account_id,omitempty"`
	// CapitalDetailList 减款方可转资金列表
	CapitalDetailList []CapitalDetail `json:"capital_detail_list,omitempty"`
	// PayeeTransferDetailAmountList 加款方可转余额信息列表
	PayeeTransferDetailAmountList []PayeeTransferDetailAmount `json:"payee_transfer_detail_amount_list,omitempty"`
}

// PayeeTransferDetailAmount 加款方可转余额信息
type PayeeTransferDetailAmount struct {
	// PayeeAccountID 加款方账户id
	PayeeAccountID uint64 `json:"payee_account_id,omitempty"`
	// CapitalDetailList 加款方可转资金列表
	CapitalDetailList []CapitalDetail `json:"capital_detail_list,omitempty"`
}
