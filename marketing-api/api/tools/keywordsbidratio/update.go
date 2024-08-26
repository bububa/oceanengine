package keywordsbidratio

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/keywordsbidratio"
)

// Update 更新优词提量系数和生效维度
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *keywordsbidratio.UpdateRequest) error {
	return clt.PostAPI(ctx, "v3.0/tools/keywords_bid_ratio/update/", req, nil, accessToken)
}
