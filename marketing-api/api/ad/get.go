package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/ad"
)

// Get 获取广告计划
// 此接口用于获取广告计划列表的信息；
// 如果创建计划时未设置对应的字段，返回的字段值会是null
// 支持filtering过滤，可按广告计划ID、出价方式、广计划状态进行过滤筛选
// 默认不获取删除的计划，如果要获取删除的计划，可在filtering中传入对应的status值；
// 对于搜索广告计划信息获取参见【搜索广告投放】
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.GetRequest) (*ad.GetResponseData, error) {
	var resp ad.GetResponse
	err := clt.Get(ctx, "2/ad/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
