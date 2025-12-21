package hotmaterialderive

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/hotmaterialderive"
)

// List 获取账户下爆款裂变任务列表
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *hotmaterialderive.ListRequest) (*hotmaterialderive.ListResult, error) {
	var resp hotmaterialderive.ListResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/hot_material_derive/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
