package report

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoFrameRequest 视频互动流失数据 API Request
type VideoFrameRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 起始日期。格式YYYY-MM-DD
	// -只支持查询2020-02-01及以后的日期
	// -与end_date的时间跨度最大支持31天，不支持获取当天数据
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate 结束日期。格式YYYY-MM-DD
	EndDate time.Time `json:"end_date,omitempty"`
	// Filtering 过滤字段。json格式
	Filtering VideoFrameFiltering `json:"filtering,omitempty"`
	// Metrics 数据指标。数组最大长度限制为5。 默认值: ['click_cnt', 'dy_comment', 'dy_follow', 'dy_like', 'user_lose_cnt']
	Metrics []string `json:"metrics,omitempty"`
}

// VideoFrameFiltering 过滤字段
type VideoFrameFiltering struct {
	// MaterislID 按素材id 过滤。素材id从上传视频接口的返回值中获取。仅支持单个素材id过滤。
	MaterialID uint64 `json:"material_id,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoFrameRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_date", r.StartDate.Format("2006-01-02"))
	values.Set("end_date", r.EndDate.Format("2006-01-02"))
	filtering, _ := json.Marshal(r.Filtering)
	values.Set("filtering", string(filtering))
	if r.Metrics != nil {
		metrics, _ := json.Marshal(r.Metrics)
		values.Set("metrics", string(metrics))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoFrameResponse 视频互动流失数据 API Response
type VideoFrameResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *VideoFrameResponseData `json:"data,omitempty"`
}

// VideoFrameResponseData json返回值
type VideoFrameResponseData struct {
	// List 应答指标列表
	List []VideoFrameResponseDataList `json:"list,omitempty"`
}

// VideoFrameResponseDataList 应答指标
type VideoFrameResponseDataList struct {
	// Metrics 指标数据
	Metrics *VideoFrameMetrics `json:"metrics,omitempty"`
	// Second 秒。指标数据对应的视频时间进度，例如：second=3表示视频第三秒时产生的指标数据
	Second int64 `json:"second,omitempty"`
}

// VideoFrameMetrics 指标
type VideoFrameMetrics struct {
	// ClickCnt 点击数
	ClickCnt int64 `json:"click_cnt,omitempty"`
	// DyComment 新增评论数
	DyComment int64 `json:"dy_comment,omitempty"`
	// DyFollow 新增关注数
	DyFollow int64 `json:"dy_follow,omitempty"`
	// DyLike 点赞数
	DyLike int64 `json:"dy_like,omitempty"`
	// UserLoseCnt 流失人数
	UserLoseCnt int64 `json:"user_lose_cnt,omitempty"`
}
