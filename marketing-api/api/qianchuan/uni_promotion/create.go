package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// Create 新建全域推广计划
// 新建全域推广计划，当前仅支持「直播全域推广-控成本投放」链路
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.CreateRequest) (uint64, error) {
	var resp unipromotion.CreateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/uni_aweme/ad/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.AdID, nil
}
