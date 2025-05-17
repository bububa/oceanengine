package agent

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateTransferRequest 转账-发起转账（代理） API Request
type CreateTransferRequest struct {
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// AgentID 代理商账户id
	AgentID uint64 `json:"agent_id,omitempty"`
	// AccountID 锚定账户id，1:N的1
	AccountID uint64 `json:"account_id,omitempty"`
	// TargetAccountDetailList 目标账户列表，1:N的N，需要列表内账户类型相同，最多支持100个
	TargetAccountDetailList []TargetAccountDetail `json:"target_account_detail_list,omitempty"`
	// TransferDirection 转账方向，以目标账户视角确定 可选值:
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	TransferDirection enum.TransferDirection `json:"transfer_direction,omitempty"`
}

// TargetAccountDetail 目标账户
type TargetAccountDetail struct {
	// AccountID 目标账户id
	AccountID uint64 `json:"account_id,omitempty"`
	// TransferCapitalDetailList 锚定账户与目标账户转账资金列表
	TransferCapitalDetailList []CapitalDetail `json:"transfer_capital_detail_list,omitempty"`
}

// Encode implements PostRequest interface
func (r CreateTransferRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateTransferResponse 转账-发起转账（代理） API Response
type CreateTransferResponse struct {
	model.BaseResponse
	Data struct {
		// TransferSerial 转账单号
		TransferSerial string `json:"transfer_serial,omitempty"`
	} `json:"data,omitempty"`
}
