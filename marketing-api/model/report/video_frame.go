package report

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type VideoFrameRequest struct {
	AdvertiserID uint64              `json:"advertiser_id,omitempty"`
	StartDate    time.Time           `json:"start_date,omitempty"`
	EndDate      time.Time           `json:"end_date,omitempty"`
	Filtering    VideoFrameFiltering `json:"filtering,omitempty"`
	Metrics      []string            `json:"metrics,omitempty"` // 数据指标。数组最大长度限制为5。 默认值: ['click_cnt', 'dy_comment', 'dy_follow', 'dy_like', 'user_lose_cnt']
}

type VideoFrameFiltering struct {
	MaterialID uint64 `json:"material_id,omitempty"` // 按 素材id 过滤。素材id从上传视频接口的返回值中获取。仅支持单个素材id过滤。
}

func (r VideoFrameRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_date", r.StartDate.Format("2006-01-02"))
	values.Set("end_date", r.EndDate.Format("2006-01-02"))
	filtering, _ := json.Marshal(r.Filtering)
	values.Set("filtering", string(filtering))
	if r.Metrics != nil {
		metrics, _ := json.Marshal(r.Metrics)
		values.Set("metrics", string(metrics))
	}
	return values.Encode()
}

type VideoFrameResponse struct {
	model.BaseResponse
	Data *VideoFrameResponseData `json:"data,omitempty"`
}

type VideoFrameResponseData struct {
	List []VideoFrameResponseDataList `json:"list,omitempty"`
}

type VideoFrameResponseDataList struct {
	Metrics *VideoFrameMetrics `json:"metrics,omitempty"` // 指标数据
	Second  int64              `json:"second,omitempty"`  // 秒。指标数据对应的视频时间进度，例如：second=3表示视频第三秒时产生的指标数据
}

type VideoFrameMetrics struct {
	ClickCnt    int64 `json:"click_cnt,omitempty"`     // 点击数
	DyComment   int64 `json:"dy_comment,omitempty"`    // 新增评论数
	DyFollow    int64 `json:"dy_follow,omitempty"`     // 新增关注数
	DyLike      int64 `json:"dy_like,omitempty"`       // 点赞数
	UserLoseCnt int64 `json:"user_lose_cnt,omitempty"` // 流失人数
}
