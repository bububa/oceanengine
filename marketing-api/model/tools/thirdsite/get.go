package thirdsite

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取第三方落地页站点列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页数默认值: 1，page必须大于0
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值:10，page_size范围为1-1000
	PageSize int `json:"page_size,omitempty"`
	// Filtering 过滤字段
	Filtering *GetFiltering `json:"filtering,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
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

// GetFiltering 过滤字段
type GetFiltering struct {
	// SiteID 搜索字段，按照建站id进行模糊匹配
	SiteID string `json:"site_id,omitempty"`
	// SiteName 搜索字段，按照建站name进行模糊匹配
	SiteName string `json:"site_name,omitempty"`
	// StartTime 时间过滤条件：时间范围内创建的落地页，开始时间,默认七天前, 时间格式 %Y-%m-%d
	StartTime string `json:"start_time,omitempty"`
	// EndTime 时间过滤条件：时间范围内创建的落地页，结束时间，默认到当天,时间格式 %Y-%m-%d
	EndTime string `json:"end_time,omitempty"`
}

// GetResponse 获取第三方落地页站点列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// List 站点列表
	List []Site `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
