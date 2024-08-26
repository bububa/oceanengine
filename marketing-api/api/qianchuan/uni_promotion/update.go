package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// Update 编辑全域推广计划
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.UpdateRequest) (*unipromotion.UpdateResult, error) {
	var resp unipromotion.UpdateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/uni_aweme/ad/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
