package sharedwallet

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DailyStatGetRequest 资金共享-查询共享钱包日流水 API Request
type DailyStatGetRequest struct {
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
func (r DailyStatGetRequest) Encode() string {
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

// DailyStatGetResponse 资金共享-查询共享钱包日流水 API Response
type DailyStatGetResponse struct {
	model.BaseResponse
	Data *DailyStatGetResult `json:"data,omitempty"`
}

type DailyStatGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Results 流水明细结果
	Results []DailyStat `json:"results,omitempty"`
}

// DailyStat 日流水列表信息
type DailyStat struct {
	// TransactionDate 日期,精确到天，格式YYYY-MM-DD
	TransactionDate string `json:"transaction_date,omitempty"`
	// SharedWalletID 共享钱包ID
	SharedWalletID uint64 `json:"shared_wallet_id,omitempty"`
	// SharedWalletName 共享钱包名称
	SharedWalletName string `json:"shared_wallet_name,omitempty"`
	// WalletType 共享钱包类型 可选值:
	// MAIN_WALLET 共享钱包 即大钱包
	// SUB_CONSUME_WALLET 投放子钱包，可以挂多个广告主，可以投放
	// SUB_MANAGE_WALLET 管理子钱包，只能挂一个广告主（如代理商），不能投放
	WalletType enum.SharedWalletType `json:"wallet_type,omitempty"`
	// Balance 日终结余(单位元）
	Balance float64 `json:"balance,omitempty"`
	// PrepayBalance 预付日终结余(单位元）
	PrepayBalance float64 `json:"prepay_balance,omitempty"`
	// CreditBalance 授信日终结余(单位元）
	CreditBalance float64 `json:"credit_balance,omitempty"`
	// Incomes 总存入(单位元)
	Incomes float64 `json:"incomes,omitempty"`
	// Expenses 总支出(单位元)
	Expenses float64 `json:"expenses,omitempty"`
	// Cost 总消耗(单位元)
	Cost float64 `json:"cost,omitempty"`
}
