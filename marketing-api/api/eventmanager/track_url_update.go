package eventmanager

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// TrackURLUpdate 事件资产下更新监测链接组
func TrackURLUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *eventmanager.TrackURLUpdateRequest) error {
	return clt.Post(ctx, "2/event_manager/track_url/update/", req, nil, accessToken)
}
