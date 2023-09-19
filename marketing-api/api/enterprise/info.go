package enterprise

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// Info 获取企业号信息
func Info(clt *core.SDKClient, accessToken string, req *enterprise.InfoRequest) ([]enterprise.Enterprise, error) {
	var resp enterprise.InfoResponse
	if err := clt.Get("v1.0/enterprise/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
