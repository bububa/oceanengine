package coupon

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/coupon"
)

// EmployeeGet 查询核销员
func EmployeeGet(clt *core.SDKClient, accessToken string, req *coupon.EmployeeGetRequest) (*coupon.EmployeeGetResponseData, error) {
	var resp coupon.EmployeeGetResponse
	if err := clt.Get("2/clue/coupon/employee/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
