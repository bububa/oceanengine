package agent

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FundTransferSeqCreateRequest 创建转账交易号（方舟） API Request
type FundTransferSeqCreateRequest struct {
	// AgentID 代理商账户ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// AccountID 广告主账户ID
	AccountID uint64 `json:"account_id,omitempty"`
	// TransferType 转账类型 可选值:
	// CASH_DEFAULT 非赠款金额
	// CREDIT_BIDDING 授信-竞价专用金额
	// CREDIT_BRAND 授信-品牌专用金额
	// CREDIT_GENERAL 授信-通用金额
	// GRANT_GENERAL 赠款金额
	// PREPAY_BIDDING 预付-竞价专用金额
	// PREPAY_BRAND 预付-品牌专用金额
	// PREPAY_GENERAL 预付-通用金额
	TransferType FundTransferType `json:"transfer_type,omitempty"`
	// Amount 转账金额，单位元，支持两位小数
	Amount float64 `json:"amount,omitempty"`
}

// Encode implement PostRequest interface
func (r FundTransferSeqCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// FundTransferSeqCreateResponse 创建转账交易号（方舟） API Response
type FundTransferSeqCreateResponse struct {
	model.BaseResponse
	Data struct {
		// TransferSeq 转账序列号，第二步提交操作需要
		TransferSeq string `json:"transfer_seq,omitempty"`
		// Status 状态码，0表示成功，其余表示失败
		Status int `json:"status,omitempty"`
		// StatusMessage 转账失败原因
		StatusMessage string `json:"status_message,omitempty"`
	} `json:"data,omitempty"`
}
