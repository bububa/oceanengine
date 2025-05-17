package customercenter

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/customercenter"
)

// CanTransferTargetList 工作台转账-获取可转列表
// 查询当前账户(锚定账户)可以互相转账的账户列表
func CanTransferTargetList(ctx context.Context, clt *core.SDKClient, accessToken string, req *customercenter.CanTransferTargetListRequest) (*customercenter.CanTransferTargetListResult, error) {
	var resp customercenter.CanTransferTargetListResponse
	if err := clt.GetAPI(ctx, "v3.0/cg_transfer/can_transfer_target/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
