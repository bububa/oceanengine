package quickappmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/quickappmanagement"
)

// QuickAppGet 查询快应用信息
// 查询当前广告主的快应用信息
func QuickAppGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *quickappmanagement.QuickAppGetRequest) (*quickappmanagement.QuickAppGetResult, error) {
	var resp quickappmanagement.QuickAppGetResponse
	if err := clt.Get(ctx, "2/tools/quick_app_management/quick_app/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
