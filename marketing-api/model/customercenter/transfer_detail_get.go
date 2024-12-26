package customercenter

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TransferDetailGetRequest 工作台转账-查询转账单信息 API Request
type TransferDetailGetRequest struct {
	// OrganizationID 组织id
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// TransferBizRequestNo 发起转账的幂等id
	TransferBizRequestNo string `json:"transfer_biz_request_no,omitempty"`
	// TransferSerial 转账单号，与transfer_biz_request_no两者传其一即可
	TransferSerial string `json:"transfer_serial,omitempty"`
	// Platform 转账业务线 可选值:
	// AD 广告
	// BENDITUI 本地推
	Platform string `json:"platform,omitempty"`
}

// Encode implements GetRequest interface
func (r TransferDetailGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("organization_id", strconv.FormatUint(r.OrganizationID, 10))
	values.Set("biz_request_no", r.BizRequestNo)
	values.Set("transfer_biz_request_no", r.TransferBizRequestNo)
	values.Set("transfer_serial", r.TransferSerial)
	values.Set("platform", r.Platform)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// TransferDetailGetResponse 工作台转账-查询转账单信息 API Response
type TransferDetailGetResponse struct {
	model.BaseResponse
	Data *TransferDetail `json:"data,omitempty"`
}

type TransferDetail struct {
	// TransferSerial 转账单号
	TransferSerial string `json:"transfer_serial,omitempty"`
	// BizRequestNo 幂等id
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// TransferDirection 转账方向(以目标账户视角确定) 可选值:
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	TransferDirection enum.TransferDirection `json:"transfer_direction,omitempty"`
	// TransferAmount 转账总金额(单位：分)
	TransferAmount int64 `json:"transfer_amount,omitempty"`
	// TransferStatus 转账总状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败(终态)
	// TRANSFER_ING 转账中
	// TRANSFER_PART 转账部分成功(终态)
	// TRANSFER_SUCCESS 转账成功(终态)
	TransferStatus enum.TransferStatus `json:"transfer_status,omitempty"`
	// TransferFinishTime 转账完成时间(yyyy-MM-dd HH:mm:ss)
	TransferFinishTime string `json:"transfer_transfer_time,omitempty"`
	// TransferCreateTime 转账创建时间(yyyy-MM-dd HH:mm:ss)
	TransferCreateTime string `json:"transfer_create_time,omitempty"`
	// TransferTargetRecordList 账户信息列表
	TransferTargetRecordList []TransferTargetRecord `json:"transfer_target_record_list,omitempty"`
}

// TransferTargetRecord 账户信息
type TransferTargetRecord struct {
	// OpponentTargetID 锚定账户id，1:N的1
	OpponetTargetID uint64 `json:"opponent_target_id,omitempty"`
	// TargetID 目标账户id，1:N的N
	TargetID uint64 `json:"target_id,omitempty"`
	// TransferAmount 转账金额(单位：分)
	TransferAmount int64 `json:"transfer_amount,omitempty"`
	// TransferStatus 转账总状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败(终态)
	// TRANSFER_ING 转账中
	// TRANSFER_PART 转账部分成功(终态)
	// TRANSFER_SUCCESS 转账成功(终态)
	TransferStatus enum.TransferStatus `json:"transfer_status,omitempty"`
	// TransferCapitalRecordList 转账资金类型列表
	TransferCapitalRecordList []TransferCapitalRecord `json:"transfer_capital_record_list,omitempty"`
}

// TransferCapitalRecord 转账资金类型
type TransferCapitalRecord struct {
	// CapitalType 转账资金类型 可选值:
	// CREDIT_BIDDING 授信竞价
	// CREDIT_BRAND 授信品牌
	// CREDIT_GENERAL 授信通用
	// GRANT_COMMON 信息流赠款
	// GRANT_DEFAULT 通用赠款
	// GRANT_SEARCH 搜索赠款
	// GRANT_UNION 穿山甲赠款
	// PREPAY_BIDDING 预付竞价
	// PREPAY_BRAND 预付品牌
	// PREPAY_GENERAL 预付通用
	CapitalType enum.CapitalType `json:"capital_type,omitempty"`
	// TransferAmount 转账资金金额(单位：分)
	TransferAmount int64 `json:"transfer_amount,omitempty"`
	// TansferStatus 转账资金状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败(终态)
	// TRANSFER_ING 转账中
	// TRANSFER_SUCCESS 转账成功(终态)        r
	TransferStatus enum.TransferStatus `json:"transfer_status,omitempty"`
	// FailReason 失败原因
	FailReason string `json:"fail_reason,omitempty"`
}
