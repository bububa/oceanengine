package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/audiencepackage/v3"
)

// Get 获取定向包
// 落地页：可用于推广目的为销售线索收集或推广目的为应用推广且下载方式为落地页的计划
// 应用推广（Android）：可用于推广目的为应用推广且下载方式为Android下载链接的计划
// 应用推广（iOS）：可用于推广目的为应用推广且下载方式为iOS下载链接的计划
// 其余类型：可应用于推广目的为该类型名称的计划
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.GetRequest) (*v3.GetResult, error) {
	var resp v3.GetResponse
	if err := clt.GetAPI(ctx, "v3.0/audience_package/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
