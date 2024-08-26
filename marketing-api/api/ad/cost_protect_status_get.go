package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/ad"
)

// CostProtectStatusGet 批量获取计划成本保障状态
// 获取成本保障状态及时同步广告主计划投放状态，可以获取计划是否处于保障中，满足赔付规则
// 接口支持批量查询
// 赔付状态：
// 成本保障生效中：表示您的计划处于赔付的时间范围，且没有做出影响赔付的操作（每天修改广告计划出价或定向中任意一个的次数超过2次），可以放心投放
// 成本保障已失效：表示在赔付的时间范围内，您每天修改广告计划出价或定向中任意一个的次数超过了2次，已经不是赔付的候选计划，系统也不会计算赔付金额
// 成本保障计算中：计划开始投放后的第5-7个自然日，系统正在对候选赔付计划，计算是否赔付以及赔付金额
// 成本保障已到账：该计划在投放期间，满足自动赔付规则，赔付已到账，可在站内信看到到账通知
// 成本保障已结束：在系统计算后超成本幅度或转化数未达到赔付标准
// 若广告主存在异常作弊行为，平台有权拒绝赔付并追回已赔付款项
// 不适用于 ROI 出价、每次付费出价、自动出价、投放管家的广告计划
// 赔付金额和实际消耗有关，总实际消耗中不包括使用「一键起量」工具期间产生的消耗
func CostProtectStatusGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.CostProtectStatusGetRequest) ([]ad.CostProtectStatus, error) {
	var resp ad.CostProtectStatusGetResponse
	if err := clt.Get(ctx, "2/ad/cost_protect_status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
