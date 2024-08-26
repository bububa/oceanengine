package creative

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// Read 创意详细信息
func Read(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.ReadRequest) (*creative.CreativeDetail, error) {
	var resp creative.ReadResponse
	err := clt.Get(ctx, "2/creative/read_v2/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
