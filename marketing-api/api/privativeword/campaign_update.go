package privativeword

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/privativeword"
)

// CampaignUpdate 设置组否定词
func CampaignUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *privativeword.AdUpdateRequest) (uint64, error) {
	var resp privativeword.AdUpdateResponse
	err := clt.Post(ctx, "2/privative_word/campaign/update", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.CampaignID, nil
}
