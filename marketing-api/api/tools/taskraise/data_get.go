package taskraise

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/taskraise"
)

// DataGet 查询优选起量任务数据
func DataGet(clt *core.SDKClient, accessToken string, req *taskraise.DataGetRequest) (*taskraise.DataGetResponseData, error) {
	var resp taskraise.DataGetResponse
	if err := clt.Get("2/tools/task_raise/data/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
