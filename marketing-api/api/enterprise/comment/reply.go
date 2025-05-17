package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise/comment"
)

// Reply 回复评论
func Reply(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ReplyRequest) (*comment.ReplyResult, error) {
	var resp comment.ReplyResponse
	if err := clt.Post(ctx, "v1.0/enterprise/comment/reply/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
