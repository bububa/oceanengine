package wechat

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/wechat"
)

// GameCreate 创建微信小游戏
func GameCreate(clt *core.SDKClient, accessToken string, req *wechat.GameCreateRequest) (*wechat.WechatGame, error) {
	var resp wechat.GameCreateResponse
	if err := clt.Post("v3.0/tools/wechat_game/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Data, nil
}
