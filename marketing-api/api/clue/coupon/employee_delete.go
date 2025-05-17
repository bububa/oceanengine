package coupon

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/coupon"
)

// EmployeeDelete 删除核销员
func EmployeeDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *coupon.EmployeeDeleteRequest) error {
	return clt.Post(ctx, "2/clue/coupon/employee/delete/", req, nil, accessToken)
}
