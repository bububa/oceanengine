package advertiser

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeAuthListGetRequest 获取千川账户下抖音号授权列表 API Request
type AwemeAuthListGetRequest struct {
	// Filtering 过滤条件
	Filtering *AwemeAuthListGetFilter `json:"filtering,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码，默认1
	Page int `json:"page,omitempty"`
	// PageSize 页码大小，可选值：10、20、50、100，默认10
	PageSize int `json:"page_size,omitempty"`
}

// AwemeAuthListGetFilter 过滤条件
type AwemeAuthListGetFilter struct {
	// AwemeShowID 抖音show_id
	// 注意：最多支持100个抖音号
	AwemeShowID []string `json:"aweme_show_id,omitempty"`
	// AwemeID 抖音id
	// 注意：最多支持100个抖音号
	AwemeID []uint64 `json:"aweme_id,omitempty"`
	// AuthStatus 授权状态筛选 可选值:
	// ALL 全部授权状态
	// EFFECTIVE 授权生效
	// EXPIRED 授权失效
	// WAIT_BIND 待达人确认授权
	AuthStatus string `json:"auth_status,omitempty"`
	// AuthType 授权类型筛选，仅支持传0，1，2，3 可选值:
	// ALL 全部授权类型
	// AWEME_COOPERATOR 合作达人
	// OFFICIAL 官方
	// SELF 自运营
	AuthType string `json:"auth_type,omitempty"`
	// AwemeName 抖音号名称
	// 注意：模糊查询
	AwemeName string `json:"aweme_name,omitempty"`
	// AuthRange 授权范围筛选 可选值:
	// ALL 全部授权
	// BY_AWEME 抖音号授权
	// BY_VIDEO 单视频授权
	AuthRange string `json:"auth_range,omitempty"`
	// IsCancellationProgress 待处理解除授权申请
	// 默认 false
	IsCancellationProgress bool `json:"is_cancellation_progress,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeAuthListGetRequest) Encode() string {
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

// AwemeAuthListGetResponse 获取千川账户下抖音号授权列表 API Response
type AwemeAuthListGetResponse struct {
	Data *AwemeAuthListGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type AwemeAuthListGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AuthorizationInfos 抖音号列表
	AuthorizationInfos []AwemeAuthInfo `json:"authorization_infos,omitempty"`
}

// AwemeAuthInfo 抖音号授权信息
type AwemeAuthInfo struct {
	// AuthVideoInfo 视频信息
	AuthVideoInfo *AwemeAuthVideoInfo `json:"auth_video_info,omitempty"`
	// AuthType 抖音号授权类型
	AuthType []string `json:"auth_type,omitempty"`
	// AwemeShowID 抖音号，即客户在手机端上看到的抖音号，若向客户披露抖音号请使用该字段
	AwemeShowID string `json:"aweme_show_id,omitempty"`
	// AwemeAvatar 抖音号头像
	AwemeAvatar string `json:"aweme_avatar,omitempty"`
	// AwemeName 抖音号名称
	AwemeName string `json:"aweme_name,omitempty"`
	// AuthSource 授权来源
	// STAR 星图
	// QIANCHUAN 千川PC
	// AWEME 随心推
	AuthSource string `json:"auth_source,omitempty"`
	// StartTime 授权开始时间
	StartTime string `json:"start_time,omitempty"`
	// EndTime 授权结束时间
	EndTime string `json:"end_time,omitempty"`
	// CreateTime 授权时间
	CreateTime string `json:"create_time,omitempty"`
	// AuthStatus 授权状态
	AuthStatus string `json:"auth_status,omitempty"`
	// AuthRange 授权范围
	AuthRange string `json:"auth_range,omitempty"`
	// AwemeID 抖音id
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// IsCancellationProgress 是否有待处理解除授权申请
	IsCancellationProgress bool `json:"is_cancellation_progress,omitempty"`
}

// AwemeAuthVideoInfo 视频信息
type AwemeAuthVideoInfo struct {
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
	// AwemeItemID 抖音短视频 ID
	AwemeItemID string `json:"aweme_item_id,omitempty"`
	// Title 抖音中的视频标题
	Title string `json:"title,omitempty"`
	// VideoCoverURL 视频封面
	VideoCoverURL string `json:"video_cover_url,omitempty"`
}
