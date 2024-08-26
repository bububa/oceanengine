package smartphone

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/smartphone"
)

// Create 创建智能电话
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *smartphone.CreateRequest) (*smartphone.CreateResponseData, error) {
	var resp smartphone.CreateResponse
	if err := clt.Post(ctx, "2/clue/smartphone/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
