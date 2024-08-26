package landinggroup

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/landinggroup"
)

// SiteOptStatusUpdate 更新落地页组站点状态
// 通过此接口，用户可以修改落地页组站点的启用状态。
func SiteOptStatusUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *landinggroup.SiteOptStatusUpdateRequest) error {
	return clt.Post(ctx, "2/tools/landing_group/site_opt_status/update/", req, nil, accessToken)
}
