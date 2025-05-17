package audiencepackage

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/audiencepackage"
)

// Create 创建定向包
// 落地页：可用于推广目的为销售线索收集或推广目的为应用推广且下载方式为落地页的计划
// 应用推广（Android）：可用于推广目的为应用推广且下载方式为Android下载链接的计划
// 应用推广（iOS）：可用于推广目的为应用推广且下载方式为iOS下载链接的计划
// 其余类型：可应用于推广目的为该类型名称的计划
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *audiencepackage.CreateRequest) (uint64, error) {
	var resp audiencepackage.CreateResponse
	err := clt.Post(ctx, "2/audience_package/create/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.AudiencePackageID, nil
}
