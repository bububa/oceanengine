package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// Get 获取账户下计划列表（不含创意）
// 用于获取千川广告账户下创建的计划列表（不含创意部分） 默认可以拉取全部计划详情，可以通过过滤拉取部分计划
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.GetRequest) (*ad.GetResponseData, error) {
	var resp ad.GetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/ad/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
