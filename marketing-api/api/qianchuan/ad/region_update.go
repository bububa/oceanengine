package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// RegionUpdate 更新计划地域定向
// 批量更新计划定向地域，已关联定向包的计划和托管计划不支持修改地域
func RegionUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.RegionUpdateRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/ad/region/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
