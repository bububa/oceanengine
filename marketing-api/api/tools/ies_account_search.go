package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// IesAccountSearch 获取绑定的抖音号
// 此接口可以帮助查询广告主账户当前绑定的抖音号信息
func IesAccountSearch(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.IesAccountSearchRequest) ([]tools.IesAccount, error) {
	var resp tools.IesAccountSearchResponse
	if err := clt.Get(ctx, "2/tools/ies_account_search/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
