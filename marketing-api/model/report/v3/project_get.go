package v3

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ProjectGetRequest 项目数据报表 API Request
type ProjectGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	EndDate time.Time `json:"end_date,omitempty"`
	// Fields 指定需要的指标名称
	Fields []string `json:"fields,omitempty"`
	// GroupBy 分组条件默认为 STAT_GROUP_BY_FIELD_STAT_TIME
	GroupBy enum.StatGroupBy `json:"group_by,omitempty"`
	// TimeGranularity 时间粒度, 默认值: STAT_TIME_GRANULARITY_DAILY
	TimeGranularity enum.StatTimeGranularity `json:"time_granularity,omitempty"`
	// OrderField 排序字段，所有的统计指标均可参与排序
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序方式；默认值: DESC；允许值: ASC, DESC
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页码；默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，即每页展示的数据量；默认值: 20；取值范围: 1-1000
	PageSize int `json:"page_size,omitempty"`
	// Filtering 过滤字段，json格式，支持字段如下
	Filtering *ProjectGetFilter `json:"filtering,omitempty"`
}

// ProjectGetFilter 数据报表过滤条件
type ProjectGetFilter struct {
	// ProjectID 项目ID，按照project_id过滤，最多支持100个
	ProjectID []uint64 `json:"project_id,omitempty"`
	// LandingType 推广目的，允许值：APP
	LandingType []enum.LandingType `json:"landing_type,omitempty"`
	// ExternalAction 转化目标，允许值：AD_CONVERT_TYPE_NOTIFY_DOWNLOAD
	// AD_CONVERT_TYPE_DOWNLOAD_FINISH
	// AD_CONVERT_TYPE_INSTALL_FINISH
	// AD_CONVERT_TYPE_ACTIVE
	// AD_CONVERT_TYPE_ACTIVE_REGISTER
	// AD_CONVERT_TYPE_GAME_ADDICTION
	// AD_CONVERT_TYPE_PAY
	ExternalAction []enum.AdConvertType `json:"external_action,omitempty"`
	// AppCode 首选位置，西瓜视频[1]
	// 抖音火山版:[3]
	// 抖音短视频[4]
	// 今日头条:[8]
	// 番茄小说:[26]
	// 穿山甲:[9]
	// Ohayoo精品游戏:[27]
	AppCode []int `json:"app_code,omitempty"`
	// Platform 平台，允许值：ANDROID 安卓、IOS 苹果、OTHER 其他
	Platform []enum.AudiencePlatform
	// PackageName 应用包名
	PackageName []string `json:"package_name,omitempty"`
	// DeliveryMode 投放模式筛选。允许值: MANUAL 手动、PROCEDURAL 自动
	DeliveryMode []string `json:"delivery_mode,omitempty"`
}

// Encode implement GetRequest interface
func (r ProjectGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("start_date", r.StartDate.Format("2006-01-02"))
	values.Set("end_date", r.EndDate.Format("2006-01-02"))
	if r.AdvertiserID > 0 {
		values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	if len(r.Fields) > 0 {
		fields, _ := json.Marshal(r.Fields)
		values.Set("fields", string(fields))
	}
	if r.GroupBy != "" {
		values.Set("group_by", string(r.GroupBy))
	}
	if r.TimeGranularity != "" {
		values.Set("time_granularity", string(r.TimeGranularity))
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
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ProjectGetResponse 项目数据报表 API Response
type ProjectGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *ProjectGetResult `json:"data,omitempty"`
}

// ProjectGetResult 返回值
type ProjectGetResult struct {
	// List 数据列表
	List []ProjectGetListItem `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// ProjectGetListItem 数据详情
type ProjectGetListItem struct {
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// ProjectName 项目名称
	ProjectName string `json:"project_name,omitempty"`
	// StatDatetime 数据时间
	StatDatetime string `json:"stat_datetime,omitempty"`
	// Data
	Data []Stat `json:"data,omitempty"`
}
