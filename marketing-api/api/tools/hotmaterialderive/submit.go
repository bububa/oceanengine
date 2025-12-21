package hotmaterialderive

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/hotmaterialderive"
)

// Submit 提交爆款裂变任务
func Submit(ctx context.Context, clt *core.SDKClient, accessToken string, req *hotmaterialderive.SubmitRequest) ([]hotmaterialderive.Task, error) {
	var resp hotmaterialderive.SubmitResponse
	if err := clt.PostAPI(ctx, "v3.0/tools/hot_material_derive/submit/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Tasks, nil
}
