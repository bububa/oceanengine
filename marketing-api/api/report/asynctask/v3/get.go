package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/asynctask"
	v3 "github.com/bububa/oceanengine/marketing-api/model/report/asynctask/v3"
)

// Get 获取任务列表
func Get(clt *core.SDKClient, accessToken string, req *v3.GetRequest) (*asynctask.GetResponseData, error) {
	var resp asynctask.GetResponse
	if err := clt.Get("v3.0/report/custom/async_task/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
