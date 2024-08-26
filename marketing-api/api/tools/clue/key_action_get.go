package clue

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/clue"
)

// KeyActionGet 获取活动记录
// 获取广告主的飞鱼获取活动记录（目前白名单开放，如需使用请联系对应销售）。为保证接口使用的安全性避免调取他人的线索信息，该接口只可用于查询自己广告主下的线索信息，即需查询的广告主账号的主体需与APPID对应开发者的主体保持一致，才可获取到线索的信息，否则会报错！
func KeyActionGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *clue.KeyActionGetRequest) (*clue.KeyActionGetResponseData, error) {
	var resp clue.KeyActionGetResponse
	if err := clt.Get(ctx, "2/tools/key_action/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
