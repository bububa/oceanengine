package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/privativeword/v3"
)

// Get 批量获取项目否定词
// 2.0项目批量获取否定词
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.GetRequest) ([]v3.Word, error) {
	var resp v3.GetResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/privative_word/project/batch_get", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
