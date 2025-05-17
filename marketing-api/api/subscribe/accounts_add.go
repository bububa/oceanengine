package subscribe

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/subscribe"
)

// AccountsAdd 新增 Adv 订阅
func AccountsAdd(ctx context.Context, clt *core.SDKClient, accessToken string, req *subscribe.AccountsAddRequest) error {
	return clt.PostAPI(ctx, "v3.0/subscribe/accounts/add/", req, nil, accessToken)
}
