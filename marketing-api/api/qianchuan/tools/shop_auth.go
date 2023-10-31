package tools

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/tools"
)

// ShopAuth 店铺新客定向授权
// 店铺新客定向授权
func ShopAuth(clt *core.SDKClient, accessToken string, req *tools.ShopAuthRequest) error {
	return clt.Post("v1.0/qianchuan/tools/shop_auth/", req, nil, accessToken)
}
