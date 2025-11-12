package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AppListRequest 查询应用信息 API Request
type AppListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SearchKey 搜索关键字
	// appid或者应用名，可以为空，可以传中文
	// 长度不超过50
	SearchKey string `json:"search_key,omitempty"`
	// SearchType 搜索类型:
	// ALL:查询全部，包括创建和被共享的应用（默认）
	// CREATE_ONLY:只查询广告主创建的应用
	// SHARED_ONLY:只查询被共享的应用
	SearchType AppSearchType `json:"search_type,omitempty"`
	// Status 应用状态:
	// ALL:所有状态
	// AUDIT_DOING:审核中
	// AUDIT_REJECTED:审核失败
	// BOOKING:预约中
	// ENABLE:已发布（默认）
	// PAST_DUE:已逾期
	Status AppStatus `json:"status,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，最大不超过200
	PageSize int `json:"page_size,omitempty"`
	// CreateTime 按创建时间查询的时间范围
	CreateTime *TimeRange `json:"create_time,omitempty"`
	// PublishTime 按发布时间查询的时间范围
	PublishTime *TimeRange `json:"publish_time,omitempty"`
	// ScheduledPublishTime 按预约发布时间查询的时间范围
	ScheduledPublishTime *TimeRange `json:"scheduled_publish_time,omitempty"`
}

// Encode implement GetRequest interface
func (r AppListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.SearchKey != "" {
		values.Set("search_key", r.SearchKey)
	}
	if r.SearchType != "" {
		values.Set("search_type", string(r.SearchType))
	}
	if r.Status != "" {
		values.Set("status", string(r.Status))
	}
	if r.CreateTime != nil {
		values.Set("create_time", string(util.JSONMarshal(r.CreateTime)))
	}
	if r.PublishTime != nil {
		values.Set("publish_time", string(util.JSONMarshal(r.PublishTime)))
	}
	if r.ScheduledPublishTime != nil {
		values.Set("scheduled_publish_time", string(util.JSONMarshal(r.ScheduledPublishTime)))
	}
	if r.Page != 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AppListResponse 查询APP信息 API Response
type AppListResponse struct {
	model.BaseResponse
	Data *AppListResponseData `json:"data,omitempty"`
}

type AppListResponseData struct {
	// List 应用列表
	List []App `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// App 应用信息
type App struct {
	// PackageID 应用包ID
	PackageID string `json:"package_id,omitempty"`
	// PackageName 包名
	PackageName string `json:"package_name,omitempty"`
	// AppCloudID app id
	AppCloudID uint64 `json:"app_cloud_id,omitempty"`
	// AppID 鸿蒙应用ID
	AppID string `json:"app_id,omitempty"`
	// AppName 应用名
	AppName string `json:"app_name,omitempty"`
	// VersionName 版本号
	VersionName string `json:"version_name,omitempty"`
	// Version 版本号
	Version string `json:"version,omitempty"`
	// DownloadURL 下载地址
	DownloadURL string `json:"download_url,omitempty"`
	// IconURL icon地址
	IconURL string `json:"icon_url,omitempty"`
	// Status 应用当前状态
	Status AppStatus `json:"status,omitempty"`
	// AssetID 鸿蒙资产ID
	AssetID string `json:"asset_id,omitempty"`
	// VersionID 版本审核ID，每次送审产生新的ID
	VersionID uint64 `json:"version_id,omitempty"`
	// PublishTime 发布时间，格式：%Y-%m-%d %H:%M:%S
	PublishTime string `json:"publish_time,omitempty"`
	// ScheduledPublishTime 预约发布时间，格式：%Y-%m-%d %H:%M:%S
	ScheduledPublishTime string `json:"scheduled_publish_time,omitempty"`
	// UpdateTime 更新时间，格式：%Y-%m-%d %H:%M:%S
	UpdateTime string `json:"update_time,omitempty"`
	// CreateTime 创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
}
