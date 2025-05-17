package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/report/v3"
)

// CustomConfigGet 获取自定义报表可用指标和维度
func CustomConfigGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.CustomConfigGetRequest) ([]v3.CustomConfig, error) {
	var resp v3.CustomConfigGetResponse
	if err := clt.GetAPI(ctx, "v3.0/report/custom/config/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
