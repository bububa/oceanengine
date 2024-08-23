package tools

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/tools"
)

// AllowCoupon 智能优惠券白名单
func AllowCoupon(clt *core.SDKClient, accessToken string, req *tools.AllowCouponRequest) (*tools.AllowCouponResult, error) {
	var resp tools.AllowCouponResponse
	if err := clt.GetAPI("v1.0/qianchuan/tools/allow_coupon/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
