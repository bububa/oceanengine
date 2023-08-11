package enterprise

import (
	"encoding/json"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
	"strconv"
)

type GetDataRequest struct {
	AdvertiserID  uint64        `json:"advertiser_id"`   // 广告主ID (必填)
	StartTime     string        `json:"start_time"`      // 报表开始时间，格式为yyyy-mm-dd，仅支持查询2020年11月1日之后的数据 (必填)
	EndTime       string        `json:"end_time"`        // 报表结束时间，格式为yyyy-mm-dd，最大时间范围365天 (必填)
	LastStartTime string        `json:"last_start_time"` // 环比周期开始时间，格式为yyyy-mm-dd
	LastEndTime   string        `json:"last_end_time"`   // 环比周期结束时间，格式为yyyy-mm-dd
	Fields        []string      `json:"fields"`          // 指标字段，默认传入指标：business_page_home_visited、business_page_new_fans_num、business_page_like、business_page_total_play，若需自定义，请点击查看各维度下支持查询的指标，否则会报错
	RatioFields   []string      `json:"ratio_fields"`    // 需要环比指标字段
	TimeDimension string        `json:"time_dimension"`  // 是否需要分天数据，默认值SUM允许值：SUM 合计数据、DAILY 分天数据
	OrderField    string        `json:"order_field"`     // 排序指标字段，不传默认为空
	OrderType     string        `json:"order_type"`      // 排序类型；默认值: DESC允许值: ASC、 DESC，未传入order_field时不生效
	Filter        *OpenIDFilter `json:"filter"`          // 筛选字段
	Offset        int           `json:"offset"`          // 偏移，类似于SQL中offset（起始为0，翻页时new_offset=old_offset+limit），默认值：0，取值范围:≥ 0
	Limit         int           `json:"limit"`           // 返回数据量，默认值：100，取值范围：1-100
}

func (r GetDataRequest) Encode() string {
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
	values.Set("time_dimension", r.TimeDimension)
	values.Set("order_field", r.OrderField)
	if r.OrderType != "" {
		values.Set("order_type", r.OrderType)
	}
	values.Set("offset", strconv.Itoa(r.Offset))
	values.Set("limit", strconv.Itoa(r.Limit))
	if r.Filter != nil {
		filtering, _ := json.Marshal(r.Filter)
		values.Set("filter", string(filtering))
	}
	if len(r.Fields) > 0 {
		fields, _ := json.Marshal(r.Fields)
		values.Set("fields", string(fields))
	}
	if len(r.RatioFields) > 0 {
		ratioFields, _ := json.Marshal(r.RatioFields)
		values.Set("ratio_fields", string(ratioFields))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

type OpenIDFilter struct {
	OpenID string `json:"open_id"` // 抖音号open_id，同时支持查询企业号和普通抖音号数据
}

type BasicData struct {
	BusinessPageHomeVisited         string `json:"business_page_home_visited"`            // 主页访问量
	BusinessPageHomeVisitedUCount   string `json:"business_page_home_visited_ucount"`     // 主页访问人数
	BusinessPageFansNumAll          string `json:"business_page_fans_num_all"`            // 总粉丝数
	BusinessPageNewFansNum          string `json:"business_page_new_fans_num"`            // 新增粉丝人数
	BusinessPageDisfollowUCount     string `json:"business_page_disfollow_ucount"`        // 流失粉丝人数
	BusinessPageComment             string `json:"business_page_comment"`                 // 评论提交数
	BusinessPageLike                string `json:"business_page_like"`                    // 点赞数
	BusinessPageShare               string `json:"business_page_share"`                   // 分享数
	BusinessPageDpDownClickCnt      string `json:"business_page_dp_down_click_cnt"`       // 保存数
	BusinessPageFavouriteVideoCnt   string `json:"business_page_favourite_video_cnt"`     // 收藏数
	BusinessPageFwVideoVV           string `json:"business_page_fw_video_vv"`             // 转发数
	BusinessPagePlayOver            string `json:"business_page_play_over"`               // 播放完成数
	BusinessPageTotalPlay           string `json:"business_page_total_play"`              // 播放数
	BusinessPageDislikeCnt          string `json:"business_page_dislike_cnt"`             // 不感兴趣数
	BusinessPageChatByShareVideoCnt string `json:"business_page_chat_by_share_video_cnt"` // 视频私信分享数
	BusinessPageClickCartCount      string `json:"business_page_click_cart_count"`        // 查看购物车数
	BusinessPageClickProduct        string `json:"business_page_click_product"`           // 商品点击数
	BusinessPageClickProductGoBuy   string `json:"business_page_click_product_go_buy"`    // 去购买点击数
	BusinessPageOrderCount          string `json:"business_page_order_count"`             // 商品下单数
	BusinessPagePayOrderCount       string `json:"business_page_pay_order_count"`         // 商品订单数
	BusinessPagePayOrderGMV         string `json:"business_page_pay_order_gmv"`           // 商品订单金额
}

type OverviewData struct {
	BasicData
	StartTimeDay string `json:"start_time_day"` // 数据统计时间
}

type GetOverviewResponse struct {
	model.BaseResponse
	Data struct {
		List         []*OverviewData `json:"list"`
		TotalMetrics BasicData       `json:"total_metrics"`
		TotalRatio   BasicData       `json:"total_ratio"`
	} `json:"data"`
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// TotalPage 总页数
	TotalPage int `json:"total_page,omitempty"`
}

type FlowCategoryData struct {
	StartTimeDay      string `json:"start_time_day"`      // 数据统计时间
	FirstFlowCategory string `json:"first_flow_category"` // 流量来源名称
	BasicData
}

type GetFlowCategoryResponse struct {
	model.BaseResponse
	Data struct {
		List         []*FlowCategoryData `json:"list"`
		TotalMetrics BasicData           `json:"total_metrics"`
		TotalRatio   BasicData           `json:"total_ratio"`
	} `json:"data"`
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// TotalPage 总页数
	TotalPage int `json:"total_page,omitempty"`
}

type VideoInfoData struct {
	StartTimeDay string `json:"start_time_day"` // 数据统计时间
	VideoId      string `json:"video_id"`       // 视频itemId
	BasicData
}

type GetVideoInfoDataResponse struct {
	model.BaseResponse
	Data struct {
		List         []*VideoInfoData `json:"list"`
		TotalMetrics BasicData        `json:"total_metrics"`
		TotalRatio   BasicData        `json:"total_ratio"`
	} `json:"data"`
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// TotalPage 总页数
	TotalPage int `json:"total_page,omitempty"`
}
