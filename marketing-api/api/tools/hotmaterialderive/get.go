package hotmaterialderive

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/hotmaterialderive"
)

// Get 查询爆款裂变任务详情
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *hotmaterialderive.GetRequest) ([]hotmaterialderive.Task, error) {
	var resp hotmaterialderive.GetResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/hot_material_derive/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.TaskDetails, nil
}
