package track

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/track"
)

// WxaActive 微信小程序转化回传
func WxaActive(ctx context.Context, clt *core.SDKClient, req *track.WxaActiveRequest) error {
	return clt.TrackActive(ctx, req, nil)
}
