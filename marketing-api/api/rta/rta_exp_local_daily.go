package rta

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/rta"
)

func RtaExpLocalDailyGet(clt *core.SDKClient, accessToken string, req *rta.GetRtaExpLocalDailyReq) ([]*rta.GetRtaExpLocalDailyData, error) {
	var resp rta.GetRtaExpLocalDailyResponse
	err := clt.Get("v3.0/report/rta_exp_local_daily/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.Data, nil
}
