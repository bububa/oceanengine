package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// PrivatewordsUpdate 全量更新否定词
// 用于全量更新搜索计划否定词
func PrivatewordsUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.PrivatewordsUpdateRequest) (*ad.PrivatewordsUpdateResult, error) {
	var resp ad.PrivatewordsUpdateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/ad/privatewords/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
