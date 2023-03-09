package customercenter

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FundTransferSeqCommitRequest 提交转账交易号 API Request
type FundTransferSeqCommitRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TargetAdvertiserID 转入账户
	TargetAdvertiserID uint64 `json:"target_advertiser_id,omitempty"`
	// TransactionSeq 交易流水号
	TransactionSeq uint64 `json:"transaction_seq,omitempty"`
}

// Encode implement PostRequest interface
func (r FundTransferSeqCommitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// FundTransferSeqCommitResponse 提交转账交易号 API Response
type FundTransferSeqCommitResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// TransactionSeq 交易流水号
		TransactionSeq uint64 `json:"transaction_seq,omitempty"`
	} `json:"data,omitempty"`
}
