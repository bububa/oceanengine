package sharedwallet

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// WalletRelationGetRequest 资金共享-查询子钱包下绑定的adv列表 API Request
type WalletRelationGetRequest struct {
	// AccountID 账户ID，对于巨量AD、千川、本地推是广告主ID，对于代理商是代理商ID
	AccountID uint64 `json:"account_id,omitempty"`
	// SharedWalletID 共享钱包ID
	SharedWalletID uint64 `json:"shared_wallet_id,omitempty"`
	// AccountType 账户类型 可选值:
	// AD 巨量广告主
	// AGENT 代理
	// LOCAL 本地推广告主
	AccountType string `json:"account_type,omitempty"`
	// Page 页码；默认为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小；默认为10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r WalletRelationGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("shared_wallet_id", strconv.FormatUint(r.SharedWalletID, 10))
	values.Set("account_type", string(r.AccountType))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// WalletRelationGetResponse 资金共享-查询子钱包下绑定的adv列表 API Response
type WalletRelationGetResponse struct {
	model.BaseResponse
	Data *WalletRelationGetResult `json:"data,omitempty"`
}

type WalletRelationGetResult struct {
	// PageInfo 翻页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Results 共享钱包关系
	Results []WalletRelation `json:"results,omitempty"`
	// SharedWalletID 共享钱包ID
	SharedWalletID uint64 `json:"shared_wallet_id,omitempty"`
}

// WalletRelation 共享钱包关系
type WalletRelation struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SharedWalletID 共享钱包ID
	SharedWalletID uint64 `json:"shared_wallet_id,omitempty"`
}
