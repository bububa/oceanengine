package duoplus

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrderReportRequest 获取订单数据报表 API Request
type OrderReportRequest struct {
	// AwemeSecUid 抖音号ID，通过已授权账户接口可以获取
	AwemeSecUid string `json:"aweme_sec_uid,omitempty"`
	// StatTime 时间范围
	StatTime *TimeRange `json:"stat_time,omitempty"`
	// GroupBy 分组方式？
	// GROUP_BY_AD_ID 按订单分组（默认值，若为按订单分组，则返回内容item_id,room_id不生效）
	// GROUP_BY_CREATIVE_ID 按素材（视频/直播间）分组
	// 目前支持传入方式包括：
	// [GROUP_BY_AD_ID]按订单拆分数据
	// [GROUP_BY_AD_ID,GROUP_BY_CREATIVE_ID] 按素材拆分数据 ，当传入该组合时，暂时将不支持分页，同时filter中的order_ids必传
	GroupBy []string `json:"group_by,omitempty"`
	// Filter 数据过滤
	Filter *OrderReportFilter `json:"filter,omitempty"`
	// Page 页码，默认值：0
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：1- 100，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// TimeRange 时间范围
type TimeRange struct {
	// BeginTime 开始时间，格式 2021-04-05
	BeginTime string `json:"begin_time,omitempty"`
	// EndTime 结束时间，格式 2021-04-05，时间跨度不能超过180天
	EndTime string `json:"end_time,omitempty"`
}

// OrderReportFilter 数据过滤
type OrderReportFilter struct {
	// ItemIDs 视频id
	ItemIDs []uint64 `json:"item_ids,omitempty"`
	// OrderIDs 订单id
	OrderIDs []uint64 `json:"order_ids,omitempty"`
	// RoomIDs 直播间id
	RoomIDs []uint64 `json:"room_ids,omitempty"`
	// AuthorUniqueIDs 达人id
	AuthorUniqueIDs []string `json:"author_unique_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r OrderReportRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("aweme_sec_uid", r.AwemeSecUid)
	if r.StatTime != nil {
		values.Set("stat_time", string(util.JSONMarshal(r.StatTime)))
	}
	if len(r.GroupBy) > 0 {
		values.Set("group_by", string(util.JSONMarshal(r.GroupBy)))
	}
	if r.Filter != nil {
		values.Set("filter", string(util.JSONMarshal(r.Filter)))
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

// OrderReportResponse 获取订单数据报表 API Response
type OrderReportResponse struct {
	model.BaseResponse
	Data *OrderReportResult `json:"data,omitempty"`
}

type OrderReportResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// OrderMetrics 订单数据
	OrderMetrics []OrderStat `json:"order_metrics,omitempty"`
}

// OrderStat 订单数据
type OrderStat struct {
	// DimensionData 数据维度
	DimensionData *OrderDimension `json:"dimension_data,omitempty"`
	// MetricsData 指标值
	MetricsData *OrderMetrics `json:"metrics_data,omitempty"`
}

// OrderDimension 数据维度
type OrderDimension struct {
	// OrderID 订单id
	OrderID uint64 `json:"order_id,omitempty"`
	// AdvertiserID 账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// IsFans 是否粉丝必见计划
	IsFans bool `json:"is_fans,omitempty"`
	// ItemID 视频id
	ItemID uint64 `json:"item_id,omitempty"`
	// RoomID 直播间id
	RoomID uint64 `json:"room_id,omitempty"`
}

// OrderMetrics 指标值
type OrderMetrics struct {
	// StatCost 消耗
	StatCost float64 `json:"stat_cost,omitempty"`
	// TotalPlay 播放次数
	TotalPlay int64 `json:"total_play,omitempty"`
	// DyFollow 新增粉丝数
	DyFollow int64 `json:"dy_follow,omitempty"`
	// DyShare 分享次数
	DyShare int64 `json:"dy_share,omitempty"`
	// DyComment 评论次数
	DyComment int64 `json:"dy_comment,omitempty"`
	// CustomLike 点赞次数
	CustomLike int64 `json:"custom_like,omitempty"`
	// DyHomeVisited 主页访问次数
	DyHomeVisited int64 `json:"dy_home_visited,omitempty"`
	// PlayDuration5sRank 5s完播率，短视频内容分析指标
	PlayDuration5sRank float64 `json:"play_duration_5s_rank,omitempty"`
	// DpTargetConvertCnt 转化数
	DpTargetConvertCnt int64 `json:"dp_target_convert_cnt,omitempty"`
	// CustomConvertCost 转化成本
	CustomConvertCost float64 `json:"custom_convert_cost,omitempty"`
	// ShowCnt 直播间展示数
	ShowCnt int64 `json:"show_cnt,omitempty"`
	// LiveClickSourceCnt 直播间新增观众数
	LiveClickSourceCnt int64 `json:"live_click_source_cnt,omitempty"`
	// LiveGiftUV 直播间打赏观众人数
	LiveGiftUV int64 `json:"live_gift_uv,omitempty"`
	// LiveGiftAmount 直播间音浪收入
	LiveGiftAmount float64 `json:"live_gift_amount,omitempty"`
	// LiveCommentCnt 直播间评论次数
	LiveCommentCnt int64 `json:"live_comment_cnt,omitempty"`
	// DuoplusLiveFollowCnt 直播间新增粉丝量
	DuoplusLiveFollowCnt int64 `json:"duoplus_live_follow_cnt,omitempty"`
	// LiveGiftCnt 直播间打赏次数
	LiveGiftCnt int64 `json:"live_gift_cnt,omitempty"`
}
