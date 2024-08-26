package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// TemplateGet 获取DPA模板
func TemplateGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.TemplateGetRequest) (*dpa.TemplateGetResponseData, error) {
	var resp dpa.TemplateGetResponse
	if err := clt.Get(ctx, "2/dpa/template/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
