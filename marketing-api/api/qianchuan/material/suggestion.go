package material

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/material"
)

// Suggestion 计划下素材审核建议
func Suggestion(ctx context.Context, clt *core.SDKClient, accessToken string, req *material.SuggestionRequest) (*material.SuggestionResult, error) {
	var resp material.SuggestionResponse
	if err := clt.GetAPI(ctx, "v1.0/qianchuan/ad/material/suggestion/", req, resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
