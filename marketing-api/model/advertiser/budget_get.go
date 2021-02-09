package advertiser

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type BudgetGetRequest struct {
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"` // 广告主ID集合(如果包含没有访问权限的ID,将返回no permission error) 取值范围: 1-100
}

func (r BudgetGetRequest) Encode() string {
	idsBytes, _ := json.Marshal(r.AdvertiserIDs)
	values := &url.Values{}
	values.Set("advertiser_ids", string(idsBytes))
	return values.Encode()
}

type BudgetGetResponse struct {
	model.BaseResponse
	Data *BudgetGetResponseData `json:"data,omitempty"`
}

type BudgetGetResponseData struct {
	List []BudgetGetResponseList `json:"list,omitempty"`
}

type BudgetGetResponseList struct {
	AdvertiserID uint64          `json:"advertiser_id,omitempty"` // 广告主ID
	Budget       float64         `json:"budget,omitempty"`        // 预算，单位：元； 精度：小数点后两位；举例：100.01。当预算类型为不限，返回的预算为0.0
	BudgetMode   enum.BudgetMode `json:"budget_mode,omitempty"`   // 预算类型
}
