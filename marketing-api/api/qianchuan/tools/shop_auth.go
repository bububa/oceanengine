package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/tools"
)

// ShopAuth 店铺新客定向授权
// 店铺新客定向授权
func ShopAuth(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.ShopAuthRequest) error {
	return clt.PostAPI(ctx, "v1.0/qianchuan/tools/shop_auth/", req, nil, accessToken)
}
