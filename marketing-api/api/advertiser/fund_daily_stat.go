package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// FundDailyStat 查询账号日流水
// 获取广告主日流水信息，一般每天8点会出来前一天的数据，如果当天存在数据延迟可往后再尝试。
func FundDailyStat(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.FundDailyStatRequest) (*advertiser.FundDailyStatResponseData, error) {
	var resp advertiser.FundDailyStatResponse
	err := clt.Get(ctx, "2/advertiser/fund/daily_stat/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
