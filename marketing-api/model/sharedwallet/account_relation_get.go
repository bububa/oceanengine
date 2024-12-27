package sharedwallet

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AccountRelationGetRequest 资金共享-查询账户对应公司下的钱包关系 API Request
type AccountRelationGetRequest struct {
	// AccountID 账户ID，对于巨量AD、千川、本地推是广告主ID，对于代理商是代理商ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型 可选值:
	// AD 巨量广告主
	// AGENT 代理
	// LOCAL 本地推广告主
	AccountType string `json:"account_type,omitempty"`
}

// Encode implements GetRequest interface
func (r AccountRelationGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AccountRelationGetResponse 资金共享-查询账户对应公司下的钱包关系 API Response
type AccountRelationGetResponse struct {
	model.BaseResponse
	Data *AccountRelationGetResult `json:"data,omitempty"`
}

type AccountRelationGetResult struct {
	// CompanyID 客户公司ID
	CompanyID uint64 `json:"company_id,omitempty"`
	// MainWalletID 共享钱包ID
	MainWalletID uint64 `json:"main_wallet_id,omitempty"`
	// SubWalletIDs 子钱包ID列表
	SubWalletIDs []uint64 `json:"sub_wallet_ids,omitempty"`
}
