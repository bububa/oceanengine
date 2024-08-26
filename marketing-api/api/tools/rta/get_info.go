package rta

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rta"
)

// GetInfo 获取RTA策略数据 API Response
func GetInfo(ctx context.Context, clt *core.SDKClient, accessToken string, req *rta.GetInfoRequest) (*rta.GetInfoData, error) {
	var resp rta.GetInfoResponse
	if err := clt.Get(ctx, "2/tools/rta/get_info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
