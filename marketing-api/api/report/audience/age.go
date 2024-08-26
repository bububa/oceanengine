package audience

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/audience"
)

// Age 年龄数据
// 获取年龄定向受众数据，当前包括6个年龄段（1-18；18-23；24-30；31-40；41-50；50岁以上），可以查看受众中不同年龄段用户的投放数据，包括花费、展示、点击等数据。
// 受众数据不支持获取当天的数据；
func Age(ctx context.Context, clt *core.SDKClient, accessToken string, req *audience.Request) ([]audience.ResponseData, error) {
	var resp audience.Response
	err := clt.Get(ctx, "2/report/audience/age/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
