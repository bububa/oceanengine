package customaudience

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dmp/customaudience"
)

// Publish 发布人群包
// 使用发布接口可以将人群包进行发布。
// 发布人群包接口简单的说是一个让人群包在平台生效的操作，存在以下4种场景将用到发布人群包接口：
// 1.在创建数据源，人群包ID首次生成后，必须先使用发布人群包接口将人群包发布，才可使用；
// 2.如果更新了数据源，需要使用发布人群包接口才能把数据源的最新的内容发布（更新）到人群包中，发布人群包后数据源的最新内容将在人群包中生效，否则更新内容将无法生效；
// 3.如果是推送人群包，推送之后同样需要调用发布人群包接口才可将人群包生效;
// 4.如果是拓展人群包或者运算人群包，人群包ID生成后同样也需要调用发布人群包接口才可将人群包生效；人群包发布是一个异步的过程，将有一段处理的时间，你可以通过【人群包详细信息】查询人群包的状态，其中返回参数delivery_status表征当前人群包可投放状态（包含是否发布完成）；发布频次限制：每个人群包每24小时只可发布一次，发布后需在24小时后才可以再次发布此人群包。
//
// - 每当人群包的内容发生变化，就需要再次发布让变化生效！
// - 发布是一个异步的过程，发布中不建议对数据源进行更新操作，这会导致人群包的发布过程推迟！
func Publish(ctx context.Context, clt *core.SDKClient, accessToken string, req *customaudience.PublishRequest) error {
	return clt.Post(ctx, "2/dmp/custom_audience/publish/", req, nil, accessToken)
}
