package comment

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/comment"
)

// AwemeBandedCreate 添加屏蔽用户
// 添加屏蔽用户接口用于添加屏蔽抖音用户，抖音用户屏蔽需要广告主绑定抖音账号，每个绑定的抖音账号最多屏蔽5000个抖音id与5000个昵称关键词，广告主当前可添加屏蔽的抖音id和昵称关键词以绑定的抖音账号中已屏蔽最多的抖音账号为准。
func AwemeBandedCreate(clt *core.SDKClient, accessToken string, req *comment.AwemeBandedCreateRequest) (*comment.AwemeBandedCreateResponseData, error) {
	var resp comment.AwemeBandedCreateResponse
	err := clt.Post("2/tools/comment/aweme_banded/create/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
