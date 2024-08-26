package customaudience

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dmp/customaudience"
)

// Push 推送人群包
// 每个人群包生成一个人群包id后，都需要经过推送，才可以在被推送的广告主下使用。同时，推送人群包可以将人群包共享给同主体的广告主。
// 您可以通过【人群包详细信息】查询人群包的状态，其中返回参数delivery_status表征当前人群包可投放状态（包含是否推送完成）；
//
// - 推送的广告主列表的主体需要与人群包所属广告主的主体一致
// - 经过拓展/运算的人群包生成了新的id，意味着需要将新的id再次推送一遍!
// - 和人群包发布不同的是，如果用户基于数据源进行了新增/删除/重置操作，不会导致对应的人群包id变化，所以不需要再次推送!
func Push(ctx context.Context, clt *core.SDKClient, accessToken string, req *customaudience.PushRequest) error {
	return clt.Post(ctx, "2/dmp/custom_audience/push_v2/", req, nil, accessToken)
}
