package adraise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adraise"
)

// Result 获取一键起量的后验数据: 用来获取一键起量的后验数据
func Result(ctx context.Context, clt *core.SDKClient, accessToken string, req *adraise.ResultRequest) (*adraise.Result, error) {
	var resp adraise.ResultResponse
	err := clt.Get(ctx, "2/tools/ad_raise_result/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
