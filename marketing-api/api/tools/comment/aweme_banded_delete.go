package comment

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// AwemeBandedDelete 删除屏蔽用户
// 删除屏蔽用户接口用于删除已屏蔽的抖音用户，抖音用户屏蔽需要广告主绑定抖音账号。
func AwemeBandedDelete(clt *core.SDKClient, accessToken string, req *comment.AwemeBandedCreateRequest) (*comment.AwemeBandedCreateResponseData, error) {
	var resp comment.AwemeBandedCreateResponse
	err := clt.Post("2/tools/comment/aweme_banded/delete/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
