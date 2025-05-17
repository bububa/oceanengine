package servemarket

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/servemarket"
)

// ActiveFuncGet 获取用户已购功能点列表
func ActiveFuncGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *servemarket.ActiveFuncGetRequest) ([]servemarket.OrderFunction, error) {
	var resp servemarket.ActiveFuncGetResponse
	if err := clt.OpenGet(ctx, "v1.0/serve_market/active_func/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.FuncList, nil
}
