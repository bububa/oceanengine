package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// Mid2ItemID 获取评论视频ID列表
// 提供由mid查询对应的item_id的能力
// 该接口仅适用于评论管理，仅返回有评论内容的对应抖音视频。若相关素材id的关联item_id均无评论，则返回为空
func Mid2ItemID(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.Mid2ItemIDRequest) (*comment.Mid2ItemIDList, error) {
	var resp comment.Mid2ItemIDResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/comment/mid2item_id/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
