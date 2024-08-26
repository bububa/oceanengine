package dmp

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/dmp"
)

// AudienceListGet 获取人群管理列表
// 用于获取千川广告账户下人群管理列表（数据同PC后台【数据】-【人群管理】-【人群列表】）
func AudienceListGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.AudienceListGetRequest) (*dmp.AudienceListGetResult, error) {
	var resp dmp.AudienceListGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/audience_list/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
