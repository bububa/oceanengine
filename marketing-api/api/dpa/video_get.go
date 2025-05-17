package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// VideoGet 获取 DPA 商品库视频模板
func VideoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.VideoGetRequest) (*dpa.VideoGetResponseData, error) {
	var resp dpa.VideoGetResponse
	if err := clt.Get(ctx, "2/dpa/video/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
