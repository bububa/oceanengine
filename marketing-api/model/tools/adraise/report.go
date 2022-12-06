package adraise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReportRequest 获取一键起量报告 API Request
type ReportRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// AdRaiseVersion 起量版本号，通过【获取起量版本信息】接口获取
	AdRaiseVersion string `json:"ad_raise_version,omitempty"`
	// StartTime 当前起量版本起量开始时间，格式：2021-03-31 16:00:00，通过【获取起量版本信息】接口获取
	StartTime string `json:"start_time,omitempty"`
	// EndTime 当前起量版本起量结束时间，格式：2021-03-31 17:00:00，结束时间必须大于开始时间，通过【获取起量版本信息】接口获取
	EndTime string `json:"end_time,omitempty"`
	// TimeDimension 报告时间维度; 允许值：SUM 获取总计报告，HOURLY 获取分时报告;默认值：SUM
	TimeDimension string `json:"time_dimension,omitempty"`
	// OrderField 排序指标，当 time_dimension 为HOURLY时可用，允许值：show、convert、 stat_time_hour 、;默认值：stat_time_hour 按照小时时间排序
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序类型，当time_dimension为HOURLY时可用，允许值：ASC 顺序、DESC：倒序; 默认值：ASC
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：1-100，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ReportRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	values.Set("ad_raise_version", r.AdRaiseVersion)
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	if r.TimeDimension != "" {
		values.Set("time_dimension", r.TimeDimension)
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
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

// ReportResponse A获取一键起量报告PI Response
type ReportResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *ReportResponseData `json:"data,omitempty"`
}

// ReportResponseData json 返回值
type ReportResponseData struct {
	// AdRaiseVersion 起量版本号
	AdRaiseVersion string `json:"ad_raise_version,omitempty"`
	// AdRaiseReport 起量报告
	AdRaiseReport []Report `json:"ad_raise_report,omitempty"`
	// PageInfo 页面信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// Report 起量报告
type Report struct {
	// StatDatetime 起量时间，当为时返回，格式：time_dimensionHOURLY2021-03-31 16:00 - 17:00
	StatDatetime string `json:"stat_datetime,omitempty"`
	// Cost 一键起量阶段产生消耗
	Cost int64 `json:"cost,omitempty"`
	// Show 一键起量阶段产生展示
	Show int64 `json:"show,omitempty"`
	// Click 一键起量阶段产生点击数
	Click int64 `json:"click,omitempty"`
	// Convert 一键起量阶段产生转换数
	Convert int64 `json:"convert,omitempty"`
	// Ctr CTR 一键起量期间点击率
	Ctr float64 `json:"ctr,omitempty"`
	// Cvr CVR 一键起量期间转化率
	Cvr float64 `json:"cvr,omitempty"`
}
