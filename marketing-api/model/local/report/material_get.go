package report

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialGetRequest 获取素材数据 API Request
type MaterialGetRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// TimeGranularity 时间粒度，允许值：
	// TIME_GRANULARITY_DAILY 天维度（默认值）
	// TIME_GRANULARITY_HOURLY 小时维度
	// TIME_GRANULARITY_TOTAL 汇总
	TimeGranularity enum.TimeGranularity `json:"time_granularity,omitempty"`
	// StartDate 查询起始日期，格式：yyyy-mm-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询结束日期，格式：yyyy-mm-dd
	// 当time_granularity = TIME_GRANULARITY_DAILY/TIME_GRANULARITY_TOTAL时，时间跨度不能超过365天
	// 当time_granularity = TIME_GRANULARITY_HOURLY时，时间跨度不能超过7天
	EndDate string `json:"end_date,omitempty"`
	// OrderType 排序方式，允许值：
	// ASC 升序（默认值）
	// DESC 降序
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// OrderField 排序字段，允许值可参考应答返回数据指标
	OrderField string `json:"order_field,omitempty"`
	// Metrics 指标集，允许值可参考应答返回数据指标
	Metrics []string `json:"metrics,omitempty"`
	// Filtering 过滤器
	Filtering *MaterialGetFilter `json:"filtering,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：10（默认值）、20、50、100
	PageSize int `json:"page_size,omitempty"`
}

type MaterialGetFilter struct {
  // MaterialIDs 素材ID
  MaterialIDs []uint64 `json:"material_ids,omitempty"`
  // MaterialType 素材类型，允许值：
// CASURAL 图文
// VIDEO 视频
  MaterialType enum.MaterialMode `json:"material_type,omitempty"`
	// CampaignType 广告类型，允许值：
	// GENERAL 通投广告
	// SEARCHING 线索广告
	CampaignType local.AdType `json:"campaign_type,omitempty"`
}

// Encode implements GetRequest interface
func (r MaterialGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	if r.TimeGranularity != "" {
		values.Set("time_granularity", string(r.TimeGranularity))
	}
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	values.Set("metrics", string(util.JSONMarshal(r.Metrics)))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// MaterialGetResponse 获取素材数据 API Response
type MaterialGetResponse struct {
	model.BaseResponse
	Data *MaterialGetResult `json:"data,omitempty"`
}

type MaterialGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// MaterialList 
	MaterialList []Report `json:"material_list,omitempty"`
}
