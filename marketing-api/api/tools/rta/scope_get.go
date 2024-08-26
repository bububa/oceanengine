package rta

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rta"
)

// ScopeGet 获取RTA策略绑定信息列表
// 获取RTA策略下的项目(广告组)列表，一个rta策略要么绑定campaign要么绑定project，不可能同时绑定
func ScopeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *rta.ScopeGetRequest) ([]rta.Scope, error) {
	var resp rta.ScopeGetResponse
	if err := clt.Get(ctx, "2/tools/rta/scope/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
