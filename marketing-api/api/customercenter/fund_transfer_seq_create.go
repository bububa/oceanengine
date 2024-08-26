package customercenter

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/customercenter"
)

// FundTransferSeqCreate 创建转账交易号
// 创建转账交易号，即发起转账请求
// 可以从客户中心账户往广告主账户转账，也可以广告主之间转账
// 在执行转账的时候动态检查保证金，保证金不是在某个时间点确定扣除的
// 不影响现有资金流水接口（拆成分步接口后），转账操作完成后，流水产生时间是对应转账成功的时间
func FundTransferSeqCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *customercenter.FundTransferSeqCreateRequest) (uint64, error) {
	var resp customercenter.FundTransferSeqCreateResponse
	if err := clt.Post(ctx, "2/customer_center/fund/transfer_seq/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.TransactionSeq, nil
}
