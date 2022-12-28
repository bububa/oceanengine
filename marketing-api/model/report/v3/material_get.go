package v3

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialGetRequest 素材数据报表 API Request
type MaterialGetRequest struct {
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
	Filtering *MaterialGetFilter `json:"filtering,omitempty"`
}

// MaterialGetFilter 数据报表过滤条件
type MaterialGetFilter struct {
	// ProjectID 项目ID，最多入参100个
	ProjectID []uint64 `json:"project_id,omitempty"`
	// PromotionID 广告ID，按照promotion_id过滤，最多支持100个
	PromotionID []uint64 `json:"promotion_id,omitempty"`
	// VideoMaterial 视频素材信息
	VideoMaterial *Material `json:"video_material,omitempty"`
	// ImageMaterial 图片素材信息
	ImageMaterial *Material `json:"image_material,omitempty"`
	// TitleMaterial 标题素材信息
	TitleMaterial *Material `json:"title_material,omitempty"`
	// DeliveryMode 投放模式筛选。允许值: MANUAL 手动、PROCEDURAL 自动
	DeliveryMode []string `json:"delivery_mode,omitempty"`
}

// Material 素材信息
type Material struct {
	// Mid 素材id
	Mid uint64 `json:"mid,omitempty"`
	// ImageMode 视频素材类型，允许值：CREATIVE_IMAGE_MODE_VIDEO 横版视频、CREATIVE_IMAGE_MODE_VIDEO_VERTICAL 竖版视频
	ImageMode []enum.ImageMode `json:"image_mode,omitempty"`
	// AppCode 首选位置，允许值：
	// 西瓜视频[1]
	// 抖音火山版:[3]
	// 抖音短视频[4]
	// 今日头条:[8]
	// 番茄小说:[26]
	// 穿山甲:[9]
	// Ohayoo精品游戏:[27]
	AppCode []int `json:"app_code"`
}

// Encode implement GetRequest interface
func (r MaterialGetRequest) Encode() string {
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

// MaterialGetResponse 素材数据报表 API Response
type MaterialGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *MaterialGetResult `json:"data,omitempty"`
}

// MaterialGetResult 返回值
type MaterialGetResult struct {
	// List 数据列表
	List []MaterialGetListItem `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// MaterialGetListItem 数据详情
type MaterialGetListItem struct {
	// MaterialID 素材ID
	MaterialID uint64 `json:"material_id,omitempty"`
	// ImageMode 素材类型
	ImageMode string `json:"image_mode,omitempty"`
	// StatDatetime 数据时间
	StatDatetime string `json:"stat_datetime,omitempty"`
	// Data
	Data []Stat `json:"data,omitempty"`
}
