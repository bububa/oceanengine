package privativeword

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/privativeword"
)

// AdAdd 批量新增计划否定词
// 批量新增计划否定词，即增量更新计划否定词
// 对于单个广告计划，最多200个短语否定词或200个精确否定词，若超过则截断。
// 多次操作添加否定词，短语否定词或精确否定词总量不能大于200个，若超过则截断。
func AdAdd(ctx context.Context, clt *core.SDKClient, accessToken string, req *privativeword.AdAddRequest) (*privativeword.AdAddResponseData, error) {
	var resp privativeword.AdAddResponse
	err := clt.Post(ctx, "2/privative_word/ad/add", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
