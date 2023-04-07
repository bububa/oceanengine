package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VersionGetRequest 获取起量版本信息 API Request
type VersionGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：1-100，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r VersionGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("promotion_id", strconv.FormatUint(r.PromotionID, 10))
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

// VersionGetResponse 获取起量版本信息 API Response
type VersionGetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *VersionGetResponseData `json:"data,omitempty"`
}

// VersionGetResponseData json 返回值
type VersionGetResponseData struct {
	// List 起量版本列表
	List []Version `json:"ad_raise_version_list,omitempty"`
	// PageInfo 页面信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// Version 起量版本
type Version struct {
	// RaiseVersion 起量版本号
	RaiseVersion uint64 `json:"raise_version,omitempty"`
	// StartTime 当前版本起量开始时间，格式：2021-03-31 16:00:00
	StartTime string `json:"start_time,omitempty"`
	// EndTime 当前起量版本结束时间，格式：2021-03-31 17:00:00
	EndTime string `json:"end_time,omitempty"`
}
