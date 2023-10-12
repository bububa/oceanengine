package eventmanager

import "github.com/bububa/oceanengine/marketing-api/util"

// TrackURLUpdateRequest 事件资产下更新监测链接组 API Request
type TrackURLUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetsID 资产ID
	AssetID uint64 `json:"asset_id,omitempty"`
	// DownloadURL 应用下载链接，应用下载链接，IOS和安卓应用资产：必填
	DownloadURL string `json:"download_url,omitempty"`
	// TrackURLGroup 监测链接组信息，IOS和安卓应用可绑定多组监测链接
	TrackURLGroup *TrackURLGroup `json:"track_url_group,omitempty"`
}

// Encode implement PostRequest interface
func (r TrackURLUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
