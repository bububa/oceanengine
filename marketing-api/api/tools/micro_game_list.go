package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// MicroGameList 获取字节小游戏
// 获取字节小游戏列表，对应在巨量工作台上的字节小游戏资产
func MicroGameList(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.MicroAppListRequest) (*tools.MicroAppListResult, error) {
	var resp tools.MicroAppListResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/micro_game/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
