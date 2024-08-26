package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise/comment"
)

// ReplyList 获取评论的回复
// 根据评论id，获取评论的回复内容。
func ReplyList(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ReplyListRequest) (*comment.ReplyListResult, error) {
	var resp comment.ReplyListResponse
	if err := clt.Get(ctx, "v1.0/enterprise/comment/reply/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
