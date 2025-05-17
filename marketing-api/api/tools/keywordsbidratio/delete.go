package keywordsbidratio

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/keywordsbidratio"
)

// Delete 删除优词提量系数和生效维度
func Delete(ctx context.Context, clt *core.SDKClient, accessToken string, req *keywordsbidratio.DeleteRequest) error {
	return clt.PostAPI(ctx, "v3.0/tools/keywords_bid_ratio/delete/", req, nil, accessToken)
}
