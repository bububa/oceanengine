package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/ad"
)

// UpdateStatus 更新计划状态
// 通过此接口可对计划做启用/暂停/删除操作；
// 一次可以处理100个计划
// 对于删除的计划不能再进行状态操作，否则会报错！
// 如果有一个计划有问题，全部计划修改都不会成功！请确保传入的计划属于此广告主以及处于非删除状态。
func UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.UpdateStatusRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	err := clt.Post(ctx, "2/ad/update/status/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
