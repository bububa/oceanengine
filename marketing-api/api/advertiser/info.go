package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// Info 广告主信息
// 获取广告主账户详细信息,可指定fields查询所需元素。
// 目前上线了新版的代理商鉴权，如果查询报no permission错误，请前往代理商一站式平台为账号申请对应的权限。
func Info(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.InfoRequest) ([]advertiser.Info, error) {
	var resp advertiser.InfoResponse
	err := clt.Get(ctx, "2/advertiser/info/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
