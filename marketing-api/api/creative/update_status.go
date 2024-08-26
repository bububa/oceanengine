package creative

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// UpdateStatus 更新创意状态
// 不可以操作已删除创意的状态，且程序化创意不能进行单独操作，需要从计划纬度进行操作
// 本接口为批量更新接口，调用结果针对单个创意存在部分成功部分失败场景，请避免根据应答code字段直接判断创意状态更新的结果
func UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.UpdateStatusRequest) (*creative.UpdateResponseData, error) {
	var resp creative.UpdateResponse
	err := clt.Post(ctx, "2/creative/status/update_v2/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
