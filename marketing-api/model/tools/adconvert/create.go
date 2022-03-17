package adconvert

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建转化目标 API Request
type CreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 转化名称
	Name string `json:"name,omitempty"`
	// AppName 应用中文名
	AppName string `json:"app_name,omitempty"`
	// ConvertSourceType 转化来源
	ConvertSourceType enum.AdConvertSource `json:"convert_source_type,omitempty"`
	// ConvertType 转化类型
	ConvertType enum.AdConvertType `json:"convert_type,omitempty"`
	// DownloadURL 下载地址
	DownloadURL string `json:"download_url,omitempty"`
	// AppType 应用类型
	AppType string `json:"app_type,omitempty"`
	// ActionTrackURL 点击监测链接
	ActionTrackURL string `json:"action_track_url,omitempty"`
	// DisplayTrackURL 展示监测链接
	DisplayTrackURL string `json:"display_track_url,omitempty"`
	// VideoPlayEffectiveTrackURL 视频有效播放监测链接
	VideoPlayEffectiveTrackURL string `json:"video_play_effective_track_url,omitempty"`
	// VideoPlayDoneTrackURL 视频播放完毕监测链接
	VideoPlayDoneTrackURL string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayTrackURL 视频播放监测链接
	VideoPlayTrackURL string `json:"video_play_track_url,omitempty"`
	// PackageName 包名
	PackageName string `json:"package_name,omitempty"`
	// DeepExternalAction 深度转化目标，转化类型不同允许深度转化目标也不同，具体参考下方的【转化目标与深度转化目标关系】介绍
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
	// ConvertDataType 转化统计方式，针对每次付费广告，投放范围是站内和穿山甲，转化来源是应用下载SDK/API这两种方式，广告平台统计该转化目标是否发生的方式，默认“仅一次”（即，每个用户最多仅统计一次转化行为）。
	// 允许值：
	// ONLY_ONE（仅一次）：对于每位转化的用户，仅统计其首次“目标事件”的转化行为，即每位用户最多仅记录一次转化。
	// EVERY_ONE（每一次）：对于每位转化的用户，统计其每次“目标事件”的发生次数，即每位用户可记录多次发生的转化；该统计方式下，创建广告计划时deep_bid_type须为BID_PER_ACTION
	// 注意：如果广告主ID不在白名单里面，且统计方式选择EVERY_ONE，请求会失败，报错信息“convertDataType not in whiteList”。
	ConvertDataType string `json:"convert_data_type,omitempty"`
	// ConvertXPathURL 转化页面
	ConvertXPathURL string `json:"convert_xpath_url,omitempty"`
	// ConvertXPathValue 转化路径
	ConvertXPathValue string `json:"convert_xpath_value,omitempty"`
	// XPathIgnoreParams 匹配规则字段(xpath下可传)，允许值:
	// "UTM_ID"、"CID"、"ADID"
	XPathIgnoreParams []string `json:"xpath_ignore_params,omitempty"`
	// ExternalURL 落地页链接
	ExternalURL string `json:"external_url,omitempty"`
	// AppID APP ID
	AppID string `json:"app_id,omitempty"`
	// OpenURL 直达链接
	OpenURL string `json:"open_url,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建转化目标 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *Convert `json:"data,omitempty"`
}
