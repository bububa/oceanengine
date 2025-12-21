package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// AuthInit 全域授权初始化
// 全域授权初始化接口，用以解决部分在拉取达人/机构广告主可投的商品列表时返回商品不可投case（当前账户无使用该抖音号投放所选商品的全域推广权限）。
func AuthInit(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.AuthInitRequest) error {
	return clt.Post(ctx, "v1.0/qianchuan/uni_promotion/auth/init/", req, nil, accessToken)
}
