package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// AwemeBannedDelete 删除屏蔽用户
// 删除屏蔽用户接口用于删除已屏蔽的抖音用户，抖音用户屏蔽需要广告主绑定抖音账号。
func AwemeBannedDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.AwemeBannedCreateRequest) (*comment.AwemeBannedCreateResponseData, error) {
	var resp comment.AwemeBannedCreateResponse
	err := clt.PostAPI(ctx, "v3.0/tools/aweme_banned/delete/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
