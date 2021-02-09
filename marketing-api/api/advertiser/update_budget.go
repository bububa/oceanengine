package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// 更新账户日预算
// 此接口可以更新广告主账号设置的预算类型与预算。
// 账户预算类别：不限预算、日预算
// 范围1000<=X <= 9999999.99、仅支持最多2位小数
func UpdateBudget(clt *core.SDKClient, accessToken string, req *advertiser.UpdateBudgetRequest) (*model.BaseResponse, error) {
	var resp model.BaseResponse
	err := clt.Post("2/advertiser/update/budget/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
