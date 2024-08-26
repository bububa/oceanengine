package creative

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// MaterialRead 创意素材信息
func MaterialRead(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.MaterialReadRequest) ([]creative.Material, error) {
	var resp creative.MaterialReadResponse
	err := clt.Get(ctx, "2/creative/material/read/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
