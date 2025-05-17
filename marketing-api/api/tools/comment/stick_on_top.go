package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// StickOnTop 置顶评论
func StickOnTop(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.StickOnTopRequest) error {
	return clt.PostAPI(ctx, "v3.0/tools/comment/stick_on_top/", req, nil, accessToken)
}
