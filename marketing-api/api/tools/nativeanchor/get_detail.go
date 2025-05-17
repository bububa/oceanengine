package nativeanchor

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/nativeanchor"
)

// GetDetail 获取原生锚点详情
// 根据锚点唯一id，获取到锚点详情，支持查询账户下锚点的详情（包括被共享和自有锚点）
// 暂不支持获取「高级在线预约」、「字节小程序」锚点详情
func GetDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *nativeanchor.GetDetailRequest) (*nativeanchor.GetDetailResult, error) {
	var resp nativeanchor.GetDetailResponse
	if err := clt.GetAPI(ctx, "v3.0/native_anchor/get/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
