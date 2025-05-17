package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// AccountRelationGet 共享钱包-查询账户对应公司下的钱包关系
func AccountRelationGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.AccountRelationGetRequest) (*sharedwallet.AccountRelationGetResult, error) {
	var resp sharedwallet.AccountRelationGetResponse
	if err := clt.GetAPI(ctx, "v3.0/shared_wallet/account_relation/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
