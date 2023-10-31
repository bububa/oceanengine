package enterprise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/enterprise"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReportRequest 获取数据 API Request
type ReportRequest struct {
	// AdvertiserID 广告主ID (必填)
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartTime 报表开始时间，格式为yyyy-mm-dd，仅支持查询2020年11月1日之后的数据 (必填)
	StartTime string `json:"start_time,omitempty"`
	// EndTime 报表结束时间，格式为yyyy-mm-dd，最大时间范围365天 (必填)
	EndTime string `json:"end_time,omitempty"`
	// LastStartTime 环比周期开始时间，格式为yyyy-mm-dd
	LastStartTime string `json:"last_start_time,omitempty"`
	// LastEndTime 环比周期结束时间，格式为yyyy-mm-dd
	LastEndTime string `json:"last_end_time,omitempty"`
	// Fields 指标字段，默认传入指标：business_page_home_visited、business_page_new_fans_num、business_page_like、business_page_total_play，若需自定义，请点击查看各维度下支持查询的指标，否则会报错
	Fields []string `json:"fields,omitempty"`
	// RatioFields 需要环比指标字段
	RatioFields []string `json:"ratio_fields,omitempty"`
	// TimeDimension 是否需要分天数据，默认值SUM允许值：SUM 合计数据、DAILY 分天数据
	TimeDimension enterprise.TimeDimension `json:"time_dimension,omitempty"`
	// OrderFIeld 排序指标字段，不传默认为空
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序类型；默认值: DESC允许值: ASC、 DESC，未传入order_field时不生效
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Filter 筛选字段
	Filter *ReportFilter `json:"filter,omitempty"`
	// Offset 偏移，类似于SQL中offset（起始为0，翻页时new_offset=old_offset+limit），默认值：0，取值范围:≥ 0
	Offset int `json:"offset,omitempty"`
	// Limit 返回数据量，默认值：100，取值范围：1-100
	Limit int `json:"limit,omitempty"`
}

type ReportFilter struct {
	// OpenID 抖音号open_id，同时支持查询企业号和普通抖音号数据
	OpenID string `json:"open_id,omitempty"`
}

func (r ReportRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	if r.LastStartTime != "" {
		values.Set("last_start_time", r.LastStartTime)
	}
	if r.LastEndTime != "" {
		values.Set("last_end_time", r.LastEndTime)
	}
	if r.TimeDimension != "" {
		values.Set("time_dimension", string(r.TimeDimension))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.Offset > 0 {
		values.Set("offset", strconv.Itoa(r.Offset))
	}
	if r.Limit > 0 {
		values.Set("limit", strconv.Itoa(r.Limit))
	}
	if r.Filter != nil {
		values.Set("filter", string(util.JSONMarshal(r.Filter)))
	}
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if len(r.RatioFields) > 0 {
		values.Set("ratio_fields", string(util.JSONMarshal(r.RatioFields)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ReportResponse 获取数据 API Response
type ReportResponse struct {
	model.BaseResponse
	Data *ReportResult `json:"data,omitempty"`
}

type ReportResult struct {
	// List 指标列表
	List         []Stat   `json:"list,omitempty"`
	TotalMetrics *Metrics `json:"total_metrics,omitempty"`
	TotalRatio   *Metrics `json:"total_ratio,omitempty"`
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// Offset 偏移
	Offset int `json:"offset,omitempty"`
}

// Stat
type Stat struct {
	Dimension
	Metrics
}

type Dimension struct {
	// StartTimeDay 数据统计时间
	StartTimeDay string `json:"start_time_day,omitempty"`
	// FirstFlowCategory 流量来源名称
	FirstFlowCategory string `json:"first_flow_category,omitempty"`
	// VideoID 视频itemId
	VideoID uint64 `json:"video_id,omitempty"`
}

// Metrics 指标
type Metrics struct {
	// BusinessPageHomeVisited 主页访问量
	BusinessPageHomeVisited model.Int64 `json:"business_page_home_visited,omitempty"`
	// BusinessPageHomeVisitedUCount 主页访问人数
	BusinessPageHomeVisitedUCount model.Int64 `json:"business_page_home_visited_ucount,omitempty"`
	// BusinessPageFansNumAll 总粉丝数
	BusinessPageFansNumAll model.Int64 `json:"business_page_fans_num_all,omitempty"`
	// BusinessPageNewFansNum 新增粉丝人数
	BusinessPageNewFansNum model.Int64 `json:"business_page_new_fans_num,omitempty"`
	// BusinessPageDisfollowUCount  流失粉丝人数
	BusinessPageDisfollowUCount model.Int64 `json:"business_page_disfollow_ucount,omitempty"`
	// BusinessPageComment 评论提交数
	BusinessPageComment model.Int64 `json:"business_page_comment,omitempty"`
	// BusinessPageLike 点赞数
	BusinessPageLike model.Int64 `json:"business_page_like,omitempty"`
	// BusinessPageShare 分享数
	BusinessPageShare model.Int64 `json:"business_page_share,omitempty"`
	// BusinessPageDpDownClickCnt 保存数
	BusinessPageDpDownClickCnt model.Int64 `json:"business_page_dp_down_click_cnt,omitempty"`
	// BusinessPageFavouriteVideoCnt 收藏数
	BusinessPageFavouriteVideoCnt model.Int64 `json:"business_page_favourite_video_cnt,omitempty"`
	// BusinessPageFwVideoVV 转发数
	BusinessPageFwVideoVV model.Int64 `json:"business_page_fw_video_vv"`
	// BusinessPagePlayOver 播放完成数
	BusinessPagePlayOver model.Int64 `json:"business_page_play_over"`
	// BusinessPageTotalPlay 播放数
	BusinessPageTotalPlay model.Int64 `json:"business_page_total_play"`
	// BusinessPageDislikeCnt 不感兴趣数
	BusinessPageDislikeCnt model.Int64 `json:"business_page_dislike_cnt"`
	// BusinessPageChatByShareVideoCnt 视频私信分享数
	BusinessPageChatByShareVideoCnt model.Int64 `json:"business_page_chat_by_share_video_cnt"`
	// BusinessPageClickCartCount 查看购物车数
	BusinessPageClickCartCount model.Int64 `json:"business_page_click_cart_count"`
	// BusinessPageClickProduct 商品点击数
	BusinessPageClickProduct model.Int64 `json:"business_page_click_product"`
	// BusinessPageClickProductGoBuy 去购买点击数
	BusinessPageClickProductGoBuy model.Int64 `json:"business_page_click_product_go_buy"`
	// BusinessPageOrderCount 商品下单数
	BusinessPageOrderCount model.Int64 `json:"business_page_order_count"`
	// BusinessPagePayOrderCount 商品订单数
	BusinessPagePayOrderCount model.Int64 `json:"business_page_pay_order_count"`
	// BusinessPagePayOrderGMV 商品订单金额
	BusinessPagePayOrderGMV model.Float64 `json:"business_page_pay_order_gmv"`
}
