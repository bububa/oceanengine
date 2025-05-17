package taskraise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/taskraise"
)

// StatusStop 关闭优选起量任务
func StatusStop(ctx context.Context, clt *core.SDKClient, accessToken string, req *taskraise.StatusStopRequest) error {
	return clt.Post(ctx, "2/tools/task_raise/status/stop/", req, nil, accessToken)
}
