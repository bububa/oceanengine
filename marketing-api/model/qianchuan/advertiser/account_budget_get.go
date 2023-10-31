package advertiser

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AccountBudgetGetRequest 获取账户日预算
type AccountBudgetGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r AccountBudgetGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AccountBudgetGetResponse 获取账户日预算
type AccountBudgetGetResponse struct {
	model.BaseResponse
	Data *Budget `json:"data,omitempty"`
}

type Budget struct {
	// BudgetMode 预算类型
	// INFINITE 不限
	// SPECIFIED 日预算
	BudgetMode qianchuan.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 账户今日竞价预算，单位元
	Budget float64 `json:"budget,omitempty"`
}
