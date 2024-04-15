package advertiser

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BalanceGetRequest 获取账户余额 API Request
type BalanceGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r BalanceGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// BalanceGetResponse 获取账户余额 API Response
type BalanceGetResponse struct {
	model.BaseResponse
	Data *Balance `json:"data,omitempty"`
}

type Balance struct {
	// AdvertiserID 广告主ID或代理商ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AccountTotal 账户总余额
	AccountTotal float64 `json:"account_total,omitempty"`
	// AccountValid 账户可用总余额
	AccountValid float64 `json:"account_valid,omitempty"`
	// AccountFrozen 账户冻结总余额
	AccountFrozen float64 `json:"account_frozen,omitempty"`
	// AccountGeneralTotal 通用总余额
	AccountGeneralTotal float64 `json:"account_general_total,omitempty"`
	// AccountGeneralValid 通用可用余额
	AccountGeneralValid float64 `json:"account_general_valid,omitempty"`
	// AccountGeneralFrozen 通用冻结余额
	AccountGeneralFrozen float64 `json:"account_general_frozen,omitempty"`
	// AccountBiddingTotal 竞价总余额
	AccountBiddingTotal float64 `json:"account_bidding_total,omitempty"`
	// AccountBiddingValid 竞价可用余额
	AccountBiddingValid float64 `json:"account_bidding_valid,omitempty"`
	// AccountBiddingFrozen 竞价冻结余额
	AccountBiddingFrozen float64 `json:"account_bidding_frozen,omitempty"`
	// AccountBrandTotal 账户品牌总余额
	AccountBrandTotal float64 `json:"account_brand_total,omitempty"`
	// AccountBrandValid 账户品牌可用余额
	AccountBrandValid float64 `json:"account_brand_valid,omitempty"`
	// AccountBrandFrozen 账户品牌冻结余额
	AccountBrandFrozen float64 `json:"account_brand_frozen,omitempty"`
	// ShareGrantTotal 总共享赠款
	ShareGrantTotal float64 `json:"share_grant_total,omitempty"`
	// ShareWalletGeneralValid 共享钱包通用可用余额
	ShareWalletGeneralValid float64 `json:"share_wallet_general_valid,omitempty"`
	// ShareWalletBiddingValid 共享钱包竞价可用余额
	ShareWalletBiddingValid float64 `json:"share_wallet_bidding_valid,omitempty"`
	// ShareWalletBrandValid 共享钱包品牌可用余额
	ShareWalletBrandValid float64 `json:"share_wallet_brand_valid,omitempty"`
	// ShareWalletID 共享钱包id
	ShareWalletID string `json:"share_wallet_id,omitempty"`
	// ShareWalletName 共享钱包名称
	ShareWalletName string `json:"share_wallet_name,omitempty"`
	// ShareWalletTotal 共享钱包可用总余额
	ShareWalletTotal float64 `json:"share_wallet_total,omitempty"`
}
