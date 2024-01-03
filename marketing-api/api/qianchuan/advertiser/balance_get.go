package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/advertiser"
)

// BalanceGet 获取账户余额
// 获取账户余额，包含通用，竞价，品牌余额
func BalanceGet(clt *core.SDKClient, accessToken string, req *advertiser.BalanceGetRequest) (*advertiser.Balance, error) {
	var resp advertiser.BalanceGetResponse
	if err := clt.Get("v1.0/qianchuan/account/balance/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
