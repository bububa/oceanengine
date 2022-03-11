package fund

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// SharedWalletBalanceGetRequest 获取共享钱包余额 API Request
type SharedWalletBalanceGetRequest struct {
	// AdvertiserIDs 广告主ID列表，长度1-100
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r SharedWalletBalanceGetRequest) Encode() string {
	values := &url.Values{}
	ids, _ := json.Marshal(r.AdvertiserIDs)
	values.Set("advertiser_ids", string(ids))
	return values.Encode()
}

// SharedWalletBalanceGetResponse 获取共享钱包余额 API Response
type SharedWalletBalanceGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List 账户列表
		List []SharedWalletBalance `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// SharedWalletBalanceStatus  余额查询状态
type SharedWalletBalanceStatus string

const (
	// SharedWalletBalanceStatus_SUCCESS 成功
	SharedWalletBalanceStatus_SUCCESS SharedWalletBalanceStatus = "SUCCESS"
	// SharedWalletBalanceStatus_FAIL 失败
	SharedWalletBalanceStatus_FAIL SharedWalletBalanceStatus = "FAIL"
	// SharedWalletBalanceStatus_NO_WALLET 钱包不存在
	SharedWalletBalanceStatus_NO_WALLET SharedWalletBalanceStatus = "NO_WALLET"
)

// SharedWalletBalance 共享钱包余额
type SharedWalletBalance struct {
	// AdvertiserID 账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// BalanceDetail 余额详情列表
	BalanceDetail []SharedWalletBalanceDetail `json:"balance_detail,omitempty"`
	// Status 余额查询状态，SUCCESS：成功、FAIL：失败、NO_WALLET：钱包不存在
	Status SharedWalletBalanceStatus `json:"status,omitempty"`
	// StatusMessage 状态说明
	StatusMessage string `json:"status_message,omitempty"`
}

// SharedWalletBalanceBillingInventory 余额可用广告位
type SharedWalletBalanceBillingInventory string

const (
	// SharedWalletBalanceBillingInventory_DEFAULT 默认
	SharedWalletBalanceBillingInventory_DEFAULT SharedWalletBalanceBillingInventory = "DEFAULT"
	// SharedWalletBalanceBillingInventory_SEARCH 搜索
	SharedWalletBalanceBillingInventory_SEARCH SharedWalletBalanceBillingInventory = "SEARCH"
	// SharedWalletBalanceBillingInventory_UNION 穿山甲
	SharedWalletBalanceBillingInventory_UNION SharedWalletBalanceBillingInventory = "UNION"
	// SharedWalletBalanceBillingInventory_COMMON 通用
	SharedWalletBalanceBillingInventory_COMMON SharedWalletBalanceBillingInventory = "COMMON"
)

// SharedWalletBalanceDetail 余额详情
type SharedWalletBalanceDetail struct {
	// Balance 余额，单位：元，精度：两位小数
	Balance float64 `json:"balance,omitempty"`
	// BillingInventory 余额可用广告位，DEFAULT：默认、SEARCH：搜索、UNION：穿山甲、COMMON：通用
	BillingInventory string `json:"billing_inventory,omitempty"`
}
