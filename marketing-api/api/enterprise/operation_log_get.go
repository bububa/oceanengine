package enterprise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// OperationLogGet 获取企业号推广操作记录
// 获取企业号在特定日期的推广操作记录，可查询的推广操作含直播开始、发布视频、开启DOU+推广、开启广告投放，并支持查询到具体的开播直播间ID、视频ID、DOU+推广金额，以及广告投放的账户与计划ID
func OperationLogGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *enterprise.OperationLogGetRequest) (*enterprise.OperationLogGetResult, error) {
	var resp enterprise.OperationLogGetResponse
	if err := clt.Get(ctx, "v1.0/enterprise/operation/log/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
