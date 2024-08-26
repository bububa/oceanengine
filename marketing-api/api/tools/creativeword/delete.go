package creativeword

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/creativeword"
)

// Delete 删除动态创意词包
// 使用此接口可以删除已创建的动态创意词包。
func Delete(ctx context.Context, clt *core.SDKClient, accessToken string, req *creativeword.DeleteRequest) (uint64, error) {
	var resp creativeword.DeleteResponse
	err := clt.Post(ctx, "2/tools/creative_word/delete/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.CreativeWordID, nil
}
