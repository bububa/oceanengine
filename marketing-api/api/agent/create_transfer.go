package agent

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// CreateTransfer 转账-发起转账（代理）
// 发起转账，支持1:N转账、不停投转账、虚客互转
func CreateTransfer(clt *core.SDKClient, accessToken string, req *agent.CreateTransferRequest) (string, error) {
	var resp agent.CreateTransferResponse
	if err := clt.Post("v3.0/cg_transfer/create_transfer/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.TransferSerial, nil
}
