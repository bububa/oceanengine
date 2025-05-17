package security

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/security"
)

// ScoreViolationEventGet 查询违规积分明细
func ScoreViolationEventGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *security.ScoreViolationEventGetRequest) (*security.ScoreViolationEventGetResult, error) {
	var resp security.ScoreViolationEventGetResponse
	if err := clt.GetAPI(ctx, "v3.0/security/score_violation_event/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
