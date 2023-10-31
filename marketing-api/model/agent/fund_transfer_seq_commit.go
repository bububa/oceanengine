package agent

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FundTransferSeqCommitRequest 提交转账交易号（方舟） API Request
type FundTransferSeqCommitRequest struct {
	// AgentID 代理商账户ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// TransferSeq 转账序列号
	TransferSeq string `json:"transfer_seq,omitempty"`
}

// Encode implement PostRequest interface
func (r FundTransferSeqCommitRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
