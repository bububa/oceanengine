package adconvert

import "github.com/bububa/oceanengine/marketing-api/enum"

// Convert 转化目标详细信息
type Convert struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ID 转化id
	ID uint64 `json:"id,omitempty"`
	// AppType 应用类型
	AppType string `json:"app_type,omitempty"`
	// PackageName 包名
	PackageName string `json:"package_name,omitempty"`
	// DownloadURL 下载地址
	DownloadURL string `json:"download_url,omitempty"`
	// OptStatus 转化工具操作状态
	OptStatus enum.AdConvertOptStatus `json:"opt_status,omitempty"`
	// ConvertSourceType 转化来源
	ConvertSourceType enum.AdConvertSource `json:"convert_source_type,omitempty"`
	// Status 转化状态
	Status enum.AdConvertStatus `json:"status,omitempty"`
	// ConvertType 转化类型
	ConvertType enum.AdConvertType `json:"convert_type,omitempty"`
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
	// ConvertActivateCallbackURL 激活回传地址
	ConvertActivateCallbackURL string `json:"convert_activate_callback_url,omitempty"`
	// AppID APP ID
	AppID string `json:"app_id,omitempty"`
	// ExternalURL 落地页链接
	ExternalURL string `json:"external_url,omitempty"`
	// ConvertTrackParams 监测参数
	ConvertTrackParams string `json:"convert_tarck_params,omitempty"`
	// ConvertBaseCode 转化基础代码
	ConvertBaseCode string `json:"convert_base_code,omitempty"`
	// ConvertJSCode 转化代码（JS方式）
	ConvertJSCode string `json:"convert_js_code,omitempty"`
	// ConvertHTMLCode 转化代码（HTML方式）
	ConvertHTMLCode string `json:"convert_html_code,omitempty"`
	// ConvertXPathURL 转化页面
	ConvertXPathURL string `json:"convert_xpath_url,omitempty"`
	// ConvertXPathValue 转化路径
	ConvertXPathValue string `json:"convert_xpath_value,omitempty"`
	// OpenURL 直达链接
	OpenURL string `json:"open_url,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTime 更新时间
	ModifyTime string `json:"modify_time,omitempty"`
	// IgnoreParams 转化类型下匹配规则字段
	IgnoreParams []string `json:"ignore_params,omitempty"`
	// ConvertDataType 转化统计方式
	ConvertDataType enum.AdConvertDataType `json:"convert_data_type,omitempty"`
}
