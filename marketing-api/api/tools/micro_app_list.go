package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// MicroAppList 获取字节小程序
// 获取巨量工作台上字节小程序资产列表
func MicroAppList(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.MicroAppListRequest) (*tools.MicroAppListResult, error) {
	var resp tools.MicroAppListResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/micro_app/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
