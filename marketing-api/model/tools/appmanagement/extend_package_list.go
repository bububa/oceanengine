package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ExtendPackageListRequest 查询应用分包列表 API Request
type ExtendPackageListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PackageID 应用包ID，获取方法见接口文档
	PackageID string `json:"package_id,omitempty"`
	// Status 状态
	// 允许值："ALL"：全部
	// "NOT_UPDATE"： 未更新
	// "CREATING"：创建中
	// "UPDATING"：更新中
	// "PUBLISHED"：已发布
	// "CREATION_FAILED"：创建失败
	// "UPDATE_FAILED"：更新失败
	// 默认值："ALL"
	Status ExtendPackageStatus `json:"status,omitempty"`
	// Page 页数 默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	// 默认值：10，范围：1～1000
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ExtendPackageListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("package_id", r.PackageID)
	if r.Status != "" {
		values.Set("status", string(r.Status))
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

// ExtendPackageListResponse 查询应用分包列表 API Request
type ExtendPackageListResponse struct {
	model.BaseResponse
	Data *ExtendPackageListData `json:"data,omitempty"`
}

type ExtendPackageListData struct {
	// List 应用分包列表
	List []ExtendPackage `json:"list,omitempty"`
	// PageInfo 页面信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// ExtendPackage 应用分包
type ExtendPackage struct {
	// Reason 分包失败原因
	Reason string `json:"reason,omitempty"`
	// Status 状态
	Status ExtendPackageStatus `json:"status,omitempty"`
	// UpdateTime 更新时间
	UpdateTime string `json:"update_time,omitempty"`
	// VersionName 版本号
	VersionName string `json:"version_name,omitempty"`
	// ChannelID 渠道号
	ChannelID string `json:"channel_id,omitempty"`
	// DownloadURL 下载链接
	DownloadURL string `json:"download_url,omitempty"`
	// PackageID 应用包ID
	PackageID string `json:"package_id,omitempty"`
	// Remark 备注
	Remark string `json:"remark,omitempty"`
}
