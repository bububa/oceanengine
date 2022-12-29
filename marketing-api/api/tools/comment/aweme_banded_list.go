package comment

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// AwemeBandedList 获取屏蔽用户列表
// 获取屏蔽用户列表用于获取广告主绑定的抖音账号（第一个绑定的抖音账号）屏蔽的抖音id列表或昵称关键词列表。
func AwemeBandedList(clt *core.SDKClient, accessToken string, req *comment.AwemeBandedListRequest) (*comment.AwemeBandedListResponseData, error) {
	var resp comment.AwemeBandedListResponse
	if err := clt.Get("v3.0/tools/comment/aweme_banded/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
