package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	ad "github.com/bububa/oceanengine/marketing-api/model/ad/v3"
)

// RejectReason 获取计划审核建议
func RejectReason(clt *core.SDKClient, accessToken string, req *ad.RejectReasonRequest) ([]ad.RejectReason, error) {
	var resp ad.RejectReasonResponse
	if err := clt.Get("v3.0/promotion/reject_reason/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
