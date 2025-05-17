package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/report/v3"
)

// ProjectGet 项目数据报表
func ProjectGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.ProjectGetRequest) (*v3.ProjectGetResult, error) {
	var resp v3.ProjectGetResponse
	if err := clt.GetAPI(ctx, "v3.0/report/project/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
