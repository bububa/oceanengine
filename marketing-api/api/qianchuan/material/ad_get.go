package material

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/material"
)

// AdGet 获取素材关联计划
func AdGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *material.AdGetRequest) (*material.AdGetResult, error) {
	var resp material.AdGetResponse
	if err := clt.GetAPI(ctx, "v1.0/qianchuan/material/ad/get/", req, resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
