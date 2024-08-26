package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// AuthorizedGet 获取可投全域推广抖音号列表
func AuthorizedGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.AuthorizedGetRequest) (*unipromotion.AuthorizedGetResult, error) {
	var resp unipromotion.AuthorizedGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/uni_aweme/authorized/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
