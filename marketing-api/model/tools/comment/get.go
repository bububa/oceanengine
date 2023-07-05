package comment

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取评论列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformVersion 查询的平台版本，查询体验版需传入V2，允许值：
	// V1: 1.0平台（默认值）
	// V2: 2.0平台
	PlatformVersion enum.PlatformVersion `json:"platform_version,omitempty"`
	// StartTime 查询起始时间，格式：yyyy-MM-dd，若不填，默认6天前（即获取最近七天的内容）
	StartTime string `json:"start_time,omitempty"`
	// EndTime 查询截止时间，格式：yyyy-MM-dd，若不填，默认当天
	EndTime string `json:"end_time,omitempty"`
	// IncludeMetrics 是否包含计算指标，传入True则返回负评数和负评率，默认为True
	IncludeMetrics bool `json:"include_metrics,omitempty"`
	// OrderField 排序字段，允许值：
	// REPLY_COUNT按评论回复数量排序
	// LIKE_COUNT 按点赞数量排序
	// CREATE_TIME 按评论时间排序（默认值）
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序类型，允许值：
	// ASC: 升序
	// DESC: 降序（默认值）
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页数 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值: 10
	PageSize int `json:"page_size,omitempty"`
	// Filtering 过滤条件，若此字段不传，或传空则视为无限制条件
	Filtering *GetFilter `json:"filtering,omitempty"`
}

type GetFilter struct {
	// AdIDs 计划id列表，一次最多10个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// CreativeIDs 广告创意id列表，一次最多100个，仅当platform_version=V1时传入生效
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// ItemIDs 广告视频id列表，一次最多100个，仅当platform_version=V2时传入生效，可通过【获取抖音授权关系】获取item_id和视频相关信息
	ItemIDs []uint64 `json:"item_ids,omitempty"`
	// BindRelation 抖音号分类，可选值:
	// ALL: 全部抖音号（默认值）
	// BIND_AWEME: 绑定的抖音号
	// VIRTUAL_AWEME: 虚拟抖音号
	BindRelation string `json:"bind_relation,omitempty"`
	// LevelType 查询的评论等级，可选值:
	// LEVEL_ALL: 所有评论（默认值）
	// LEVEL_ONE: 一级评论
	// LEVEL_TWO: 二级评论
	LevelType enum.CommentLevelType `json:"level_type,omitempty"`
	// IsReplied 是否已回复
	IsReplied bool `json:"is_replied,omitempty"`
	// HideStatus 隐藏状态， 可选值:
	// ALL: 全部(默认值）
	// HIDE: 已隐藏
	// NOT_HIDE: 未隐藏
	HideStatus enum.CommentHideStatus `json:"hide_status,omitempty"`
	// EmotionType 评论情感， 可选值:
	// NEGATIVE: 负向评论
	// NEUTRAL: 中性评论
	// POSITIVE: 正向评论
	EmotionType enum.CommentEmotionType `json:"emotion_type,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("platform_version", r.PlatformVersion)
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	if !r.IncludeMetrics {
		values.Set("include_metrics", "false")
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
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

// GetResponse 获取评论列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json 返回值
type GetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// NegativeVolumn 负评数
	NegativeVolumn int `json:"negative_volumn,omitempty"`
	// NegativeVolumnRate 负评率
	NegativeVolumnRate float64 `json:"negative_volumn_rate,omitempty"`
	// CommentsList 评论列表
	CommentsList []Comment `json:"comment_list,omitempty"`
}
