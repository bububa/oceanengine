package wechat

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/wechat"
)

// AppletCreate 创建微信小程序
func AppletCreate(clt *core.SDKClient, accessToken string, req *wechat.AppletCreateRequest) (uint64, error) {
	var resp wechat.AppletCreateResponse
	if err := clt.PostAPI("v3.0/tools/wechat_applet/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.Data.InstanceID, nil
}
