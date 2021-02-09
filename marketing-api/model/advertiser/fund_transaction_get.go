package advertiser

import (
	"net/url"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type FundTransactionGetRequest struct {
	AdvertiserID    uint64               `json:"advertiser_id,omitempty"`
	StartDate       time.Time            `json:"start_date,omitempty"`       // 开始时间，格式YYYY-MM-DD，默认当前年份的1月1日
	EndDate         time.Time            `json:"end_date,omitempty"`         // 结束时间，格式YYYY-MM-DD，默认为今天
	TransactionType enum.TransactionType `json:"transaction_type,omitempty"` // 流水类型
	Page            int                  `json:"page,omitempty"`             // 页码. 默认值: 1
	PageSize        int                  `json:"page_size,omitempty"`        // 页面数据量. 默认值: 10
}

func (r FundTransactionGetRequest) Encode() string {
	values := &url.Values{}
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
	return values.Encode()
}

type FundTransactionGetResponse struct {
	model.BaseResponse
	Data *FundTransactionGetResponseData `json:"data,omitempty"`
}

type FundTransactionGetResponseData struct {
	List []FundTransactionGetResponseList `json:"list,omitempty"`
}

type FundTransactionGetResponseList struct {
	AdvertiserID    uint64               `json:"advertiser_id,omitempty"`    // 广告主ID
	TransactionType enum.TransactionType `json:"transaction_type,omitempty"` // 流水类型
	CreateTime      string               `json:"create_time,omitempty"`      // 流水产生时间
	Amount          float64              `json:"amount,omitempty"`           // 交易总金额(单位元)
	Cash            float64              `json:"cash,omitempty"`             // 现金总金额(单位元)
	Frozen          float64              `json:"frozen,omitempty"`           // 冻结(单位元)
	Grant           float64              `json:"grant,omitempty"`            // 赠款总金额(单位元）
	ReturnGoods     float64              `json:"return_goods,omitempty"`     // 返货总金额(单位元)
	TransactionSeq  uint64               `json:"transaction_seq,omitempty"`  // 交易流水号
	Remitter        uint64               `json:"remitter,omitempty"`         // 付款方，即广告主id。
	Payee           uint64               `json:"payee,omitempty"`            // 收款方，即广告主id。
	Dealbase        float64              `json:"dealbase,omitempty"`         // 返点
}
