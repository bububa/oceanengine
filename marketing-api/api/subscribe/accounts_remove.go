package subscribe

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/subscribe"
)

// AccountsRemove 删除 Adv 订阅
func AccountsRemove(ctx context.Context, clt *core.SDKClient, accessToken string, req *subscribe.AccountsRemoveRequest) error {
	return clt.PostAPI(ctx, "v3.0/subscribe/accounts/remove/", req, nil, accessToken)
}
