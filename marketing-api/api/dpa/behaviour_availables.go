package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// BehaviourAvailables 获取DPA可用行为
func BehaviourAvailables(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.BehaviourAvailablesRequest) ([]dpa.Behaviour, error) {
	var resp dpa.BehaviourAvailablesResponse
	if err := clt.Get(ctx, "2/dpa/behaviour/availables/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
