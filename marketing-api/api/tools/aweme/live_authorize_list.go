package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/aweme"
)

// LiveAuthorizeList 查询授权直播抖音达人列表
// 授权直播抖音达人列表
func LiveAuthorizeList(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.LiveAuthorizeListRequest) (*aweme.LiveAuthorizeListResponseData, error) {
	var resp aweme.LiveAuthorizeListResponse
	err := clt.Get(ctx, "2/tools/live_authorize/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
