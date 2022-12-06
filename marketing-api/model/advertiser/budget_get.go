package advertiser

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// BudgetGetRequest 获取账户日预算 API Request
type BudgetGetRequest struct {
	// AdvertiserIDs 广告主ID集合(如果包含没有访问权限的ID,将返回no permission error) 取值范围: 1-100
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r BudgetGetRequest) Encode() string {
	idsBytes, _ := json.Marshal(r.AdvertiserIDs)
	values := util.GetUrlValues()
	values.Set("advertiser_ids", string(idsBytes))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// BudgetGetResponse 获取账户日预算 API Response
type BudgetGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *BudgetGetResponseData `json:"data,omitempty"`
}

// BudgetGetResponseData json返回值
type BudgetGetResponseData struct {
	// List 返回值列表
	List []BudgetGetResponseList `json:"list,omitempty"`
}

// BudgetGetResponseList 返回值列表
type BudgetGetResponseList struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Budget 预算，单位：元； 精度：小数点后两位；举例：100.01。当预算类型为不限，返回的预算为0.0
	Budget float64 `json:"budget,omitempty"`
	// BudgetMode 预算类型
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
}
