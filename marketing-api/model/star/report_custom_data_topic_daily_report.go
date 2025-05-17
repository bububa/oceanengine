package star

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReportCustomDataTopicDailyReportRequest 获取投后每日趋势数据（短视频） API Request
type ReportCustomDataTopicDailyReportRequest struct {
	// StartTime 起始时间: yyyy-MM-dd
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间: yyyy-MM-dd
	EndTime string `json:"end_time,omitempty"`
	// Topics 数据主题:
	// BASIC_DATA：基础信息
	// FLOW_DATA：流量表现
	// CONVERT_DATA：转化表现
	// SEARCH_DATA：搜索表现
	// RECOMMEND_DATA： 种草表现
	// DY_SHOP_DATA：抖音进店
	// USER_DISTRIBUTION_DATA：用户画像、
	// INDEX_SCORE_DATA： 指数得分
	// COMMENT_DATA：评论数据
	// 直播用户画像仅保留近90天且直播时长 >= 25 分钟直播数据
	Topics []enum.StarReportTopic `json:"topics,omitempty"`
	// StarID 发起请求的客户的starId
	StarID uint64 `json:"star_id,omitempty"`
	// WorkID 交付作品Id，如：视频Id，直播间room_id.
	WorkID uint64 `json:"work_id,omitempty"`
	// DemandID 任务Id
	DemandID uint64 `json:"demand_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ReportCustomDataTopicDailyReportRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("star_id", strconv.FormatUint(r.StarID, 10))
	values.Set("work_id", strconv.FormatUint(r.WorkID, 10))
	values.Set("demand_id", strconv.FormatUint(r.DemandID, 10))
	values.Set("topics", string(util.JSONMarshal(r.Topics)))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ReportCustomDataTopicDailyReportResponse 获取投后数据主题累计数据 API Response
type ReportCustomDataTopicDailyReportResponse struct {
	Data *ReportCustomDataTopicDailyReportResult `json:"data,omitempty"`
	model.BaseResponse
}

type ReportCustomDataTopicDailyReportResult struct {
	// Stats 统计数据列表
	Stats []ReportCustomDataTopicDailyReport `json:"stats,omitempty"`
}

type ReportCustomDataTopicDailyReport struct {
	// Date 该指标数据的yyyy-MM-dd日期时间, 按天聚合。BASIC_DATA作为通用指标在每日趋势数据中date为"0"。
	Date string `json:"date,omitempty"`
	// Data 对应请求的数据主题数组
	Data []DataTopicConfig `json:"data,omitempty"`
	// WorkID 交付作品Id：如视频Id
	ItemID uint64 `json:"item_id,omitempty"`
	// DemandID 请求的任务Id
	DemandID uint64 `json:"demand_id,omitempty"`
}
