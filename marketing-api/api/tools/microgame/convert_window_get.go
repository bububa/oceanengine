package microgame

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model/tools/microgame"
)

// ConvertWindowGet 查询字节小游戏归因激活时间窗
// 使用场景：可通过该接口查询当前字节小游戏最新的归因激活时间窗配置内容，如希望修改归因激活时间窗，需要前往编辑接口。
func ConvertWindowGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *microgame.ConvertWindowGetRequest) (enum.ConvertWindow, error) {
	var resp microgame.ConvertWindowGetResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/micro_game/convert_window/get/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.ConvertWindow, nil
}
