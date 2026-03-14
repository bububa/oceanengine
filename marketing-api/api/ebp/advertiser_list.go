package ebp

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/ebp"
)

// AdvertiserList 获取升级版巨量引擎工作台下账户列表
func AdvertiserList(ctx context.Context, clt *core.SDKClient, accessToken string, req *ebp.AdvertiserListRequest) (*ebp.AdvertiserListResponseData, error) {
	var resp ebp.AdvertiserListResponse
	if err := clt.Get(ctx, "2/ebp/advertiser/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
