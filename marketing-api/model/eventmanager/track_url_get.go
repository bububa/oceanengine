package eventmanager

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TrackURLGetRequest 获取事件资产下的监测链接组 API Request
type TrackURLGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetID 资产ID
	AssetID uint64 `json:"asset_id,omitempty"`
	// DownloadURL 应用下载链接，应用下载链接，IOS和安卓应用资产：必填
	DownloadURL string `json:"download_url,omitempty"`
	// TrackURLGroupName 监测链接组名称
	TrackURLGroupName string `json:"track_url_group_name,omitempty"`
	// TrackURLGroupID 监测链接组ID
	TrackURLGroupID uint64 `json:"track_url_group_id,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 分页个数
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r TrackURLGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("asset_id", strconv.FormatUint(r.AssetID, 10))
	if r.DownloadURL != "" {
		values.Set("download_url", r.DownloadURL)
	}
	if r.TrackURLGroupName != "" {
		values.Set("track_url_group_name", r.TrackURLGroupName)
	}
	if r.TrackURLGroupID > 0 {
		values.Set("track_url_group_id", strconv.FormatUint(r.TrackURLGroupID, 10))
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

// TrackURLGetResponse 获取事件资产下的监测链接组 API Response
type TrackURLGetResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *TrackURLGetData `json:"data,omitempty"`
}

// TrackURLGetData 返回数据
type TrackURLGetData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// TrackURLGroups 监测链接组信息
	TrackURLGroups []TrackURLGroup `json:"track_url_groups,omitempty"`
}
