package report

import (
	"github.com/bububa/oceanengine/marketing-api/model"
)

type GetResponse struct {
	model.BaseResponse
	Data *GetResponseData `json:"data,omitempty"`
}

type GetResponseData struct {
	List     []GetResponseList `json:"list,omitempty"`
	PageInfo *model.PageInfo   `json:"page_info,omitempty"`
}

type GetResponseList struct {
	Dimensions
	Metrics
}
