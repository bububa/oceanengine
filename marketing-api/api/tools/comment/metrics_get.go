package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// MetricsGet 获取评论统计指标
// 获取广告账户下评论的「可见评论数」、「可见负评数」、「可见评论负评率」
func MetricsGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.MetricsGetRequest) (*comment.MetricsGetResult, error) {
	var resp comment.MetricsGetResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/comment_metrics/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
