package tools

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeAuthListRequest 获取抖音授权关系
type AwemeAuthListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 筛选条件
	Filtering *AwemeAuthListFilter `json:"filtering,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

// AwemeAuthListFilter 筛选条件
type AwemeAuthListFilter struct {
	// AuthType 授权类型，可选值:
	// AWEME_ACCOUNT: 抖音号授权、VIDEO_ITEM: 单视频授权
	AuthType []enum.AwemeAuthType `json:"auth_type,omitempty"`
	// AuthStatus 授权状态， 可选值:
	// AUTHRIZED: 授权中、AUTHRIZING: 待授权确认、INVALID: 授权失效
	AuthStatus []enum.AwemeAuthStatus `json:"auth_status,omitempty"`
	// AwemeIDs 按抖音号id过滤结果，最长传入50个
	AwemeIDs []string `json:"aweme_ids,omitempty"`
	// ItemIDs 按抖音视频id过滤结果，一次最多允许查询50个
	ItemIDs []uint64 `json:"item_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeAuthListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		buf, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(buf))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AwemeAuthListResponse 获取抖音授权关系
type AwemeAuthListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AwemeAuthListData `json:"data,omitempty"`
}

// AwemeAuthListData .
type AwemeAuthListData struct {
	// List .
	List []AwemeAuthItem `json:"list,omitempty"`
	// PageInfo .
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type AwemeAuthItem struct {
	// AuthType 授权类型
	AuthType enum.AwemeAuthType `json:"auth_type,omitempty"`
	// AwemeID 抖音号
	AwemeID string `json:"aweme_id,omitempty"`
	// AwemeName 抖音账号名称
	AwemeName string `json:"aweme_name,omitempty"`
	// AuthStatus 授权状态
	AuthStatus enum.AwemeAuthStatus `json:"auth_status,omitempty"`
	// SubStatus 授权子状态，返回值
	// INVALID_CANCEL: 主动操作解除授权、INVALID_EXPIRED: 授权期限已到、INVALID_REJECT: C端拒绝授权、INVALID_TIME_OUT: 超时未确认、RENEWING: 续期待确认、RENEW_FAIL: 续期申请失效、RENEW_SUCCESS: 续期成功
	SubStatus enum.AwemeAuthSubStatus `json:"sub_status,omitempty"`
	// StartTime 授权开始时间，格式为yyyy-MM-dd HH:mm:ss
	StartTime string `json:"start_time,omitempty"`
	// EndTime 授权结束时间，格式为yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty"`
	// VideoInfo 授权视频信息，若为单视频授权会返回
	VideoInfo *AwemeAuthVideo `json:"video_info,omitempty"`
}

// AwemeAuthVideo 授权视频信息，若为单视频授权会返回
type AwemeAuthVideo struct {
	// ItemID 抖音视频ID
	ItemID uint64 `json:"item_id,omitempty"`
	// ImageMode 素材类型
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// Title 视频名称
	Title string `json:"title,omitempty"`
	// AwemePlayURL 视频播放链接
	AwemePlayURL string `json:"aweme_play_url,omitempty"`
	// Duration 视频时长，单位为秒
	Duration float64 `json:"duration,omitempty"`
	// VideoCoverID 视频封面ID
	VideoCoverID string `json:"video_cover_id,omitempty"`
	// VideoCoverURL 视频封面链接
	VideoCoverURL string `json:"video_cover_url,omitempty"`
}
