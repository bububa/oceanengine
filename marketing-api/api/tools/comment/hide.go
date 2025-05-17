package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// Hide 隐藏评论
func Hide(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.OperateRequest) ([]uint64, error) {
	var resp comment.OperateResponse
	if err := clt.PostAPI(ctx, "v3.0/tools/comment/hide/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.SuccessCommentIDs, nil
}
