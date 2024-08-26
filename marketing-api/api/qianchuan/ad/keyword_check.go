package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// KeywordCheck 关键词合规校验
// 用于确认创建搜索广告之前提供关键词合规校验及违规提醒
func KeywordCheck(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.KeywordCheckRequest) (*ad.KeywordCheckData, error) {
	var resp ad.KeywordCheckResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/keyword/check/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
