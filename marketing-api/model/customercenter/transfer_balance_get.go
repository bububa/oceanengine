package customercenter

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TransferBalanceGetRequest 工作台转账-查询账户转账余额 API Request
type TransferBalanceGetRequest struct {
	// OrganizationID 组织id
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// TargetIDList 查询账户id列表(限制长度100)
	TargetIDList []uint64 `json:"target_id_list,omitempty"`
	// Platform 业务线 可选值:
	// AD 广告
	// BENDITUI 本地推
	Platform string `json:"platform,omitempty"`
}

// Encode implements GetRequest interface
func (r TransferBalanceGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("organization_id", strconv.FormatUint(r.OrganizationID, 10))
	values.Set("biz_request_no", r.BizRequestNo)
	values.Set("target_id_list", string(util.JSONMarshal(r.TargetIDList)))
	values.Set("platform", r.Platform)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// TransferBalanceGetResponse 工作台转账-查询账户转账余额 API Response
type TransferBalanceGetResponse struct {
	model.BaseResponse
	Data struct {
		// TargetAmountDetailList 账户金额列表
		TargetAmountDetailList []TargetAmountDetail `json:"target_amount_detail_list,omitempty"`
	} `json:"data,omitempty"`
}

// TargetAmountDetail 账户金额
type TargetAmountDetail struct {
	// TargetID 账户id
	TargetID uint64 `json:"target_id,omitempty"`
	// CapitalDetailList 可转资金列表
	CapitalDetailList []CapitalDetail `json:"capital_detail_list,omitempty"`
	// DepositAmount 竞价消耗保证金金额(单位：分)
	DepositAmount int64 `json:"deposit_amount,omitempty"`
	// TotalTransferAmount 总可转金额(单位：分)
	TotalTransferAmount int64 `json:"total_transfer_amount,omitempty"`
}
