package carousel

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/carousel"
)

// AwemeGet 获取抖音号下图文
// 通过商品id获取抖音号下已有的图文素材，当前仅短视频带货支持，必须传商品id
func AwemeGet(ctx context.Context, clt *core.SDKClient, req *carousel.AwemeGetRequest, accessToken string) (*carousel.AwemeGetResult, error) {
	var resp carousel.AwemeGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/carousel/aweme/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
