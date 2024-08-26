package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/keyword"
	v3 "github.com/bububa/oceanengine/marketing-api/model/keyword/v3"
)

// Create 创建关键词
// 在指定的promotion_id下创建搜索关键词，支持在原有关键词基础上进行新增
// 目前仅支持搜索直投广告
// 创建关键词时会自动将优词添加为关键词，请您知悉
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.CreateRequest) (*keyword.ResponseData, error) {
	var resp keyword.Response
	if err := clt.PostAPI(ctx, "v3.0/keyword/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
