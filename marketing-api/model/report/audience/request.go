package audience

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// Request 受众分析数据报表 API Request
type Request struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 起始日期，从0时起，格式2020-08-15; 默认15天前，即不指定起止时间获取最近15天数据
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate 结束日期，至24时止，格式2020-08-29; 默认昨天，即不指定起止时间获取最近15天数据; 起始时间与结束时间之差小于15天，否则报错并提示"max time span is 15 days"
	EndDate time.Time `json:"end_date,omitempty"`
	// Page 页码;默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，即每页展示的数据量，限制为1-100; 默认值: 20
	PageSize int `json:"page_size,omitempty"`
	// IDType 查询ID类型
	IDType enum.AudienceStatIDType `json:"id_type,omitempty"`
	// IDs 查询ID列表，长度1-100; id_type为AUDIENCE_STAT_ID_TYPE_ADVERTISER时，选填；其他类型，必填
	IDs     []uint64 `json:"ids,omitempty"`
	Metrics []string `json:"metrics,omitempty"` // Metrics 查询指标列表
}

// Encode implement GetRequest interface
func (r Request) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_date", r.StartDate.Format("2006-01-02"))
	values.Set("end_date", r.EndDate.Format("2006-01-02"))
	values.Set("id_type", string(r.IDType))
	if r.IDs != nil {
		ids, _ := json.Marshal(r.IDs)
		values.Set("ids", string(ids))
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
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
