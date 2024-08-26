package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// DetailGet 获取计划详情（含创意信息）
// 用于获取千川广告账户下创建的单个计划的规则信息（含创意生成规则），以及计划下的创意
// 注意： 程序化创意在通过审核后才可以获取到
func DetailGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.DetailGetRequest) (*ad.Ad, error) {
	var resp ad.DetailGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/ad/detail/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
