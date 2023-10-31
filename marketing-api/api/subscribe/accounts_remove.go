package subscribe

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/subscribe"
)

// AccountsRemove 删除 Adv 订阅
func AccountsRemove(clt *core.SDKClient, accessToken string, req *subscribe.AccountsRemoveRequest) error {
	return clt.Post("v3.0/subscribe/accounts/remove/", req, nil, accessToken)
}
