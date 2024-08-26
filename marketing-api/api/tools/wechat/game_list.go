package wechat

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/wechat"
)

// GameList 获取微信小游戏列表
func GameList(ctx context.Context, clt *core.SDKClient, accessToken string, req *wechat.GameListRequest) (*wechat.GameListResult, error) {
	var resp wechat.GameListResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/wechat_game/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
