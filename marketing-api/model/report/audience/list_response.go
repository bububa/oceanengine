package audience

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// ListResponse 行为兴趣数据/抖音达人数据 API Response
type ListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *ListResponseData `json:"data,omitempty"`
}

// ListResponseData json返回值
type ListResponseData struct {
	// List 数据列表
	List []ListResponseList `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// ListResponseList 数据
type ListResponseList struct {
	// LabelName 标签名称
	LabelName string `json:"label_name,omitempty"`
	// SuperiorLabelName 上级标签名称，一级类目与关键词时返回null
	SuperiorLabelName string `json:"superior_label_name,omitempty"`
	// AudienceLevel 类目词级别
	AudienceLevel enum.AudienceLevel `json:"audience_level,omitempty"`
	// Metrics 查询指标列表
	Metrics *Metrics `json:"metrics,omitempty"`
}
