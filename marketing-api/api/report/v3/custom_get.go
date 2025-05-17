package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/report/v3"
)

// CustomGet 自定义数据报表
func CustomGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.CustomGetRequest) (*v3.CustomGetResult, error) {
	var resp v3.CustomGetResponse
	if err := clt.GetAPI(ctx, "v3.0/report/custom/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
