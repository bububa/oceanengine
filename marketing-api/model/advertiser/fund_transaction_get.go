package advertiser

import (
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FundTransactionGetRequest 查询账号流水明细 API Request
type FundTransactionGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 开始时间，格式YYYY-MM-DD，默认当前年份的1月1日
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate 结束时间，格式YYYY-MM-DD，默认为今天
	EndDate time.Time `json:"end_date,omitempty"`
	// TransactionType 流水类型
	TransactionType enum.TransactionType `json:"transaction_type,omitempty"`
	// Page 页码. 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量. 默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r FundTransactionGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_date", r.StartDate.Format("2006-01-02"))
	values.Set("end_date", r.EndDate.Format("2006-01-02"))
	values.Set("transaction_type", string(r.TransactionType))
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

// FundTransactionGetResponse 查询账号流水明细 API Response
type FundTransactionGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *FundTransactionGetResponseData `json:"data,omitempty"`
}

// FundTransactionGetResponseData json返回值
type FundTransactionGetResponseData struct {
	// List 明细list
	List []FundTransactionGetResponseList `json:"list,omitempty"`
}

// FundTransactionGetResponseList 明细
type FundTransactionGetResponseList struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TransactionType 流水类型
	TransactionType enum.TransactionType `json:"transaction_type,omitempty"`
	// CreateTime 流水产生时间
	CreateTime string `json:"create_time,omitempty"`
	// Amount 交易总金额(单位元)
	Amount float64 `json:"amount,omitempty"`
	// Cash 现金总金额(单位元)
	Cash float64 `json:"cash,omitempty"`
	// Frozen 冻结(单位元)
	Frozen float64 `json:"frozen,omitempty"`
	// Grant 赠款总金额(单位元）
	Grant float64 `json:"grant,omitempty"`
	// ReturnGoods 返货总金额(单位元)
	ReturnGoods float64 `json:"return_goods,omitempty"`
	// TransactionSeq 交易流水号
	TransactionSeq uint64 `json:"transaction_seq,omitempty"`
	// Remitter 付款方，即广告主id。
	Remitter uint64 `json:"remitter,omitempty"`
	// Payee 收款方，即广告主id。
	Payee uint64 `json:"payee,omitempty"`
	// Dealbase 返点
	Dealbase float64 `json:"dealbase,omitempty"`
}
