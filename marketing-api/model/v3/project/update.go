package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 更新项目 API Request
type UpdateRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// DownloadMode 优先从系统应用商店下载（下载模式），枚举值：APP_STORE_DELIVERY 优先商店下载、 DEFAULT 默认下载
	DownloadMode enum.DownloadMode `json:"download_mode,omitempty"`
	// OpenURL Deeplink直达链接
	OpenURL string `json:"open_url,omitempty"`
	// UlinkURL ulink直达链接
	UlinkURL string `json:"urlink_url,omitempty"`
	// Audience 定向设置
	Audience *Audience `json:"audience,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// TrackURLSetting 监测链接设置
	TrackURLSetting *TrackURLSetting `json:"track_url_setting,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
