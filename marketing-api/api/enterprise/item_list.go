package enterprise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// ItemList 获取企业号视频列表
func ItemList(ctx context.Context, clt *core.SDKClient, accessToken string, req *enterprise.ItemListRequest) ([]enterprise.Item, error) {
	var resp enterprise.ItemListResponse
	if err := clt.Get(ctx, "v1.0/enterprise/item/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ItemList, nil
}
