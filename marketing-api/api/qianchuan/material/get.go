package material

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/material"
)

// Get 获取账户下素材列表
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *material.GetRequest) (*material.GetResult, error) {
	var resp material.GetResponse
	if err := clt.GetAPI(ctx, "v1.0/qianchuan/material/get/", req, resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
