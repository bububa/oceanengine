package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/file/v3"
)

// QualityGet 连山投前分析结果查询
// 素材属性结果查询
func QualityGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.QualityGetRequest) ([]v3.MaterialQuality, error) {
	var resp v3.QualityGetResponse
	if err := clt.GetAPI(ctx, "v3.0/file/quality/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
