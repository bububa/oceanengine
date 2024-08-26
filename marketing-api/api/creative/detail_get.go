package creative

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// DetailGet 创意详细信息(新)
func DetailGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.DetailGetRequest) (*creative.CreativeDetailV2, error) {
	var resp creative.DetailGetResponse
	err := clt.GetAPI(ctx, "v3.0/creative/detail/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
