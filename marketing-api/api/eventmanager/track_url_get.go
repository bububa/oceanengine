package eventmanager

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

func TrackURLGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *eventmanager.TrackURLGetRequest) (*eventmanager.TrackURLGetData, error) {
	var resp eventmanager.TrackURLGetResponse
	if err := clt.Get(ctx, "2/event_manager/track_url/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
