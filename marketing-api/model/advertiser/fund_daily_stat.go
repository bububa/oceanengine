package advertiser

import (
	"net/url"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type FundDailyStatRequest struct {
	AdvertiserID uint64    `json:"advertiser_id,omitempty"`
	StartDate    time.Time `json:"start_date,omitempty"` // 开始时间，格式YYYY-MM-DD，默认当前年份的1月1日
	EndDate      time.Time `json:"end_date,omitempty"`   // 结束时间，格式YYYY-MM-DD，默认为今天
	Page         int       `json:"page,omitempty"`       // 页码. 默认值: 1
	PageSize     int       `json:"page_size,omitempty"`  // 页面数据量. 默认值: 10
}

func (r FundDailyStatRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if !r.StartDate.IsZero() {
		values.Set("start_date", r.StartDate.Format("2006-01-02"))
	}
	if !r.EndDate.IsZero() {
		values.Set("end_date", r.EndDate.Format("2006-01-02"))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

type FundDailyStatResponse struct {
	model.BaseResponse
	Data *FundDailyStatResponseData `json:"data,omitempty"`
}

type FundDailyStatResponseData struct {
	List     []FundDailyStatResponseList `json:"list,omitempty"`
	PageInfo *model.PageInfo             `json:"page_info,omitempty"`
}

type FundDailyStatResponseList struct {
	AdvertiserID    uint64  `json:"advertiser_id,omitempty"`     // 广告主ID
	Date            string  `json:"date,omitempty"`              // 日期
	Balance         float64 `json:"balance,omitempty"`           // 日终结余(单位元）
	CashCost        float64 `json:"cash_cost,omitempty"`         // 现金支出(单位元)
	Cost            float64 `json:"cost,omitempty"`              // 总支出(单位元)
	Frozen          float64 `json:"frozen,omitempty"`            // 冻结(单位元)
	Income          float64 `json:"income,omitempty"`            // 总存入(单位元)
	RewardCost      float64 `json:"reward_cost,omitempty"`       // 赠款支出(单位元)
	ReturnGoodsCost float64 `json:"return_goods_cost,omitempty"` // 返货支出(单位元)
	TransferIn      float64 `json:"transfer_in,omitempty"`       // 总转入(单位元)
	TransferOut     float64 `json:"transfer_out,omitempty"`      // 总转出(单位元)
}
