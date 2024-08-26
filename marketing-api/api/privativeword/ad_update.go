package privativeword

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/privativeword"
)

// AdUpdate 设置计划否定词
// 单个设置计划否定词，即全量更新单个计划否定词
// phrase_words 与 precise_words 如果传入空数组，则更新为空。两者至少传入一个，如果只传一个，另一个字段则更新为空。
// 如果单个否定词有误，则计划否定词全部更新失败。
func AdUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *privativeword.AdUpdateRequest) (uint64, error) {
	var resp privativeword.AdUpdateResponse
	err := clt.Post(ctx, "2/privative_word/ad/update", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.AdID, nil
}
