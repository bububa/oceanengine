package wechat

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/wechat"
)

// AppletUpdate 更新微信小程序
func AppletUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *wechat.AppletUpdateRequest) (*wechat.AppletUpdateResult, error) {
	var resp wechat.AppletUpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/tools/wechat_applet/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Data, nil
}
