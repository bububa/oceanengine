package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// List 获取全域推广列表
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.ListRequest) (*unipromotion.ListResult, error) {
	var resp unipromotion.ListResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/uni_promotion/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
