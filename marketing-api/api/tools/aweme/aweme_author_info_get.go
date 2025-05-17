package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/aweme"
)

// AwemeAuthorInfoGet 查询抖音号id对应的达人信息
// 查询抖音号id对应的抖音达人信息。使用场景如下：当计划中设置抖音达人账号时，可以根据其抖音账号id，查询对应的抖音达人信息。
func AwemeAuthorInfoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.AwemeAuthorInfoGetRequest) ([]aweme.Author, error) {
	var resp aweme.AwemeAuthorInfoGetResponse
	err := clt.Get(ctx, "2/tools/aweme_author_info/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
