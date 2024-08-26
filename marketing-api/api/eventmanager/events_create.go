package eventmanager

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// EventsCreate 资产下创建事件
// 此接口用于在资产下创建单个事件
// 需要先通过【获取推广内容】接口获取要创建的资产ID（asset_id）
// 需要先通过【获取可创建事件】接口获取到某资产下支持创建的事件ID（event_id）
// 当资产类型是三方落地页时，同资产同事件支持多种回传方式，但XPATH不能与JSSDK和EXTERNAL_API同时存在
func EventsCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *eventmanager.EventsCreateRequest) error {
	return clt.Post(ctx, "2/event_manager/events/create/", req, nil, accessToken)
}
