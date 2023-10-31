package finance

import "github.com/bububa/oceanengine/marketing-api/enum/qianchuan"

// Transaction 财务流水
type Transaction struct {
	// Date 日期，格式 2021-04-05
	Date string `json:"date,omitempty"`
	// ViewDeliveryType 资金池类型，枚举值：
	// ALL 全部（默认值）
	// DEFAULT 通用
	// BRAND 品牌
	ViewDeliveryType qianchuan.ViewDeliveryType `json:"view_delivery_type,omitempty"`
	// DeductionCost 消返红包消耗
	DeductionCost float64 `json:"deduction_cost,omitempty"`
	// Cost 账户余额总消耗，不包括消返红包消耗和共享赠款消耗
	Cost float64 `json:"cost,omitempty"`
	// ShareCost 共享赠款消耗
	ShareCost float64 `json:"share_cost,omitempty"`
	// CashCost 非赠款消耗
	CashCost float64 `json:"cash_cost,omitempty"`
	// GrantCost 赠款消耗
	GrantCost float64 `json:"grant_cost,omitempty"`
	// Income 总存入
	Income float64 `json:"income,omitempty"`
	// TransferIn 总转入
	TransferIn float64 `json:"transfer_in,omitempty"`
	// TransferOut 总转出
	TransferOut float64 `json:"transfer_out,omitempty"`
	// CashBalance 非赠款余额
	CashBalance float64 `json:"cash_balance,omitempty"`
	// GrantBalance 赠款余额
	GrantBalance float64 `json:"grant_balance,omitempty"`
	// TotalBalance 总余额
	TotalBalance float64 `json:"total_balance,omitempty"`
	// QcAwemeCost 随心推消耗
	QcAwemeCost float64 `json:"qc_aweme_cost,omitempty"`
	// QcAwemeCashCost 随心推余额消耗
	QcAwemeCashCost float64 `json:"qc_aweme_cash_cost,omitempty"`
	// QcAwemeGrantCost 随心推赠款消耗
	QcAwemeGrantCost float64 `json:"qc_aweme_grant_cost,omitempty"`
}
