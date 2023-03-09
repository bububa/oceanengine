package customercenter

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TransferType 转账类型
type TransferType string

const (
	// TransferType_GRANT 赠款
	TransferType_GRANT TransferType = "GRANT"
	// TransferType_PREPAY_UNIVERSAL 通用预付
	TransferType_PREPAY_UNIVERSAL TransferType = "PREPAY_UNIVERSAL"
	// TransferType_PREPAY_BRAND 品牌预付
	TransferType_PREPAY_BRAND TransferType = "PREPAY_BRAND"
	// TransferType_PREPAY_BID 竞价预付
	TransferType_PREPAY_BID TransferType = "PREPAY_BID"
	// TransferType_CREDIT_UNIVERSAL 通用授信
	TransferType_CREDIT_UNIVERSAL TransferType = "CREDIT_UNIVERSAL"
	// TransferType_CREDIT_BRAND 品牌授信
	TransferType_CREDIT_BRAND TransferType = "CREDIT_BRAND"
	// TransferType_CREDIT_BID 竞价授信
	TransferType_CREDIT_BID TransferType = "CREDIT_BID"
)

// FundTransferSeqCreateRequest 创建转账交易号 API Request
type FundTransferSeqCreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TargetAdvertiserID 转入账户
	TargetAdvertiserID uint64 `json:"target_advertiser_id,omitempty"`
	// Amount 转账金额
	Amount float64 `json:"amount,omitempty"`
	// TransferType 转账类型
	TransferType TransferType `json:"transfer_type,omitempty"`
}

// Encode implement PostRequest interface
func (r FundTransferSeqCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// FundTransferSeqCreateResponse 创建转账交易号 API Response
type FundTransferSeqCreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// TransactionSeq 交易流水号
		TransactionSeq uint64 `json:"transaction_seq,omitempty"`
	} `json:"data,omitempty"`
}
