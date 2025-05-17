package security

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/security"
)

// ScoreDisposalInfoGet 查看积分处置详情
func ScoreDisposalInfoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *security.ScoreDisposalInfoGetRequest) (*security.ScoreDisposalInfoGetResult, error) {
	var resp security.ScoreDisposalInfoGetResponse
	if err := clt.GetAPI(ctx, "v3.0/security/score_disposal_info/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
