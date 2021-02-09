package audience

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type ListResponse struct {
	model.BaseResponse
	Data *ListResponseData `json:"data,omitempty"`
}

type ListResponseData struct {
	List     []ListResponseList `json:"list,omitempty"`
	PageInfo *model.PageInfo    `json:"page_info,omitempty"`
}

type ListResponseList struct {
	LabelName         string             `json:"label_name,omitempty"`          // 标签名称
	SuperiorLabelName string             `json:"superior_label_name,omitempty"` // 上级标签名称，一级类目与关键词时返回null
	AudienceLevel     enum.AudienceLevel `json:"audience_level,omitempty"`      // 类目词级别
	Metrics           *Metrics           `json:"metrics,omitempty"`             // 查询指标列表
}
