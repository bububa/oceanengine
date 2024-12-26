package customercenter

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/customercenter"
)

// TransferCreate 工作台转账-发起转账
// 发起转账，支持1:N转账、不停投转账
func TransferCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *customercenter.TransferCreateRequest) (string, error) {
	var resp customercenter.TransferCreateResponse
	if err := clt.PostAPI(ctx, "v3.0/cg_transfer/transfer/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.TransferSerial, nil
}
