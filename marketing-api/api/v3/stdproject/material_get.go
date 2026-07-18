package stdproject

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/stdproject"
)

// MaterialGet 获取投放项目下素材
func MaterialGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *stdproject.MaterialGetRequest) (*stdproject.MaterialGetResponseData, error) {
	var resp stdproject.MaterialGetResponse
	if err := clt.GetAPI(ctx, "v3.0/oc_project/material/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
