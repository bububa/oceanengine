package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/tools/adraise/v3"
)

// StatusCurrentIDsGet 获取广告起量状态
// 批量获取广告当前的一键起量状态
func StatusCurrentIDsGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.StatusCurrentIDsGetRequest) (*v3.StatusCurrentIDsGetResult, error) {
	var resp v3.StatusCurrentIDsGetResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/promotion_status_current_ids/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
