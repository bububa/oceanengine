package rta

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rta"
)

// ExpGet 获取穿山甲渠道RTA联合实验数据
// 功能
// 该接口用于查询穿山甲渠道的RTA联合实验数据
func RtaExpGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *rta.RtaExpGetRequest) ([]rta.Report, error) {
	var resp rta.RtaExpGetResponse
	if err := clt.Get(ctx, "2/report/rta_exp/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Data, nil
}
