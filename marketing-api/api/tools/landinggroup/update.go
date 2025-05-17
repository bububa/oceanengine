package landinggroup

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/landinggroup"
)

// Update 更新落地页组信息
// 通过此接口，用户可以更新落地页组的基本信息。
// 落地页组信息会包括落地页组 ID、落地页组名称、落地页组 URL、落地页组状态、流量分配方式； 站点列表信息会包括成员 ID（即站点在落地页组中的唯一标识）、站点 ID、站点URL、站点启用状态、站点审核状态。
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *landinggroup.UpdateRequest) (*landinggroup.LandingGroup, error) {
	var resp landinggroup.UpdateResponse
	if err := clt.Post(ctx, "2/tools/landing_group/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
