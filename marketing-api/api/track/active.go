package track

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/track"
)

// Active API上报数据(new)
func Active(ctx context.Context, clt *core.SDKClient, req *track.ActiveRequest) (int, error) {
	var resp track.Response
	if err := clt.TrackActive(ctx, req, &resp); err != nil && resp.Code != 128 {
		return resp.Code, err
	}
	return resp.Code, nil
}
