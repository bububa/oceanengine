package comment

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MetricsGetRequest 获取评论统计指标 API Request
type MetricsGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartTime 查询起始时间，格式：yyyy-MM-dd，时间跨度最大90天
	StartTime string `json:"start_time,omitempty"`
	// EndTime 查询截止时间，格式：yyyy-MM-dd，时间跨度最大90天
	EndTime string `json:"end_time,omitempty"`
	// Filtering 筛选条件
	Filtering *MetricsGetFilter `json:"filtering,omitempty"`
}

type MetricsGetFilter struct {
	// LevelType 评论层级，不传返回所有评论，可选值:
	// LEVEL_ALL 所有等级
	// LEVEL_ONE 一级评论
	// LEVEL_TWO 二级评论
	LevelType enum.CommentLevelType `json:"level_type,omitempty"`
	// AuthorIDs 抖音号，一次最多查询100个抖音号id
	AuthorIDs []uint64 `json:"author_ids,omitempty"`
	// ItemIDs 广告视频id列表，一次最多100个，可通过【获取抖音授权关系】获取item_id和视频相关信息
	ItemIDs []uint64 `json:"item_ids,omitempty"`
	// HideStatus 隐藏状态，不传返回全部，可选值:
	// ALL 全部
	// HIDE 已隐藏
	// NOT_HIDE 未隐藏
	HideStatus enum.CommentHideStatus `json:"hide_status,omitempty"`
}

func (r MetricsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MetricsGetResponse 获取评论统计指标 API Response
type MetricsGetResponse struct {
	model.BaseResponse
	Data *MetricsGetResult `json:"data,omitempty"`
}

type MetricsGetResult struct {
	// TotalCount 可见评论数
	TotalCount int64 `json:"total_count,omitempty"`
	// NegativeCount 可见负评数
	NegativeCount int64 `json:"negative_count,omitempty"`
	// NegativeRate 可见评论负评率
	NegativeRate float64 `json:"negative_rate,omitempty"`
}
