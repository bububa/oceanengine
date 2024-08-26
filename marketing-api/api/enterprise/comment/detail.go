package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise/comment"
)

// Detail 获取评论从属信息
// 根据评论id，获取评论的上级评论和所属视频信息。
func Detail(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.DetailRequest) (*comment.DetailResult, error) {
	var resp comment.DetailResponse
	if err := clt.Get(ctx, "v1.0/enterprise/comment/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
