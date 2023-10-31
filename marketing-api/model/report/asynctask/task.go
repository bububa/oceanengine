package asynctask

import "github.com/bububa/oceanengine/marketing-api/enum"

// Task 任务
type Task struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TaskID 任务 id
	TaskID uint64 `json:"task_id,omitempty"`
	// TaskName 任务名称。最大长度 25 个字符，不能为空字符。
	TaskName string `json:"task_name,omitempty"`
	// TaskType 任务类型
	TaskType enum.AsyncTaskType `json:"task_type,omitempty"`
	// TaskParams 任务参数
	TaskParams *TaskParams `json:"task_params,omitempty"`
	// CreateTime 任务创建时间。格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
	// TaskStatus 任务状态。详见【附录-异步任务状态】
	TaskStatus enum.AsyncTaskStatus `json:"task_status,omitempty"`
	// DataTopic 可选值:
	// BASIC_DATA 广告基础数据
	// BIDWORD_DATA 关键词数据
	// MATERIAL_DATA 素材数据
	// ONE_KEY_BOOST_DATA 一键起量数据
	// PRODUCT_DATA 产品数据
	// QUERY_DATA 搜索词数据
	DataTopic string `json:"data_topic,omitempty"`
	// ErrMsg 任务失败时，这个字段表示失败原因
	ErrMsg string `json:"err_msg,omitempty"`
	// FileSize 任务完成时，这个字段表示生成文件的大小
	FileSize int64 `json:"file_size,omitempty"`
}

// TaskParams 任务的参数
type TaskParams struct {
	// StartDate 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期。
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期。如果分组条件中不包括时间类型，那么时间区间不能超过一年 366 天。
	// 如果分组条件包含STAT_GROUP_BY_TIME_MONTH，那么时间区间不能超过一年 366 天。
	// 如果分组条件包含STAT_GROUP_BY_TIME_WEEK，那么时间区间不能超过三个月 93 天。
	// 如果分组条件包括STAT_GROUP_BY_TIME_DAY，那么时间区间不能超过一个月 31 天。
	EndDate string `json:"end_date,omitempty"`
	// GroupBy 分组条件，【附录：分组条件】
	// 假设一次查询中共有m个campaign_id，n天数据，有以下三种情况：
	// ①group_by=["STAT_GROUP_BY_TIME_DAY"]表示数据按照天粒度聚合，即本次返回最多为n个数据，表示将m个campaign_id的数据按天各自累计。
	// ②group_by=["STAT_GROUP_BY_CAMPAIGN_ID"]表示按照campaign_id聚合，本次返回最多返回m条数据，即按照m个campaign_id各自累加.
	// ③group_by=["STAT_GROUP_BY_CAMPAIGN_ID", "STAT_GROUP_BY_TIME_DAY"]表示按照时间和id同时聚合，最多返回m * n个数据，返回数据中会同时存在id和时间。
	// 同理如果group_by=["STAT_GROUP_BY_TIME_DAY", "STAT_GROUP_BY_CAMPAIGN_ID","STAT_GROUP_BY_INVENTORY"]表示按照天、广告组id和广告位（或者其他细分维度）同时聚合。
	GroupBy []enum.StatGroupBy `json:"group_by,omitempty"`
	// OrderField 排序字段，所有的统计指标均可参与排序
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序方式
	// 默认值: DESC
	// 允许值: "ASC", "DESC"
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Fields 指定需要的消耗相关指标名称。注意：
	// 1. 如果没有指定，那么只返回支持的默认指标名称；
	// 2. 对于不同的分组条件，支持不同的指标；具体说明可以参考：异步任务报表支持指标
	// 目前还不支持返回如下字段：
	// "attribution_convert_cost",
	// "attribution_deep_convert_cost",
	// "attribution_next_day_open_cnt",
	// "attribution_next_day_open_cost",
	// "attribution_next_day_open_rate"
	Fields []string `json:"fields,omitempty"`
	// Filtering 过滤字段
	Filtering *TaskParamsFilter `json:"filtering,omitempty"`
}

// TaskParamsFilter 过滤字段
type TaskParamsFilter struct {
	// CampaignID 按照campaign_id过滤
	CampaignID []uint64 `json:"campaign_id,omitempty"`
	// AdID 按照ad_id过滤
	AdID []uint64 `json:"ad_id,omitempty"`
	// CreativeID 按照creative_id过滤
	CreativeID []uint64 `json:"creative_id,omitempty"`
	// LandingType 按照广告组推广目的过滤
	LandingType []enum.LandingType `json:"landing_type,omitempty"`
	// Pricing 按照出价方式过滤，详见【附录-计划出价类型】
	Pricing []enum.PricingType `json:"pricing,omitempty"`
	// InventoryType 按照广告首选位置过滤，详见【附录-首选投放位置】
	InventoryType []enum.StatInventoryType `json:"inventory_type,omitempty"`
	// CreativeMaterialMode 按照创意类型过滤，STATIC_ASSEMBLE 表示程序化创意，INTERVENE表示自定义创意。
	CreativeMaterialMode []enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// ConvertType 按照转化类型过滤，只支持 dpa 报表。详见【附录-转化类型】
	ConvertType []enum.AdConvertType `json:"convert_type,omitempty"`
	// Platform 按照平台类型过滤，只支持 dpa 报表。详见【附录-受众平台类型】
	Platform []enum.AudiencePlatform `json:"platform,omitempty"`
	// Bidword 按照搜索词过滤，只支持关键词/搜索词报表，关键词可通过【搜索广告-获取关键词】接口进行获取
	Bidword []string `json:"bidword,omitempty"`
	// Query 按照搜索词过滤。搜索词表示用户是通过【搜索词】进行检索到相关广告，可用于分析怎样的搜索词比较常用
	Query []string `json:"query,omitempty"`
	// ImageMode 按照素材类型过滤，详见【附录-素材类型】
	ImageMode []enum.ImageMode `json:"image_mode,omitempty"`
}
