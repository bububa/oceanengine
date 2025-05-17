package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// FundTransferSeqCreate 创建转账交易号（方舟）
// 转账接口分为两步, 包含创建转账接口 和 提交转账接口，本接口为第一步
func FundTransferSeqCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.FundTransferSeqCreateRequest) (string, error) {
	var resp agent.FundTransferSeqCreateResponse
	if err := clt.Post(ctx, "2/agent/fund/transfer_seq/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	if resp.Data.Status != 0 {
		return "", model.BaseResponse{
			Code:    resp.Data.Status,
			Message: resp.Data.StatusMessage,
		}
	}
	return resp.Data.TransferSeq, nil
}
