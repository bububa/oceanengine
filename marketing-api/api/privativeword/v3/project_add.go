package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/privativeword/v3"
)

// Add 批量添加项目否定词
// 2.0批量添加否定词
func Add(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.AddRequest) (*v3.AddResponseData, error) {
	var resp v3.AddResponse
	if err := clt.PostAPI(ctx, "v3.0/tools/privative_word/project/add", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
