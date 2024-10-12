package clue

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/clue"
)

// LifeCallback 本地推线索回传
func LifeCallback(ctx context.Context, clt *core.SDKClient, accessToken string, req *clue.LifeCallbackRequest) error {
	return clt.PostAPI(ctx, "2/tools/clue/life/callback/", req, nil, accessToken)
}
