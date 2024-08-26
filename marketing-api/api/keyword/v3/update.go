package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/keyword"
	v3 "github.com/bububa/oceanengine/marketing-api/model/keyword/v3"
)

// Update 更新关键词
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.UpdateRequest) (*keyword.ResponseData, error) {
	var resp keyword.Response
	if err := clt.PostAPI(ctx, "v3.0/keyword/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
