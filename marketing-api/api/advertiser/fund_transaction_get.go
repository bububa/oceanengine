package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// FundTransactionGet 查询账号流水明细
// 获取广告主或代理商账户流水明细信息。接口规则如下：
// 1.假设某广告主A是代理商下广告主，对于A来说，如果流水类型是转账，而收款方是A则意味着是转账加款，如果付款方是A则意味着是转账退款。
// 2.假设某广告主B是直客广告主，未在任何代理商下，如果交易流水类型为充值，则收款方和付款方将都显示广告主B，如果金额为正则是充值交易，金额为负则是退款交易。
func FundTransactionGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.FundTransactionGetRequest) ([]advertiser.FundTransactionGetResponseList, error) {
	var resp advertiser.FundTransactionGetResponse
	err := clt.Get(ctx, "2/advertiser/fund/transaction/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.List, nil
}
