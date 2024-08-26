package rta

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rta"
)

// StatusUpdate 批量启停账户下RTA策略
// 修改RTA策略状态
func StatusUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *rta.StatusUpdateRequest) error {
	return clt.Post(ctx, "2/tools/rta/status_update/", req, nil, accessToken)
}
