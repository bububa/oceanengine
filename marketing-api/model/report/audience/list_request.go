package audience

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

type ListRequest struct {
	AdvertiserID uint64     `json:"advertiser_id,omitempty"` // 广告主ID
	StartDate    time.Time  `json:"start_date,omitempty"`    // 起始日期，从0时起，格式2020-08-15; 默认15天前，即不指定起止时间获取最近15天数据
	EndDate      time.Time  `json:"end_date,omitempty"`      // 结束日期，至24时止，格式2020-08-29; 默认昨天，即不指定起止时间获取最近15天数据; 起始时间与结束时间之差小于15天，否则报错并提示"max time span is 15 days"
	Filtering    *Filtering `json:"filtering,omitempty"`     // 过滤条件
	Metrics      []string   `json:"metrics,omitempty"`       // 查询指标列表
	Page         int        `json:"page,omitempty"`          // 页码;默认值: 1
	PageSize     int        `json:"page_size,omitempty"`     // 页面大小，即每页展示的数据量，限制为1-100; 默认值: 20
}

func (r ListRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if !r.StartDate.IsZero() {
		values.Set("start_date", r.StartDate.Format("2006-01-02"))
	}
	if !r.EndDate.IsZero() {
		values.Set("end_date", r.EndDate.Format("2006-01-02"))
	}
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if r.Metrics != nil {
		metrics, _ := json.Marshal(r.Metrics)
		values.Set("metrics", string(metrics))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}
