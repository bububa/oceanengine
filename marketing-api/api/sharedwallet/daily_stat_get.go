package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// DailyStatGet 资金共享-查询共享钱包日流水
// 资金共享-查询共享钱包日流水信息，一般每天8点会出来前一天的数据，如果当天存在数据延迟可往后再尝试。
func DailyStatGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.DailyStatGetRequest) (*sharedwallet.DailyStatGetResult, error) {
	var resp sharedwallet.DailyStatGetResponse
	if err := clt.GetAPI(ctx, "v3.0/shared_wallet/daily_stat/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
