package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoGetRequest 获取素材库视频 API Request
type VideoGetRequest struct {
	// LocalAccountID 本地推账户id
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// Filtering 过滤器
	Filtering *VideoGetFilter `json:"filtering,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页大小,默认值：20，最大值：100
	PageSize int `json:"page_size,omitempty"`
}

type VideoGetFilter struct {
	// SearchKeyWord 根据视频名称或ID筛选
	SearchKeyWord string `json:"search_key_word,omitempty"`
	// ImageMode 素材类型 允许值：
	// IMAGE_MODE_VIDEO 横版视频
	// IMAGE_MODE_VIDEO_VERTICAL 竖版视频
	ImageMode local.ImageMode `json:"image_mode,omitempty"`
	// MaterialSource 素材来源 允许值：
	// BP_PLATFORM 巨量引擎工作平台共享视频
	// CREATIVE_AIGC 即创
	// LOCAL_ADS_UPLOAD 本地上传
	// STAR 星图平台
	// MAPI MAPI上传
	MaterialSource local.MaterialSource `json:"material_source,omitempty"`
	// AnalysisType 评估类型 允许值：
	// FIRST_PUBLISH 首发
	// FIRST_PUBLISH_AND_HIGH_QUALITY 首发&优质
	// HIGH_QUALITY 优质
	AnalysisType local.MaterialAnalysisType `json:"analysis_type,omitempty"`
	// StartTime 根据视频上传时间进行过滤的起始时间，与end_time搭配使用，格式：yyyy-MM-dd HH:mm:ss
	StartTime string `json:"start_time,omitempty"`
	// EndTime 根据视频上传时间进行过滤的截止时间，与start_time搭配使用，格式：yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty"`
	// IsFilterUnqualified 是否过滤低质素材，允许值：
	// false 不过滤
	// true 过滤（默认值）
	IsFilterUnqualified *bool `json:"is_filter_unqualified,omitempty"`
	// OrderField 排序字段，允许值：
	// CONVERSION_COST 转化成本
	// CONVERSION_RATE 转化率
	// CREATE_TIME 创建时间（默认值）
	// CTR 点击率
	// DURATION 视频时长
	// STAT_COST 消耗
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序顺序，允许值：
	// ASC 升序
	// DESC 降序（默认值）
	OrderType enum.OrderType `json:"order_type,omitempty"`
}

// Encode implements GetRequest interface
func (r VideoGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// VideoGetResponse 获取素材库视频 API Response
type VideoGetResponse struct {
	model.BaseResponse
	Data *VideoGetResult `json:"data,omitempty"`
}

type VideoGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// VideoList 素材库视频列表
	VideoList []Video `json:"video_list,omitempty"`
}
