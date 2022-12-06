package adraise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VersionRequest 获取起量版本信息 API Request
type VersionRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：1-100，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r VersionRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
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

// VersionResponse 获取起量版本信息 API Response
type VersionResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *VersionResponseData `json:"data,omitempty"`
}

// VersionResponseData json 返回值
type VersionResponseData struct {
	// List 起量版本列表
	List []Version `json:"ad_raise_version_list,omitempty"`
	// PageInfo 页面信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// Version 起量版本
type Version struct {
	// Num 起量版本号
	Num uint64 `json:"version,omitempty"`
	// StartTime 当前版本起量开始时间，格式：2021-03-31 16:00:00
	StartTime string `json:"start_time,omitempty"`
	// EndTime 当前起量版本结束时间，格式：2021-03-31 17:00:00
	EndTime string `json:"end_time,omitempty"`
}
