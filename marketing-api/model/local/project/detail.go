package project

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailRequest 获取项目详情 API Request
type DetailRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
}

// Encode implements GetRequest interface
func (r DetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	values.Set("project_id", strconv.FormatUint(r.ProjectID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DetailResponse 获取项目详情 API Response
type DetailResponse struct {
	// Project 项目详情
	Data *ProjectDetail `json:"data,omitempty"`
	model.BaseResponse
}
