package agent

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdvertiserRechargeRequest 代理商转账 API Request
type AdvertiserRechargeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AgentID 代理商ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// TransferType 转账类型，允许值：
	// GRANT 赠款
	// PREPAY_UNIVERSAL 通用预付
	// PREPAY_BRAND 品牌预付
	// PREPAY_BID 竞价预付
	// CREDIT_UNIVERSAL 通用授信
	// CREDIT_BRAND 品牌授信
	// CREDIT_BID 竞价授信
	TransferType AdvertiserTransferType `json:"transfer_type,omitempty"`
	// Amount 金额,单位(元),最低转账金额500元
	Amount float64 `json:"amount,omitempty"`
}

// Encode implement PostRequest interface
func (r AdvertiserRechargeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdvertiserRechargeResponse 代理商转账 API Response
type AdvertiserRechargeResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdvertiserRechargeResponseData `json:"data,omitempty"`
}

// AdvertiserRechargeResponseData json返回值
type AdvertiserRechargeResponseData struct {
	// TransactionSeq 交易序列号
	TransactionSeq string `json:"transaction_seq,omitempty"`
}
