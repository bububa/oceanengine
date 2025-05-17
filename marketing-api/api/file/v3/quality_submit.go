package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/file/v3"
)

// QualitySubmit 连山视频投前分析提交
// 素材属性提交分析
func QualitySubmit(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.QualitySubmitRequest) (uint64, error) {
	var resp v3.QualitySubmitResponse
	if err := clt.PostAPI(ctx, "v3.0/file/quality/submit/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.MaterialID, nil
}
