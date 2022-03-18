package event

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/event"
)

// AvailableEventsGet 获取可创建事件列表
// 此接口用于查询资产下支持创建的事件列表
func AvailableEventsGet(clt *core.SDKClient, accessToken string, req *event.AvailableEventsGetRequest) ([]event.EventConfig, error) {
	var resp event.AvailableEventsGetResponse
	err := clt.Get("2/tools/event_manager/available_events/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
