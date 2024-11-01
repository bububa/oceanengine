package security

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/security"
)

// ScoreTotalGet 查询账户累计积分
func ScoreTotalGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *security.ScoreTotalGetRequest) (*security.ScoreTotalGetResult, error) {
	var resp security.ScoreTotalGetResponse
	if err := clt.GetAPI(ctx, "v3.0/security/score_total/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
