package eventmanager

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// EventConfigsGet 获取已创建事件列表
// 此接口用于查询资产下已经创建的事件及其相关信息
func EventConfigsGet(clt *core.SDKClient, accessToken string, req *eventmanager.EventConfigsGetRequest) ([]eventmanager.EventConfig, error) {
	var resp eventmanager.EventConfigsGetResponse
	err := clt.Get("2/event_manager/event_configs/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
