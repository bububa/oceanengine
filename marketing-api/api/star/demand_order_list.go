package star

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/star"
)

// DemandOrderList 获取星图客户任务订单列表
// 此接口用于根据星图id和星图任务id，获取星图客户所发布任务的订单列表，包含订单总数、订单id、订单创建时间、订单状态、作品名称、作品封面图、视频id、视频链接、接单的达人id、达人名称、达人头像。
func DemandOrderList(ctx context.Context, clt *core.SDKClient, accessToken string, req *star.DemandOrderListRequest) (*star.DemandOrderListResult, error) {
	var resp star.DemandOrderListResponse
	err := clt.GetAPI(ctx, "2/star/demand/order/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
