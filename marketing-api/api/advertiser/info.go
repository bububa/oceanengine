package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// Info 广告主信息
// 获取广告主账户详细信息,可指定fields查询所需元素。
// 目前上线了新版的代理商鉴权，如果查询报no permission错误，请前往代理商一站式平台为账号申请对应的权限。
func Info(clt *core.SDKClient, accessToken string, req *advertiser.InfoRequest) (*advertiser.InfoResponse, error) {
	var resp advertiser.InfoResponse
	err := clt.Get("2/advertiser/info/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
