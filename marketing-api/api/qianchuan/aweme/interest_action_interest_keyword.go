package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// InterestActionInterestKeyword 获取随心推兴趣标签
// 该接口用于获取随心推创编链路上使用的兴趣标签
func InterestActionInterestKeyword(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) ([]aweme.InterestKeyword, error) {
	var resp aweme.InterestActionInterestKeywordResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/aweme/interest_action/interest/keyword/", &aweme.InterestActionInterestKeywordRequest{AdvertiserID: advertiserID}, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
