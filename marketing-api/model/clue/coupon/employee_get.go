package coupon

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// EmployeeGetRequest 查询核销员 API Request
type EmployeeGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值10，不超过50
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r EmployeeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
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

// EmployeeGetResponse 查询核销员 API Response
type EmployeeGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *EmployeeGetResponseData `json:"data,omitempty"`
}

// EmployeeGetResponseData json返回值
type EmployeeGetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 核销员列表
	List []Employee `json:"list,omitempty"`
}
