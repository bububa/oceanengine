package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	diagnosis "github.com/bububa/oceanengine/marketing-api/model/tools/diagnosis/v3"
)

// SuggestionAccept 采纳计划诊断建议
// 通过采纳诊断建议接口，广告主可以采纳【获取诊断建议】接口所获得的所有建议
func SuggestionAccept(ctx context.Context, clt *core.SDKClient, accessToken string, req *diagnosis.SuggestionAcceptRequest) (*diagnosis.SuggestionAcceptResponseData, error) {
	var resp diagnosis.SuggestionAcceptResponse
	err := clt.PostAPI(ctx, "v3.0/tools/promotion_diagnosis/suggestion/accept/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
