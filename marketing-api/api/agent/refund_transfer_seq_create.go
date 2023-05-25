package agent

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// RefundTransferSeqCreate 创建退款交易号（方舟）
// 退款接口分为两步, 包含创建退款接口 和 提交退款接口，本接口为第一步
func RefundTransferSeqCreate(clt *core.SDKClient, accessToken string, req *agent.FundTransferSeqCreateRequest) (string, error) {
	var resp agent.FundTransferSeqCreateResponse
	if err := clt.Post("2/agent/refund/transfer_seq/create/", req, &resp, accessToken); err != nil {
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
