package eventmanager

import "github.com/bububa/oceanengine/marketing-api/util"

// TrackURLCreateRequest 事件资产下创建监测链接组 API Request
type TrackURLCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetsID 资产ID
	AssetsID uint64 `json:"assets_id,omitempty"`
	// DownloadURL 应用下载链接，应用下载链接，IOS和安卓应用资产：必填
	DownloadURL string `json:"download_url,omitempty"`
	// TrackURLGroups 监测链接组信息，IOS和安卓应用可绑定多组监测链接
	TrackURLGroups []TrackURLGroup `json:"track_url_groups,omitempty"`
}

// Encode implement PostRequest interface
func (r TrackURLCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
