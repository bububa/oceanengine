package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/keyword"
	v3 "github.com/bububa/oceanengine/marketing-api/model/keyword/v3"
)

// Delete 删除关键词
// 删除指定keyword_id的搜索词，可批量删除。
func Delete(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.DeleteRequest) (*keyword.DeleteResponseData, error) {
	var resp keyword.DeleteResponse
	if err := clt.PostAPI(ctx, "v3.0/keyword/delete/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
