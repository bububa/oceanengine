package customercenter

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/customercenter"
)

// FundTransferSeqCommit 提交转账交易号
// 在成功“创建交易号”后，通过提交，实际发起转账操作
func FundTransferSeqCommit(clt *core.SDKClient, accessToken string, req *customercenter.FundTransferSeqCommitRequest) (uint64, error) {
	var resp customercenter.FundTransferSeqCommitResponse
	if err := clt.Post("2/customer_center/fund/transfer_seq/commit/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.TransactionSeq, nil
}
