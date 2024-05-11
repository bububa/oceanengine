package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryTransferBalanceRequest 转账-查询账户转账余额（代理） API Request
type QueryTransferBalanceRequest struct {
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// AgentID 代理商账户id
	AgentID uint64 `json:"agent_id,omitempty"`
	// AccountIDList 查询账户id列表
	AccountIDList []uint64 `json:"account_id_list,omitempty"`
}

// Encode implements GetRequest interface
func (r QueryTransferBalanceRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("biz_request_no", r.BizRequestNo)
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	values.Set("account_id_list", string(util.JSONMarshal(r.AccountIDList)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QueryTransferBalanceResponse 转账-查询账户转账余额（代理） API Response
type QueryTransferBalanceResponse struct {
	model.BaseResponse
	Data struct {
		// AccountAmountDetailList 账户金额列表
		AccountAmountDetailList []AccountAmountDetail `json:"account_amount_detail_list,omitempty"`
	} `json:"data,omitempty"`
}

// AccountAmountDetail 账户金额
type AccountAmountDetail struct {
	// AccountID 账户id
	AccountID uint64 `json:"account_id,omitempty"`
	// CapitalDetailList 可转资金列表
	CapitalDetailList []CapitalDetail `json:"capital_detail_list,omitempty"`
	// DepositAmount 竞价消耗保证金金额(单位：分)
	DepositAmount int64 `json:"deposit_amount,omitempty"`
	// TotalTransferAmount 总可转金额(单位：分)
	TotalTransferAmount int64 `json:"total_transfer_amount,omitempty"`
}

// CapitalDetail 可转资金
type CapitalDetail struct {
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
