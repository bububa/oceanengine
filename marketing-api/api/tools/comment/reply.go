package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// Reply 回复评论
func Reply(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.OperateRequest) ([]uint64, error) {
	var resp comment.OperateResponse
	if err := clt.PostAPI(ctx, "v3.0/tools/comment/reply/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.SuccessCommentIDs, nil
}
