package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// AwemeBannedList 获取屏蔽用户列表
// 获取屏蔽用户列表用于获取广告主绑定的抖音账号（第一个绑定的抖音账号）屏蔽的抖音id列表或昵称关键词列表。
func AwemeBannedList(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.AwemeBannedListRequest) (*comment.AwemeBannedListResponseData, error) {
	var resp comment.AwemeBannedListResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/aweme_banned/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
