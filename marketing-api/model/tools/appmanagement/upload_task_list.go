package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UploadTaskListRequest 查询文件异步上传任务 API Request
type UploadTaskListRequest struct {
	// AccountID 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountTYpe 账户类型 可选值:
	// AD 广告主类型、BP 巨量纵横账号类型
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// Filtering 过滤条件
	Filtering *UploadTaskListFilter `json:"filtering,omitempty"`
}

type UploadTaskListFilter struct {
	// UploadIDs 上传文件id，通过【创建异步文件上传任务】获得
	UploadIDs []uint64 `json:"upload_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r UploadTaskListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	if r.AccountType != "" {
		values.Set("account_type", string(r.AccountType))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// UploadTaskListResponse 查询文件异步上传任务 API Response
type UploadTaskListResponse struct {
	model.BaseResponse
	Data struct {
		// List 任务列表
		List []UploadTask `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
