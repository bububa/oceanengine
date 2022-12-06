package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TemplateGetRequest 获取DPA模板 API Request
type TemplateGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码， 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量， 默认值: 10，最大 100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r TemplateGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
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

// TemplateGetResponse 获取DPA模板 API Response
type TemplateGetResponse struct {
	model.BaseResponse
	// Data 返回值
	Data *TemplateGetResponseData `json:"data,omitempty"`
}

// TemplateGetResponseData 返回值
type TemplateGetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 模板列表
	List []Template `json:"list,omitempty"`
}
