package sharedwallet

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/sharedwallet"
)

// CanTransferBalance 资金共享-最大可转余额查询
// 支持查询减款方与加款方之间最大可转金额、减款方非品牌最大可转出金额、加款方非品牌最小转入金额，
// 为了最大限度保证转账申请成功率，在发起转账申请之前，建议调用方做以下校验：
// - 减款钱包非品牌资金申请转账金额 <= 减款钱包非品牌资金最大可转出金额（non_brand_max_transfer_balance）
// - 减款钱包资金-业务线维度申请转账金额  <= 减款钱包资金-业务线维度可转金额（remitter_capital_detail_list）
// - 加款钱包非品牌资金申请转账金额 >= 加款钱包非品牌资金最小转入金额（non_brand_min_transfer_balance）
// - 加款钱包资金-业务线维度申请转账金额 <= 加款钱包资金-业务线维度可转金额（payee_transfer_amount_detail_list.capital_detail_list）
func CanTransferBalance(ctx context.Context, clt *core.SDKClient, accessToken string, req *sharedwallet.CanTransferBalanceRequest) ([]sharedwallet.CanTransferDetail, error) {
	var resp sharedwallet.CanTransferBalanceResponse
	if err := clt.GetAPI(ctx, "v3.0/cg_transfer/wallet/transfer/can_transfer_balance/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.CanTransferDetailList, nil
}
