package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AccountFundGetRequest 批量查询账户余额 API Request
type AccountFundGetRequest struct {
	// AccountIDs 需要查询的账户ids，支持广告主、代理商、星图客户、星图MCN、星图服务商、本地推
	// 注意：
	// 传入账户id的类型需统一
	// 最多支持20个账户
	AccountIDs []uint64 `json:"account_ids,omitempty"`
	// AccountType 账户业务类型，可选值:
	// AD AD
	// STAR 星图
	// LOCAL 本地推
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// GrantTypeSplit 是否拆分赠款类型，允许值：
	// ON 开启
	// OFF 关闭（默认）
	GrantTypeSplit enum.OnOff `json:"grant_type_split,omitempty"`
}

// Encode implement GetRequest interface
func (r AccountFundGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_ids", string(util.JSONMarshal(r.AccountIDs)))
	values.Set("account_type", string(r.AccountType))
	if r.GrantTypeSplit != "" {
		values.Set("grant_type_split", string(r.GrantTypeSplit))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AccountFundGetResponse 批量查询账号余额 API Response
type AccountFundGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []AccountFund `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

type AccountFund struct {
	// AccountID 账户id
	AccountID uint64 `json:"account_id,omitempty"`
	FundInfo
}
