package taskraise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/taskraise"
)

// Get 查询优选起量任务
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *taskraise.GetRequest) (*taskraise.GetResponseData, error) {
	var resp taskraise.GetResponse
	if err := clt.Get(ctx, "2/tools/task_raise/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
