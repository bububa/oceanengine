package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CarouselListRequest 获取图集素材 API Request
type CarouselListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 图集过滤信息
	Filtering *CarouselListFilter `json:"filtering,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值20
	PageSize int `json:"page_size,omitempty"`
}

// CarouselListFilter 图集过滤信息
type CarouselListFilter struct {
	// CarouselIDs 图集id
	CarouselIDs []uint64 `json:"carousel_ids,omitempty"`
	// FileName 文件名
	FileName string `json:"file_name,omitempty"`
	// CarouselType 图集素材类型
	CarouselType []enum.ImageMode `json:"carousel_type,omitempty"`
	// ImageIDs 图片id
	ImageIDs []string `json:"image_ids,omitempty"`
	// VideoID 音频id（旧版）（该字段将在12月19日下线，暂不对您的调用产生影响，请及时调整，使用下方audio_id）
	VideoID string `json:"video_id,omitempty"`
	// AudioID 音频id
	AudioID string `json:"audio_id,omitempty"`
	// StartTime 根据图集上传时间进行过滤的起始时间，与end_time搭配使用。
	// 格式：yyyy-mm-dd
	StartTime string `json:"start_time,omitempty"`
	// EndTime 根据图集上传时间进行过滤的结束时间，与start_time搭配使用。
	// 格式：yyyy-mm-dd
	EndTime string `json:"end_time,omitempty"`
	// Source 图文素材来源，可选值:
	// 本地上传：AD_SITE
	// 组织共享：BP
	// 账户推送：ACCOUNT_PUSH
	// 即创：AIC
	// MarketingAPI：OPEN_API
	// 抖音主页：CEWEBRITY_CAROUSEL
	Source string `json:"source,omitempty"`
}

// Encode implement GetRequest interface
func (r CarouselListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
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

// CarouselListRespones 获取图集素材 API Response
type CarouselListResponse struct {
	model.BaseResponse
	Data *CarouselListResult `json:"data,omitempty"`
}

type CarouselListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Carousels 图集信息
	Carousels []Carousel `json:"carousels,omitempty"`
}
