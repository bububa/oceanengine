package stdproject

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/stdproject"
)

// List 获取标准项目列表
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *stdproject.ListRequest) (*stdproject.ListResponseData, error) {
	var resp stdproject.ListResponse
	if err := clt.GetAPI(ctx, "v3.0/std_project/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
