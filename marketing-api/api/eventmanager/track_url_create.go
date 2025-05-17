package eventmanager

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// TrackURLCreate 事件资产下创建监测链接组
func TrackURLCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *eventmanager.TrackURLCreateRequest) error {
	return clt.Post(ctx, "2/event_manager/track_url/create/", req, nil, accessToken)
}
