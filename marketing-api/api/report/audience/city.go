package audience

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/audience"
)

// City 市级数据
// 获取城市定向受众数据，包括你所投放城市的总花费、展示数、点击数等数据。受众分析报表数据不支持获取当天的数据。
// 受众数据不支持获取当天的数据；
func City(ctx context.Context, clt *core.SDKClient, accessToken string, req *audience.Request) ([]audience.ResponseData, error) {
	var resp audience.Response
	err := clt.Get(ctx, "2/report/audience/city/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
