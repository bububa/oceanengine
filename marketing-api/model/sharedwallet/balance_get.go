package sharedwallet

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BalanceGetRequest 获取共享钱包余额 API Request
type BalanceGetRequest struct {
	// AdvertiserIDs 广告主ID列表，长度1-100
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r BalanceGetRequest) Encode() string {
	values := util.GetUrlValues()
	ids, _ := json.Marshal(r.AdvertiserIDs)
	values.Set("advertiser_ids", string(ids))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// BalanceGetResponse 获取共享钱包余额 API Response
type BalanceGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List 账户列表
		List []BalanceInfo `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// BalanceStatus  余额查询状态
type BalanceStatus string

const (
	// BalanceStatus_SUCCESS 成功
	BalanceStatus_SUCCESS BalanceStatus = "SUCCESS"
	// BalanceStatus_FAIL 失败
	BalanceStatus_FAIL BalanceStatus = "FAIL"
	// BalanceStatus_NO_WALLET 钱包不存在
	BalanceStatus_NO_WALLET BalanceStatus = "NO_WALLET"
)

// BalanceInfo 共享钱包余额
type BalanceInfo struct {
	// AdvertiserID 账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// BalanceDetail 余额详情列表
	BalanceDetail []BalanceDetail `json:"balance_detail,omitempty"`
	// Status 余额查询状态，SUCCESS：成功、FAIL：失败、NO_WALLET：钱包不存在
	Status BalanceStatus `json:"status,omitempty"`
	// StatusMessage 状态说明
	StatusMessage string `json:"status_message,omitempty"`
}

// BalanceBillingInventory 余额可用广告位
type BalanceBillingInventory string

const (
	// BalanceBillingInventory_DEFAULT 默认
	BalanceBillingInventory_DEFAULT BalanceBillingInventory = "DEFAULT"
	// BalanceBillingInventory_SEARCH 搜索
	BalanceBillingInventory_SEARCH BalanceBillingInventory = "SEARCH"
	// BalanceBillingInventory_UNION 穿山甲
	BalanceBillingInventory_UNION BalanceBillingInventory = "UNION"
	// BalanceBillingInventory_COMMON 通用
	BalanceBillingInventory_COMMON BalanceBillingInventory = "COMMON"
)

// BalanceDetail 余额详情
type BalanceDetail struct {
	// Balance 余额，单位：元，精度：两位小数
	Balance float64 `json:"balance,omitempty"`
	// BillingInventory 余额可用广告位，DEFAULT：默认、SEARCH：搜索、UNION：穿山甲、COMMON：通用
	BillingInventory string `json:"billing_inventory,omitempty"`
}
