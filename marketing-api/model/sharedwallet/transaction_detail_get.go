package sharedwallet

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TransactionDetailGetRequest 资金共享-查询共享钱包流水明细 API Request
type TransactionDetailGetRequest struct {
	// AccountID 账户ID，对于巨量AD、千川、本地推是广告主ID，对于代理商是代理商ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型 可选值:
	// AD 巨量广告主
	// AGENT 代理
	// LOCAL 本地推广告主
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// SharedWalletID 共享钱包ID
	SharedWalletID uint64 `json:"shared_wallet_id,omitempty"`
	// StartDate 开始时间，格式YYYY-MM-DD
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式YYYY-MM-DD
	EndDate string `json:"end_date,omitempty"`
	// Page 页码；默认为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小；默认为10; 注意：page*page_size不可大于10000
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r TransactionDetailGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	values.Set("shared_wallet_id", strconv.FormatUint(r.SharedWalletID, 10))
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// TransactionDetailGetResponse 资金共享-查询共享钱包流水明细 API Response
type TransactionDetailGetResponse struct {
	model.BaseResponse
	Data *TransactionDetailGetResult `json:"data,omitempty"`
}

type TransactionDetailGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Results 流水明细结果
	Results []TransactionDetail `json:"results,omitempty"`
}

// TransactionDetail 流水明细
type TransactionDetail struct {
	// BusinessTime 交易时间,精确到秒
	BusinessTime string `json:"business_time,omitempty"`
	// TransactionSeq 交易流水id
	TransactionSeq uint64 `json:"transaction_seq,omitempty"`
	// BizType 流水类型(充值、退款等) 可选值:
	// ADJUST_DECREASE 调差减款
	// ADJUST_FREEZE 调差冻结/解冻
	// ADJUST_INCREASE 调差加款
	// CREDIT_RECHARGE 信控充值
	// CREDIT_REFUND 信控减款
	// INIT 钱包初始化
	// ORDER_PAY 订单支付
	// ORDER_WITHDRAW 订单还款
	// REFUND_FREEZE 冻结冻结/解冻
	// SHARED_CANCEL_RECHARGE 共享钱包充值撤销
	// SHARED_FROZEN_REFUND 共享钱包退款
	// SHARED_RECHARGE 共享钱包充值
	// TRANSFER 转账
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	BizType enum.TransactionBizType `json:"biz_type,omitempty"`
	// Payee 收款钱包ID
	Payee uint64 `json:"payee,omitempty"`
	// PayeeName 收款钱包名称
	PayeeName string `json:"payee_name,omitempty"`
	// Remitter 付款钱包ID
	Remitter uint64 `json:"remitter,omitempty"`
	// RemitterName 付款钱包名称
	RemitterName string `json:"remitter_name,omitempty"`
	// Amount 总金额(单位元)
	Amount float64 `json:"amount,omitempty"`
	// PrepayAmount 预付金额(单位元)
	PrepayAmount float64 `json:"prepay_amount,omitempty"`
	// CreditAmount 授信金额(单位元)
	CreditAmount float64 `json:"credit_amount,omitempty"`
}
