package coupon

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/coupon"
)

// EmployeeCreate 添加核销员
func EmployeeCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *coupon.EmployeeCreateRequest) (*coupon.EmployeeCreateResponseData, error) {
	var resp coupon.EmployeeCreateResponse
	if err := clt.Post(ctx, "2/clue/coupon/employee/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
