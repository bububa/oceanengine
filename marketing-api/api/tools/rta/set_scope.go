package rta

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rta"
)

// SetScope 设置账户下RTA策略生效范围
// 设置广告账户下某个RTA策略的生效范围，设置一个新的RTA策略时，默认为停用状态，需要继续调用修改RTA策略状态
func SetScope(ctx context.Context, clt *core.SDKClient, accessToken string, req *rta.SetScopeRequest) error {
	return clt.Post(ctx, "2/tools/rta/set_scope/", req, nil, accessToken)
}
