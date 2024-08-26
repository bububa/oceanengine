package taskraise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/taskraise"
)

// Create 账户优选起量
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *taskraise.CreateRequest) (uint64, error) {
	var resp taskraise.CreateResponse
	if err := clt.Post(ctx, "2/tools/task_raise/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ReportID, nil
}
