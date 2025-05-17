package comment

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// AwemeBannedCreate 添加屏蔽用户
// 添加屏蔽用户接口用于添加屏蔽抖音用户，抖音用户屏蔽需要广告主绑定抖音账号，每个绑定的抖音账号最多屏蔽5000个抖音id与5000个昵称关键词，广告主当前可添加屏蔽的抖音id和昵称关键词以绑定的抖音账号中已屏蔽最多的抖音账号为准。
func AwemeBannedCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.AwemeBannedCreateRequest) (*comment.AwemeBannedCreateResponseData, error) {
	var resp comment.AwemeBannedCreateResponse
	if err := clt.PostAPI(ctx, "v3.0/tools/aweme_banned/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
