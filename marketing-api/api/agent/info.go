package agent

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// 获取代理商信息
func Info(clt *core.SDKClient, accessToken string, req *agent.InfoRequest) (*agent.InfoResponse, error) {
	var resp agent.InfoResponse
	err := clt.Get("2/agent/info/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
