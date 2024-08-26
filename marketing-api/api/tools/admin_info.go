package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// AdminInfo 获取行政信息
func AdminInfo(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.AdminInfoRequest) (*tools.AdminInfoResponseData, error) {
	var resp tools.AdminInfoResponse
	err := clt.Get(ctx, "2/tools/admin/info/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
