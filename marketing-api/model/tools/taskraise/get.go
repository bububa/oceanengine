package taskraise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 查询优选起量任务 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformVersion 创建优选起量任务的平台版本，不传值默认为巨量引擎平台原版，允许值：V2 巨量引擎体验版。可用于创建巨量引擎体验版上的账户优选起量任务
	PlatformVersion enum.PlatformVersion `json:"platform_version,omitempty"`
	// Page 页数默认值: 1，page范围为[1,99999]
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值:10，page_size范围为[1,100]
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.PlatformVersion != "" {
		values.Set("platform_version", string(r.PlatformVersion))
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

// GetResponse 查询优选起量任务 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData .
type GetResponseData struct {
	// Reports 任务详情
	Reports []Report `json:"reports,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
