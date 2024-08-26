package privativeword

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/privativeword"
)

// CampaignAdd 批量新增组否定词
// 批量新增组否定词，即增量更新组否定词
// 对于单个广告组，最多200个短语否定词或200个精确否定词，若超过则截断。
// 多次操作添加否定词，短语否定词或精确否定词总量不能大于200个，若超过则截断。
func CampaignAdd(ctx context.Context, clt *core.SDKClient, accessToken string, req *privativeword.CampaignAddRequest) (*privativeword.CampaignAddResponseData, error) {
	var resp privativeword.CampaignAddResponse
	err := clt.Post(ctx, "2/privative_word/campaign/add", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
