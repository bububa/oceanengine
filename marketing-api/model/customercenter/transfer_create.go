package customercenter

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TransferCreateRequest 工作台转账-发起转账 API Request
type TransferCreateRequest struct {
	// OrganizationID 组织id
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// OpponentTargetID 锚定账户id，1:N的1
	OpponentTargetID uint64 `json:"opponent_target_id,omitempty"`
	// TargetDetailList 目标账户列表，1:N的N，需要列表内账户类型相同，最多支持100个
	TargetDetailList []TransferTarget `json:"target_detail_list,omitempty"`
	// TransferDirection 转账方向，以目标账户视角确定 可选值:
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	TransferDirection enum.TransferDirection `json:"transfer_direction,omitempty"`
	// Remark 备注
	Remark string `json:"remark,omitempty"`
	// Platform 转账业务线 可选值:
	// AD 广告
	// BENDITUI 本地推
	Platform string `json:"platform,omitempty"`
}

// Encode implements PostRequest interface
func (r TransferCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TransferCreateResponse 工作台转账-发起转账
type TransferCreateResponse struct {
	model.BaseResponse
	Data struct {
		// TransferSerial 转账单号
		TransferSerial string `json:"transfer_serial,omitempty"`
	} `json:"data,omitempty"`
}
