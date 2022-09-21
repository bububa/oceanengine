package eventmanager

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

func TrackURLGet(clt *core.SDKClient, accessToken string, req *eventmanager.TrackURLGetRequest) (*eventmanager.TrackURLGetData, error) {
	var resp eventmanager.TrackURLGetResponse
	if err := clt.Get("2/event_manager/track_url/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
