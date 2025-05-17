package landinggroup

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/landinggroup"
)

// Create 创建落地页组
// 通过此接口，用户可以创建按照流量分配的落地页组，创建成功后，接口会返回"code_0"。
// 落地页组信息会包括落地页组 ID、落地页组名称、落地页组 URL、落地页组状态、流量分配方式； 站点列表信息会包括成员 ID（即站点在落地页组中的唯一标识）、站点 ID、站点URL、站点启用状态、站点审核状态。
// 落地页组名称长度为1-20，否则无法创建成功！
// 橙子建站站点id列表的数量范围：2 <= site数量 <= 10
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *landinggroup.CreateRequest) (*landinggroup.LandingGroup, error) {
	var resp landinggroup.CreateResponse
	if err := clt.Post(ctx, "2/tools/landing_group/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
