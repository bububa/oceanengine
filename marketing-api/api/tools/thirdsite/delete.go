package thirdsite

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/thirdsite"
)

// Delete 删除第三方落地页站点
// 通过此接口，用户可以删除第三方落地页站点。
func Delete(ctx context.Context, clt *core.SDKClient, accessToken string, req *thirdsite.DeleteRequest) (uint64, error) {
	var resp thirdsite.DeleteResponse
	err := clt.Post(ctx, "2/tools/third_site/delete/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	if resp.Data == nil {
		return 0, nil
	}
	return resp.Data.SiteID, nil
}
