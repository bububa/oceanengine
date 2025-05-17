package dmp

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/dmp"
)

// AudienceGroupGet 获取人群分组
// 用于获取千川广告账户下的人群分组信息
func AudienceGroupGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.AudienceGroupGetRequest) ([]string, error) {
	var resp dmp.AudienceGroupGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/audience_group/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.AudienceGroupList, nil
}
