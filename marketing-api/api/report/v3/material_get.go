package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/report/v3"
)

// MaterialGet 素材数据报表
func MaterialGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.MaterialGetRequest) (*v3.MaterialGetResult, error) {
	var resp v3.MaterialGetResponse
	if err := clt.GetAPI(ctx, "v3.0/report/material/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
