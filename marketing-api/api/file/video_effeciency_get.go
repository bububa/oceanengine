package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoEfficiencyGet 获取低效素材
// 支持查询素材是否是低效素材，传入素材ID列表，返回低效素材列表。
func VideoEfficiencyGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoEfficiencyGetRequest) ([]string, error) {
	var resp file.VideoEfficiencyGetResponse
	if err := clt.GetAPI(ctx, "2/file/video/efficiency/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.IDs, nil
}
