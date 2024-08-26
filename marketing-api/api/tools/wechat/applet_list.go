package wechat

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/wechat"
)

// AppletList 获取微信小程序列表
func AppletList(ctx context.Context, clt *core.SDKClient, accessToken string, req *wechat.AppletListRequest) (*wechat.AppletListResult, error) {
	var resp wechat.AppletListResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/wechat_applet/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
