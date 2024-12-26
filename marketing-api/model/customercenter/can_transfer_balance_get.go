package customercenter

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CanTransferBalanceGetRequest 工作台转账-获取最大可转余额 API Request
type CanTransferBalanceGetRequest struct {
	// OrganizationID 组织id
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// OpponentTargetID 锚定账户id，查询该账户的可转账账户列表
	OpponentTargetID uint64 `json:"opponent_target_id,omitempty"`
	// TargetIDs 目标账户id列表，1:N的N，最多支持100个
	TargetIDs []uint64 `json:"target_ids,omitempty"`
	// TransferDirection 转账方向，以可转列表视角确定 可选值:
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	TransferDirection enum.TransferDirection `json:"transfer_direction,omitempty"`
	// Platform 业务线 可选值:
	// AD 广告
	// BENDITUI 本地推
	Platform string `json:"platform,omitempty"`
}

// Encode implements GetRequest interface
func (r CanTransferBalanceGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("organization_id", strconv.FormatUint(r.OrganizationID, 10))
	values.Set("biz_request_no", r.BizRequestNo)
	values.Set("opponent_target_id", strconv.FormatUint(r.OpponentTargetID, 10))
	values.Set("target_ids", string(util.JSONMarshal(r.TargetIDs)))
	values.Set("transfer_direction", string(r.TransferDirection))
	values.Set("platform", r.Platform)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CanTransferBalanceGetResponse 工作台转账-获取最大可转余额 API Response
type CanTransferBalanceGetResponse struct {
	model.BaseResponse
	Data struct {
		// CanTransferDetailList 可转余额信息列表
		CanTransferDetailList []CanTransferDetail `json:"can_transfer_detail_list,omitempty"`
	} `json:"data,omitempty"`
}

// CanTransferDetail 可转余额信息
type CanTransferDetail struct {
	// RemitterTargetID 转出方账户id
	RemitterTargetID uint64 `json:"remitter_target_id,omitempty"`
	// PayeeTransferAmountDetailList 转入方可转余额信息列表
	PayeeTransferAmountDetailList []PayeeTransferAmountDetail `json:"payee_transfer_amount_detail_list,omitempty"`
	// CapitalDetailList 转出方可转资金列表
	CapitalDetailList []CapitalDetail `json:"capital_detail_list,omitempty"`
}

// PayeeTransferAmountDetail 转入方可转余额信息
type PayeeTransferAmountDetail struct {
	// PayeeTargetID 转入方账户id
	PayeeTargetID uint64 `json:"payee_target_id,omitempty"`
	// CapitalDetailList 转入方可转资金列表
	CapitalDetailList []CapitalDetail `json:"capital_detail_list,omitempty"`
}

// CapitalDetail 转入方可转资金
type CapitalDetail struct {
	// CapitalType 转入方可转资金类型 可选值:
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
	// TransferBalance 转入方可转资金余额(单位：分)
	TrainsferBalance int64 `json:"transfer_balance,omitempty"`
}
