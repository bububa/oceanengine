package taskraise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model/tools/taskraise"
)

// OptimizationIDsGet 查询优选起量状态
func OptimizationIDsGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (enum.TaskRaiseStatus, error) {
	var resp taskraise.OptimizationIDsGetResponse
	if err := clt.Get(ctx, "2/tools/task_raise/optimization_ids/get/", &taskraise.OptimizationIDsGetRequest{AdvertiserID: advertiserID}, &resp, accessToken); err != nil {
		return enum.TaskRaiseStatus_UNKNOWN, err
	}
	return resp.Data.Status, nil
}
