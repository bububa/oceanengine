package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// AdNameUpdate 更新商品全域推广计划名称
func AdNameUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.AdNameUpdateRequest) error {
	return clt.Post(ctx, "v1.0/qianchuan/uni_promotion/ad/name/update/", req, nil, accessToken)
}
