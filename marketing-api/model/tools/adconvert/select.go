package adconvert

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SelectRequest 转化目标列表 API Request
type SelectRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ConvertIDs 指定需要查询的转化目标ID，如不填写默认返回所有的转化目标ID
	ConvertIDs []uint64 `json:"convert_id,omitempty"`
	// OptStatus 转化工具操作状态
	OptStatus enum.AdConvertOptStatus `json:"opt_status,omitempty"`
	// Page 页数
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r SelectRequest) Encode() string {
	values := util.GetUrlValues()
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.ConvertIDs) > 0 {
		buf, _ := json.Marshal(r.ConvertIDs)
		values.Add("convert_ids", string(buf))
	}
	if r.OptStatus != "" {
		values.Add("opt_status", string(r.OptStatus))
	}
	if r.Page > 1 {
		values.Add("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Add("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// SelectResponse 转化目标列表 API Response
type SelectResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *SelectResponseData `json:"data,omitempty"`
}

// SelectResponseData json返回值
type SelectResponseData struct {
	// PageInfo 分页相关信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 转化的数据list
	List []Convert `json:"ad_convert_list,omitempty"`
}
