package campaign

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/campaign"
)

// Get 获取广告组
// 此接口用于获取广告组列表的信息；
// 当预算类型为不限，返回的预算为0.0元
// 支持filtering过滤，可按广告组ID、推广目的、广告组状态进行过滤筛选
// 默认不获取删除的广告组，如果要获取删除的广告组，可在filtering中传入对应的status值；
// 对于搜索广告组信息获取参见【搜索广告投放】
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.GetRequest) (*campaign.GetResponseData, error) {
	var resp campaign.GetResponse
	err := clt.Get(ctx, "2/campaign/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
