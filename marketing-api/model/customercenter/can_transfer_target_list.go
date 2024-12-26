package customercenter

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CanTransferTargetListRequest 工作台转账-获取可转列表 API Request
type CanTransferTargetListRequest struct {
	// OrganizationID 组织id
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// OpponentTargetID 锚定账户id，查询该账户的可转账账户列表
	OpponentTargetID uint64 `json:"opponent_target_id,omitempty"`
	// TransferDirection 转账方向，以可转列表视角确定 可选值:
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	TransferDirection enum.TransferDirection `json:"transfer_direction,omitempty"`
	// Platform 业务线 可选值:
	// AD 广告
	// BENDITUI 本地推
	Platform string `json:"platform,omitempty"`
	// Page 页码,从1开始
	Page int `json:"page,omitempty"`
	// PageSize 每页最多100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r CanTransferTargetListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("organization_id", strconv.FormatUint(r.OrganizationID, 10))
	values.Set("biz_request_no", r.BizRequestNo)
	values.Set("opponent_target_id", strconv.FormatUint(r.OpponentTargetID, 10))
	values.Set("transfer_direction", string(r.TransferDirection))
	values.Set("platform", r.Platform)
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

// CanTransferTargetListResponse 工作台转账-获取可转列表 API Response
type CanTransferTargetListResponse struct {
	model.BaseResponse
	Data *CanTransferTargetListResult `json:"data,omitempty"`
}

type CanTransferTargetListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// CanTransferTargetList 可转账户列表
	CanTransferTargetList []TransferTarget `json:"can_transfer_target_list,omitempty"`
}

// TransferTarget 可转账户
type TransferTarget struct {
	// TargetID 可转账户id
	TargetID uint64 `json:"target_id,omitempty"`
	// TransferCaptialDetailList 	锚定账户与目标账户转账资金列表
	TransferCaptialDetailList []CapitalDetail `json:"transfer_captial_detail_list,omitempty"`
}
