package clue

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/clue"
)

// Callback 回传有效线索
// 回传广告主的线索信息。为保证接口使用的安全性避免操作他人的线索信息，该接口只可用于操作自己广告主下的线索信息，即需查询的广告主账号的主体需与APPID对应开发者的主体保持一致，才可操作线索的信息，否则会报错！
func Callback(ctx context.Context, clt *core.SDKClient, accessToken string, req *clue.CallbackRequest) error {
	return clt.Post(ctx, "2/tools/clue/callback/", req, nil, accessToken)
}
