package report

import (
	"github.com/bububa/oceanengine/marketing-api/model"
)

// GetResponse 数据报表API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData 返回值
type GetResponseData struct {
	// List 数据列表
	List []GetResponseList `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// GetResponseList 数据详情
type GetResponseList struct {
	Dimensions
	Metrics
}
