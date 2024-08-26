package thirdsite

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/thirdsite"
)

// Update 修改第三方落地页站点
// 通过此接口，用户可以修改第三方落地页站点名称name，修改成功后接口会返回"code_0"。 修改站点名称前后，站点id：site_id不变。
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *thirdsite.UpdateRequest) (uint64, error) {
	var resp thirdsite.UpdateResponse
	err := clt.Post(ctx, "2/tools/third_site/update/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	if resp.Data == nil {
		return 0, nil
	}
	return resp.Data.SiteID, nil
}
